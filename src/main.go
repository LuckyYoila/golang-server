package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	Message string "json:message"
	Name string "json:name"
}



func main(){

	Server := http.Server{
		Addr: ":8080",
		Handler: nil,
	}

	http.HandleFunc("/api", func (w http.ResponseWriter, r *http.Request) {
		message := APIResponse{
			 "Hello World",
			 "John Doe",
		}
		
		output , err := json.Marshal(message)

		println(string(output))

		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(output)
		// fmt.Fprintf(w, string(output))

	})
	print("Listening on port 8080 \n")
	running:= Server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(running)
	if running != nil {
		fmt.Println("Server is running")
		return
	}
}