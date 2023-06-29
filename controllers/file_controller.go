package controllers

import (
	"go-http-server/responses"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func UploadHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Error Retrieving the File"}})
	}

	//verify file type
	if file.Header.Get("Content-Type") != "video/mp4" {
		log.Println("Invalid file type uploaded")
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Invalid file type uploaded"}})
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-files", "upload-*.mp4")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()

	tempFileName := tempFile.Name()

	// This will copy the uploaded file to the temp file created
	fileContent, err := file.Open()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(fileContent)

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": "Error Retrieving the File"}})
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
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Successfully Uploaded File"}})
}
