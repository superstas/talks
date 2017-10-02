package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadToS3(body io.Reader) error {
	region := "us-east-2"
	bucketName, key := "test", "key"

	s3Sess := session.Must(session.NewSession(&aws.Config{
		Region: &region,
	}))

	s3Client := s3.New(s3Sess)
	s3Uploader := s3manager.NewUploaderWithClient(s3Client)

	_, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &key,
		Body:   body,
	})
	return err
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// 1 OMIT
	pr, pw := io.Pipe() // HL
	// uploading to S3
	go func(data io.Reader) {
		defer wg.Done()
		defer pr.Close()
		if err := uploadToS3(data); err != nil {
			log.Fatal(err)
		}
	}(pr)

	// raw data
	c, err := http.Get("https://golang.org/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Body.Close()

	http.Post()
	// data compression
	gz := gzip.NewWriter(pw)
	io.Copy(gz, c.Body)
	gz.Flush()
	gz.Close()
	pw.Close()
	// END 1 OMIT
	wg.Wait()
}
