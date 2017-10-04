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

	// raw data
	c, err := http.Get("https://golang.org/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Body.Close()

	// 1 OMIT
	pipeReader, pipeWriter := io.Pipe() // HL
	// uploading to S3
	go func() {
		defer wg.Done()
		defer pipeReader.Close()                       // HL
		if err := uploadToS3(pipeReader); err != nil { // HL
			log.Fatal(err)
		}
	}()
	gz := gzip.NewWriter(pipeWriter) // HL
	io.Copy(gz, c.Body)

	pipeWriter.Close() // HL
	// END 1 OMIT
	gz.Flush()
	gz.Close()
	wg.Wait()
}
