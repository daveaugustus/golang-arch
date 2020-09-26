package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	// p1 := person{
	// 	First: "Dave",
	// }

	// p2 := person{
	// 	First: "James",
	// }

	// xp := []person{p1, p2}

	// fmt.Println("Marshaling")
	// jsonbyte, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(string(jsonbyte))

	// // Unmarshaling the same

	// fmt.Println("Unmarshaling")
	// xp2 := []person{}
	// if err := json.Unmarshal(jsonbyte, &xp2); err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(xp2)

	// Server and services

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Dave",
	}

	p2 := person{
		First: "James",
	}

	xs := []person{p1, p2}

	if err := json.NewEncoder(w).Encode(xs); err != nil {
		log.Println("Encode bad data!", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person

	if err := json.NewDecoder(r.Body).Decode(&p1); err != nil {
		log.Println("Decode bad data!", err)
	}

	log.Println("Person: ", p1)

}
