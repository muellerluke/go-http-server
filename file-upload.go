package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func VideoUploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(32 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")

	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}

	//verify file type
	fileTypeAcceptable := handler.Header.Get("Content-Type") == "video/mp4"

	if !fileTypeAcceptable {
		log.Println("File type not acceptable")
		return
	}

	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-files", "upload-*.mp4")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()

	tempFileName := tempFile.Name()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	err1 := ffmpeg.Input("./"+tempFileName).Output("./output-files/out1.mp4", ffmpeg.KwArgs{"c:v": "libx264", "vf": "scale=-2:720"}).OverWriteOutput().Run()
	ffmpeg.Input("./"+tempFileName).Output("./output-files/out2.mp4", ffmpeg.KwArgs{"c:v": "libx264", "vf": "scale=-2:1080"}).OverWriteOutput().Run()
	ffmpeg.Input("./"+tempFileName).Output("./output-files/out3.mp4", ffmpeg.KwArgs{"c:v": "libx264", "vf": "scale=-2:1440"}).OverWriteOutput().Run()
	ffmpeg.Input("./"+tempFileName).Output("./output-files/out4.mp4", ffmpeg.KwArgs{"c:v": "libx264", "vf": "scale=-2:2160"}).OverWriteOutput().Run()

	// delete temp file
	os.Remove(tempFileName)

	if err1 != nil {
		log.Println(err1)
	}

	// return that we have successfully uploaded our file!
	log.Println("Successfully Uploaded File")
}
