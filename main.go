package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string;
}

func main () {

	p1 := person {
		First: "Jenny",
	}

	p2 := person {
		First: "James",
	}

	xp := []person{p1, p2}

 bs, err := json.Marshal(xp)
 if err != nil {
	 log.Panic(err)
 }
 fmt.Println("PRINT JSON", string(bs))

 xp2 := [] person {}

 err = json.Unmarshal(bs, &xp2)
 if (err != nil ) {
	 log.Panic(err)
 }

 fmt.Println("back into a go data structure", xp2)

 http.HandleFunc("/encode", foo)
 http.HandleFunc("/decode", bar)

 http.HandleFunc("/encodel", encode)
 http.HandleFunc("/decodel", decode)
 http.ListenAndServe(":8080", nil)

}

func encode ( w http.ResponseWriter, r *http.Request) {
	p1 := person {
		First: "Jenny",
	}

	p2 := person {
		First: "James",
	}

	xp := []person{p1, p2}

	err := json.NewEncoder(w).Encode(xp)
	 if err != nil {
		 log.Println("Error in encode ", err)
	 }

}

func decode( w http.ResponseWriter, r *http.Request) {
	var xp []person
	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Person list received", xp)
	
}

func foo( w http.ResponseWriter, r *http.Request) {
   p1 := person {
		 First: "Jenny",
	 }

	 err := json.NewEncoder(w).Encode(p1)
	 if err != nil {
		 log.Println("Error in foo ", err)
	 }

}


func bar( w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Person received", p1)
	
}