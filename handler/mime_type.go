package handler

var validImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

func isAllowedImageType(mimeType string) bool {
	_, exists := validImageTypes[mimeType]

	return exists
}

var validFileTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"audio/mp3":  true,
	"audio/wave": true,
}

func isAllowedFileType(mimeType string) bool {
	_, exists := validFileTypes[mimeType]

	return exists
}
