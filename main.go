package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Dave",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}

	jsonbyte, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(jsonbyte))
}
