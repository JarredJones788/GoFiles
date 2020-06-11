package main

import (
	"client"
	"fmt"
	"io/ioutil"
	"types"
)

func main() {

	//Setup config
	config := &types.Config{
		Endpoint: "nyc3.digitaloceanspaces.com",
		AccessID: "B7DGWCR5E3MOXTQANMRX",
		Secret:   "pKnaWLIcGvNiRe5hCa9y8dCx6M/dXap96w8fd4nT/1k",
		UseSSL:   true,
	}

	//Create instance
	uploadClient := client.FileClient{}

	//Try and connect to our file store
	err := uploadClient.Create(config)

	if err != nil {
		fmt.Println(err)
		return
	}

	//--------------DELETE FILE-------------

	// removeConfig := &client.FileRemove{
	// 	BucketName: "tester788",
	// 	FileName:   "test2/server.go",
	// }

	// err = uploadClient.RemoveFile(removeConfig)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Removed File")

	//--------------GET FILE----------------

	getConfig := &client.FileGet{
		BucketName: "tester788",
		FileName:   "test2/server.go",
	}

	object, err := uploadClient.GetFile(getConfig)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileBytes, err := ioutil.ReadAll(object)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fileBytes)

	//-----------UPLOAD FILE--------------------

	// file, err := os.Open("./server.go")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()

	// uploadConfig := &client.FileUpload{
	// 	BucketName: "tester788",
	// 	FileName:   "test2/server.go",
	// 	FileType:   "txt",
	// }

	// err = uploadClient.UploadFile(file, uploadConfig)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("File Uploaded!")

}
