package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type API struct {
	message string 
	name string 
}



func main(){

	Server := http.Server{
		Addr: ":8080",
		Handler: nil,
	}

	http.HandleFunc("/api", func (w http.ResponseWriter, r *http.Request) {
		message := API{
			 "Hello World",
			 "John Doe",
		}
		
		output , err := json.Marshal(message)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(w, string(output))

	})
	print("Listening on port 8080")
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