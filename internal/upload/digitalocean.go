package upload

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"os"
)

func UploadToDigitalOcean(file multipart.File, header *multipart.FileHeader) (string, error) {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv("DO_SPACE_KEY"), os.Getenv("DO_SPACE_SECRET"), ""),
		Endpoint:    aws.String(os.Getenv("DO_SPACE_ENDPOINT")),
		Region:      aws.String(os.Getenv("DO_SPACE_REGION")),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return "", err
	}

	s3Client := s3.New(newSession)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("DO_SPACE_BUCKET")),
		Key:    aws.String(header.Filename),
		Body:   file,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}

	fileURL := "https://" + os.Getenv("DO_SPACE_BUCKET") + "." + os.Getenv("DO_SPACE_REGION") + ".digitaloceanspaces.com/" + header.Filename

	return fileURL, nil
}