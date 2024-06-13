package jager

import (
	"fmt"
	"net/http"
	"testing"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Test_Server(t *testing.T) {
	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		Map(w, map[string]interface{}{
			"type":      "jager map",
			"name":      "John",
			"surname":   "Doe",
			"age":       30,
			"isStudent": true,
		})
	})
	http.HandleFunc("/string", func(w http.ResponseWriter, r *http.Request) {
		String(w, `{"type":"jager string","name":"john","surname":"doe","age":30,"isStudent":true}`)
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Write([]byte("Method Not Allowed, only POST method"))
			return
		}

		jsonData, err := Read(w, r)
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		Write(w, jsonData)
	})

	http.HandleFunc("/struct", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:  "John Doe",
			Age:   30,
			Email: "johndoe@example.com",
		}

		Struct(w, person)
	})

	fmt.Println("Jager Test Server Running on :5501")
	http.ListenAndServe(":5501", nil)
}
