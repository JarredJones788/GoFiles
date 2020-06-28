package client

import (
	"config"
	"net/url"
	"os"
	"time"
	"types"

	"github.com/minio/minio-go"
)

//FileClient - Structure of the client to access bucket
type FileClient struct {
	minio *minio.Client
}

//Create - returns an instance of minio.client
func (c *FileClient) Create(config *config.Config) error {
	client, err := minio.New(config.S3Config.Endpoint, config.S3Config.AccessID, config.S3Config.Secret, config.S3Config.UseSSL)
	if err != nil {
		return err
	}
	c.minio = client

	return nil
}

//PresignedUploadFile - provides a valid link to upload a file with expiry
func (c *FileClient) PresignedUploadFile(config *types.FileData) (*url.URL, error) {

	//If duration is not set then default to 60 seconds
	if config.Duration == 0 {
		config.Duration = 60
	}

	url, err := c.minio.PresignedPutObject(config.BucketName, config.FileName, time.Duration(config.Duration)*time.Second)
	if err != nil {
		return nil, err
	}

	return url, nil
}

//PresignedGetFile - provides a valid link to a file with an expiry
func (c *FileClient) PresignedGetFile(config *types.FileData) (*url.URL, error) {

	//If duration is not set then default to 60 seconds
	if config.Duration == 0 {
		config.Duration = 60
	}

	url, err := c.minio.PresignedGetObject(config.BucketName, config.FileName, time.Duration(config.Duration)*time.Second, nil)
	if err != nil {
		return nil, err
	}

	return url, nil
}

//UploadFile - uploades provided file to a bucket
func (c *FileClient) UploadFile(file *os.File, config *types.FileData) error {

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = c.minio.PutObject(config.BucketName, config.FileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: config.FileType})
	if err != nil {
		return err
	}

	return nil
}

//GetFile - gets a file from a bucket
func (c *FileClient) GetFile(config *types.FileData) (*minio.Object, error) {

	object, err := c.minio.GetObject(config.BucketName, config.FileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return object, nil
}

//RemoveFile - delete a file from a bucket
func (c *FileClient) RemoveFile(config *types.FileData) error {

	err := c.minio.RemoveObject(config.BucketName, config.FileName)
	if err != nil {
		return err
	}

	return nil

}
