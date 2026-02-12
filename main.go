package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {


	/*
	200 OK 
	201 CREATED

	300 Redirect

	400 Error 

	500 Server error
	
	*/


	//--- GET

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		
		if request.Method != http.MethodGet {
			http.Error(writer, "Endast GET Stöds på /", http.StatusMethodNotAllowed)
			return 
		}

		writer.Header().Set("Content-type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(writer, "Hej! Det här är min server som körs på port 8080")
	})

	//--- POST

	http.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {

		
		type Request struct {
			Msg string `json:"msg"`
		}

		if request.Method != http.MethodPost {
			http.Error(writer, "Bara POST method tillåten", http.StatusMethodNotAllowed)
			return
		}

		var req Request

		err := json.NewDecoder(request.Body).Decode(&req)


		
		if err != nil {
			http.Error(writer, "Fel JSON format", http.StatusBadRequest)
			return
		}

		
		writer.Header().Set("Content-Type", "text/plain")

		fmt.Fprintf(writer, "Svar från POST endpoint, du skickade: %s", req.Msg)

		/*


		// http://localhost:8080/post?msg=Yahya

		value := request.FormValue("msg")

		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusCreated)
		fmt.Fprintf(writer, "Svar från POST endpoint, du skickade: %s", value)

		*/

		})

		http.HandleFunc("/put", func(writer http.ResponseWriter, request *http.Request) {

			if request.Method != http.MethodPut {
				http.Error(writer, "Bara PUT Method tillåtet", http.StatusMethodNotAllowed)
				return
			}

			value := request.FormValue("msg")

			writer.Header().Set("Content-Type", "text/plain")
			writer.WriteHeader(http.StatusOK)

			fmt.Fprintf(writer, "PUT endpointen. du uppdaterade msg till: %s", value)

		})

		http.HandleFunc("/delete", func(writer http.ResponseWriter, request *http.Request) {
			
			if request.Method != http.MethodDelete {
				http.Error(writer, "Bara Delete Method tillåtet", http.StatusMethodNotAllowed)
				return
			}


			msg := request.FormValue("msg")
			writer.Header().Set("Content-Type", "text/plain")
			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, "Delete endpointen, du raderade: %s", msg)

		})


	port := "8080"

	log.Println("Servern startat på http://localhost:" + port)

	err := http.ListenAndServe(":"+port,nil)


	if err != nil {
		log.Fatal(err)
	}
	
}

// Våran app kommer att köras på port 8080
// http://localhost:8080