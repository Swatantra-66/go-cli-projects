package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   uint   `json:"phone"`
	LoginId uint   `json:"loginId" validate:"required"`
	Salary  uint   `json:"salary"`
	JobRole string `json:"jobRole"`
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull\n")
	name := r.FormValue("name")
	loginId := r.FormValue("loginId")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "LoginId = %v\n", loginId)
}

func main() {

	user := User{
		Name:    "Swatantra Yadav",
		Email:   "maverickswatantra@gmail.com",
		Phone:   9151356121,
		LoginId: 9151,
		Salary:  80000,
		JobRole: "DevOps Enginner",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/encoding", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/form", formHandler)

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method is not supported", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "Hello from Go Server")
	})

	fmt.Println("Server running on http://localhost:3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
