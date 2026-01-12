package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"unicode"
	//fyne
)

type PasswordResponse struct {
	Password         string `json:"password"`
	Length           int    `json:"length"`
	PasswordStrength string `json:"passwordStrength"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()
	mux.HandleFunc("/password", PasswordGenerator)

	fmt.Println("Server running on http://localhost:1010")
	err := http.ListenAndServe(":1010", mux)
	if err != nil {
		log.Fatal("Server Error: 404 not found")
	}
}

func CheckStrength(password string) string {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true

		case unicode.IsLower(char):
			hasLower = true

		case unicode.IsDigit(char):
			hasNumber = true

		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	score := 0
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSpecial {
		score++
	}

	switch {
	case score <= 2:
		return "Weak"
	case score < 4:
		return "Strong"
	default:
		return "Very Strong"
	}
}

func PasswordGenerator(w http.ResponseWriter, r *http.Request) {
	lowCase := "abcdefghijklmnopqrstuvwxyz"
	uppCase := "ABCDEFGHIJKLMOPQRSTUVWXYZ"
	Numbers := "0123456789"
	SpecialChar := "$&()*+[]@#_!?"
	passwordLength := 8
	password := ""

	for n := 0; n < passwordLength; n++ {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(4)
		// fmt.Println(randNum)

		switch randNum {
		case 0:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(lowCase))
			password = password + string(lowCase[randNum])

		case 1:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(uppCase))
			password = password + string(uppCase[randNum])

		case 2:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(Numbers))
			password = password + string(Numbers[randNum])

		case 3:
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(len(SpecialChar))
			password = password + string(SpecialChar[randNum])
		}
	}

	strength := CheckStrength(password)

	response := PasswordResponse{
		Password:         password,
		Length:           len(password),
		PasswordStrength: strength,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
