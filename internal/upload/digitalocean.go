package upload

import (
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToDigitalOcean(file multipart.File, header *multipart.FileHeader) (string, error) {
	log.Println("Starting upload to DigitalOcean")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("DO_SPACE_KEY"),
			os.Getenv("DO_SPACE_SECRET"),
			"",
		),
		Endpoint: aws.String(os.Getenv("DO_SPACE_ENDPOINT")),
		Region:   aws.String(os.Getenv("DO_SPACE_REGION")),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Printf("Error creating new session: %v", err)
		return "", err
	}

	s3Client := s3.New(newSession)

	log.Printf("Uploading file: %s", header.Filename)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("DO_SPACE_BUCKET")),
		Key:         aws.String(header.Filename),
		Body:        file,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(header.Header.Get("Content-Type")),
	})
	if err != nil {
		log.Printf("Error uploading file: %v", err)
		return "", err
	}

	log.Println("File uploaded successfully")

	fileURL := "https://" + os.Getenv("DO_SPACE_BUCKET") + "." + os.Getenv("DO_SPACE_REGION") + ".digitaloceanspaces.com/" + header.Filename

	return fileURL, nil
}
