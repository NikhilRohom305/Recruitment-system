package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func UploadResume(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")

	err := req.ParseMultipartForm(30)
	if err != nil {
		return
	}
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		return
	}
	if !strings.HasSuffix(fileHeader.Filename, ".pdf") && !strings.HasSuffix(fileHeader.Filename, ".docx") {
		http.Error(w, "file should be either pdf or docx", 501)
		return
	}

	dst, err := os.Create(fileHeader.Filename)
	if err != nil {
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		return
	}
	err = ParseResume(dst.Name()) // Pass the temporary file path
	if err != nil {
		http.Error(w, "Error parsing resume file", http.StatusInternalServerError)
		return
	}

	resp := make(map[string]interface{})
	resp["message"] = "file uploaded successfully"
	res, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(res)
}

func ParseResume(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	// Read the file content
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	reqBody := bytes.NewReader(fileContent)
	apiURL := "https://api.apilayer.com/resume_parser/upload"
	apiKey := "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1"

	req, err := http.NewRequest("POST", apiURL, reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	log.Print("statuscode", resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Print(string(respBody))
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(respBody, &parsedResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println(parsedResponse)
	return
}
