package main

import (

	"fmt"
	"log"
	"net/http"
	"time"
)


var Port = ":5555"
func main(){

	http.HandleFunc("/",ServeFiles)
	fmt.Println("Serving @ : ","http://127.0.0.1"+Port)
	log.Fatal(http.ListenAndServe(Port,nil))
}

func ServeFiles(w http.ResponseWriter, r *http.Request){

	switch r.Method{

	case "GET":

		path := r.URL.Path

		fmt.Println(path)

		if path == "/"{

			path = "./static/index.html"
		}else{

			path = "."+path
		}

		http.ServeFile(w,r,path)

	case "POST":

		r.ParseMultipartForm(0)

		message := r.FormValue("message")

		fmt.Println("----------------------------------")
		fmt.Println("Message from Client: ",message)
		// respond to client's request
		fmt.Fprintf(w, "Server: %s \n", message+ " | " + time.Now().Format(time.RFC3339))
	
	default:
		fmt.Fprintf(w,"Request type other than GET or POSt not supported")


	}

}