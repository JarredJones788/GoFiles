package types

//FileData - struct to interact with api
type FileData struct {
	Key        string `json:"key"`
	BucketName string `json:"bucketName"`
	FileName   string `json:"fileName"`
	FileType   string `json:"fileType"`
	Duration   int64  `json:"duration"`
}
