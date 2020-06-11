package client

//FileUpload - struct to upload a file
type FileUpload struct {
	BucketName string
	FileName   string
	FileType   string
}

//FileGet - struct to get a file
type FileGet struct {
	BucketName string
	FileName   string
}

//FileRemove - struct to remove a file
type FileRemove struct {
	BucketName string
	FileName   string
}
