package main

import (
	"log"

	"github.com/google/uuid"
)

func main() {
	database, err := GetDatabase(MONGO)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Everything working fine\n")

	p := &Person{Id: uuid.NewString(), Name: "Danilo"}

	database.Save(p)

	fP, err := database.FindByName(p.Name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("P = %#v\n", fP)
}
