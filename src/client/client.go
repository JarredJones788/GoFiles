package client

import (
	"os"
	"types"

	"github.com/minio/minio-go"
)

//FileClient - Structure of the client to access bucket
type FileClient struct {
	minio *minio.Client
}

//Create - returns an instance of minio.client
func (c *FileClient) Create(config *types.Config) error {
	client, err := minio.New(config.Endpoint, config.AccessID, config.Secret, config.UseSSL)
	if err != nil {
		return err
	}

	c.minio = client
	return nil
}

//UploadFile - uploades provided file to a bucket
func (c *FileClient) UploadFile(file *os.File, config *FileUpload) error {

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
func (c *FileClient) GetFile(config *FileGet) (*minio.Object, error) {

	object, err := c.minio.GetObject(config.BucketName, config.FileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return object, nil
}

//RemoveFile - delete a file from a bucket
func (c *FileClient) RemoveFile(config *FileRemove) error {

	err := c.minio.RemoveObject(config.BucketName, config.FileName)
	if err != nil {
		return err
	}

	return nil

}
