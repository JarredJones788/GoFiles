package router

import (
	"client"
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"types"

	"github.com/gorilla/mux"
)

//Router type
type Router struct {
	Host       string
	FileClient *client.FileClient
	FileKey    string
}

//Init - inits all routes.
func (router Router) Init(config *config.Config) error {

	router.Host = config.Host

	router.FileClient = &client.FileClient{}

	//Try and connect to our file store
	err := router.FileClient.Create(config)

	if err != nil {
		return err
	}

	//Get file secret.
	data, err := ioutil.ReadFile(config.Tokens.FileKey)
	if err != nil {
		return err
	}

	//Set key from the file
	router.FileKey = string(data)

	//Setup mux router
	r := mux.NewRouter()
	router.setUpRoutes(r)
	fmt.Println("Server Started")
	http.ListenAndServe(config.Port, r)

	return nil
}

//setUpRoutes - sets up all endpoints for the service
func (router Router) setUpRoutes(r *mux.Router) {
	r.HandleFunc("/api/file/upload", router.upload)
	r.HandleFunc("/api/file/get", router.getFile)
	r.HandleFunc("/api/file/remove", router.removeFile)
}

//-----------------HELPERS BELOW-----------------\\

//badRequest - returns a generic bad response
func (router Router) badRequest(w http.ResponseWriter) {
	failed, err := json.Marshal(types.GenericResponse{Response: false})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(failed)
}

//goodRequest - returns a generic good response
func (router Router) goodRequest(w http.ResponseWriter) {
	good, err := json.Marshal(types.GenericResponse{Response: true})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(good)
}

//reasonRequest - returns a response with a reason
func (router Router) reasonResponse(w http.ResponseWriter, response bool, reason string) {
	good, err := json.Marshal(types.ReasonResponse{Response: response, Reason: reason})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(good)
}

//urlRequrest - returns a response with a url
func (router Router) urlResponse(w http.ResponseWriter, response bool, url string) {
	good, err := json.Marshal(types.URLResponse{Response: response, URL: url})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(good)
}

//setUpHeaders - sets the desired headers for an http response
func (router Router) setUpHeaders(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", router.Host)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Max-Age", "120")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return false
	}
	return true
}

//-----------------ROUTES BELOW-----------------\\

//upload - endpoint to get a link to upload a new file
func (router Router) upload(w http.ResponseWriter, r *http.Request) {
	if !router.setUpHeaders(w, r) {
		return //request was an OPTIONS which was handled.
	}

	var data types.FileData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("Upload Error: " + err.Error())
		router.badRequest(w)
		return
	}

	//Key provided does not match the private key
	if data.Key != router.FileKey {
		fmt.Println("File key does not match")
		router.badRequest(w)
		return
	}

	//Info below is needed to contiue
	if data.BucketName == "" || data.FileName == "" {
		fmt.Println("Bucket name or file name is blank")
		router.reasonResponse(w, false, "Bucket name or file name is blank")
		return
	}

	//Get the generated upload url
	url, err := router.FileClient.PresignedUploadFile(&data)
	if err != nil {
		fmt.Println(err)
		router.badRequest(w)
		return
	}

	router.urlResponse(w, true, url.String())
}

//getFile - endpoint to get a link to a file
func (router Router) getFile(w http.ResponseWriter, r *http.Request) {
	if !router.setUpHeaders(w, r) {
		return //request was an OPTIONS which was handled.
	}

	var data types.FileData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("Upload Error: " + err.Error())
		router.badRequest(w)
		return
	}

	//Key provided does not match the private key
	if router.FileKey != data.Key {
		fmt.Println("File key does not match")
		router.badRequest(w)
		return
	}

	//Info below is needed to contiue
	if data.BucketName == "" || data.FileName == "" {
		fmt.Println("Bucket name or file name is blank")
		router.reasonResponse(w, false, "Bucket name or file name is blank")
		return
	}

	//Get the generated upload url
	url, err := router.FileClient.PresignedGetFile(&data)
	if err != nil {
		fmt.Println(err)
		router.badRequest(w)
		return
	}

	router.urlResponse(w, true, url.String())
}

//removeFile - endpoint to remove a file
func (router Router) removeFile(w http.ResponseWriter, r *http.Request) {
	if !router.setUpHeaders(w, r) {
		return //request was an OPTIONS which was handled.
	}

	var data types.FileData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("Upload Error: " + err.Error())
		router.badRequest(w)
		return
	}

	//Key provided does not match the private key
	if data.Key != router.FileKey {
		fmt.Println("File key does not match")
		router.badRequest(w)
		return
	}

	//Info below is needed to contiue
	if data.BucketName == "" || data.FileName == "" {
		fmt.Println("Bucket name or file name is blank")
		router.reasonResponse(w, false, "Bucket name or file name is blank")
		return
	}

	//Get the generated upload url
	err := router.FileClient.RemoveFile(&data)
	if err != nil {
		fmt.Println(err)
		router.badRequest(w)
		return
	}

	router.goodRequest(w)
}
