package repository

import (
	"fmt"
	"log"
	"murmur-server/model"
	"net/smtp"
)

// mailRepository contains the gmail username and password
// as well as the frontend origin.
type mailRepository struct {
	username string
	password string
	origin   string
}

// NewMailRepository is a factory for initializing Mail Repositories
func NewMailRepository(username string, password string, origin string) model.MailRepository {
	return &mailRepository{
		username: username,
		password: password,
		origin:   origin,
	}
}

// SendResetMail sends a password reset email with the given reset token
func (m *mailRepository) SendResetMail(email string, token string) error {

	msg := "From: " + m.username + "\n" +
		"To: " + email + "\n" +
		"Subject: Reset Email\n\n" +
		fmt.Sprintf("<a href=\"%s/reset-password/%s\">Reset Password</a>", m.origin, token)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", m.username, m.password, "smtp.gmail.com"),
		m.username, []string{email}, []byte(msg))

	if err != nil {
		log.Printf("Failed to send email to %s: %v", email, err)
		return err
	}

	log.Printf("Email sent to %s successfully", email)
	return nil
}

func (m *mailRepository) SendVerificationMail(email string, token string) error {

	msg := "From: " + m.username + "\n" +
		"To: " + email + "\n" +
		"Subject: Reset Email\n\n" +
		fmt.Sprintf("<a href=\"%s/verification/%s\">Verification</a>", m.origin, token)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", m.username, m.password, "smtp.gmail.com"),
		m.username, []string{email}, []byte(msg))

	if err != nil {
		log.Printf("Failed to send email to %s: %v", email, err)
		return err
	}

	log.Printf("Email sent to %s successfully", email)
	return nil
}
