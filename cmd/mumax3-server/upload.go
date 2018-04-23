package main

// Lets users upload simulation files.

import (
	"fmt"
	"log"
	"io"
	"os"
	"net/http"
	"regexp"
)

func HandleUpload(w http.ResponseWriter, req *http.Request) {
	RLock()
	defer RUnlock()

	isAlpha := regexp.MustCompile(`^[A-Za-z0-9.+_\s-]+$`).MatchString

	fmt.Println("method:", req.Method)
	if req.Method != "POST" {
		http.Error(w, "Does not compute", http.StatusNotFound)
		return; 
		}

	req.ParseMultipartForm(20*1024) //max 20k
	user := req.FormValue("user")
	_, ok := Users[user]	//is the form user one of the existing users?
	if !ok {
	//if !isAlpha(user) {  
		log.Println(	"upload: user" ,user, "is not valid.");
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	inFile, handler, err := req.FormFile("uploadfile")
	if err != nil { fmt.Println(err); return }
	defer inFile.Close()
	
	if !isAlpha(handler.Filename) {  
		log.Println(	"Upload not successful. filename",handler.Filename,"has unacceptable characters.")
		fmt.Fprintln(w,	"Upload not successful. filename",handler.Filename,"has unacceptable characters.")
		return
	}
	fextension := handler.Filename[len(handler.Filename)-4:];
	if fextension != ".mx3" {
		log.Println(	"Upload not successful. .mx3 is not:" , fextension)
		fmt.Fprintln(w,	"Upload not successful. .mx3 is not:" , fextension)
		return
	}

	fmt.Fprintln(w, handler.Header)
	outFile, err := os.OpenFile("./" +user+"/"+ handler.Filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	//os.Create("./test/" + handler.Filename)
	if err != nil { 
		log.Println(	"Upload not successful. Error creating file:",err)
		fmt.Fprintln(w,	"Upload not successful. Error creating file:",err)
		return }
	defer outFile.Close()

	written, err := io.Copy(outFile, inFile);
	if err != nil { 
		log.Println(	"Upload not successful. Error writing file:",err)
		fmt.Fprintln(w,	"Upload not successful. Error writing file:",err)
		return }
	log.Println(	"\nSuccessfully written uploaded file:", handler.Filename, "\nlength:", written)
	fmt.Fprintln(w,	"\nSuccessfully written uploaded file:", handler.Filename, "\nlength:", written)
}
