package main

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// setup minio server info
	endpoint := "192.168.200.20:9000"
	accessKeyID := "1Y1P8ZMW7A4X8SXRD1DD"
	secretAccessKey := "dhAUR7JKmdKcdFAcn6grRxnO7MuKZy8moMc+C9V+"
	userSSL := false

	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, userSSL)
	if err != nil {
		log.Fatal(err)
	}

	bucketName := "mymusice"
	location := "us-east-1"

	// create bucket name mysusice
	err = client.MakeBucket(bucketName, location)
	if err != nil {
		exsites, err := client.BucketExists(bucketName)
		if err == nil && exsites {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}

	log.Printf("successffully created %s\n", bucketName)

	// upload the zip file
	objectName := "test.zip"
	filePath, err := filepath.Abs(filepath.Dir(os.Args[0])
	fmt.Println(filePath)
	contentType := "application/zip"

	n, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successful upload %s of size %d\n", objectName, n)
}