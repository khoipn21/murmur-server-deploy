package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"murmur-server/config"
	"murmur-server/model"
	"os"
)

type dataSources struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	S3Session   *session.Session
}

func initDS(ctx context.Context, cfg config.Config) (*dataSources, error) {
	log.Printf("Initializing data sources\n")

	log.Printf("Connecting to Postgresql\n")
	db, err := gorm.Open(postgres.Open(cfg.DatabaseUrl))

	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Migrate models and setup join tables
	if err = db.AutoMigrate(
		&model.User{},
		&model.Guild{},
		&model.Member{},
		&model.Channel{},
		&model.DMMember{},
		&model.Message{},
		&model.Attachment{},
		&model.VCMember{},
	); err != nil {
		return nil, fmt.Errorf("error migrating models: %w", err)
	}

	if err = db.SetupJoinTable(&model.Guild{}, "Members", &model.Member{}); err != nil {
		return nil, fmt.Errorf("error creating join table: %w", err)
	}

	if err = db.SetupJoinTable(&model.Guild{}, "VCMembers", &model.VCMember{}); err != nil {
		return nil, fmt.Errorf("error creating join table: %w", err)
	}

	// Initialize redis connection
	opt, err := redis.ParseURL(cfg.RedisUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing the redis url: %w", err)
	}

	log.Println("Connecting to Redis")
	rdb := redis.NewClient(opt)

	// verify redis connection
	_, err = rdb.Ping(ctx).Result()

	if err != nil {
		return nil, fmt.Errorf("error connecting to redis: %w", err)
	}

	// Initialize S3 Session
	sess, err := session.NewSession(
		&aws.Config{
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				"",
			),
			Region: aws.String(cfg.Region),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("error creating s3 session: %w", err)
	}

	return &dataSources{
		DB:          db,
		RedisClient: rdb,
		S3Session:   sess,
	}, nil
}

// close to be used in graceful server shutdown
func (d *dataSources) close() error {
	if err := d.RedisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}
