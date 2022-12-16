package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Person struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsWorking bool   `json:"is_working"`
}

func SendKafka(reader io.Reader) {
	b, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sending to kafka... %v\n", string(b))
}

type PersonAdapter struct {
	p Person
}

func (pd PersonAdapter) Read(p []byte) (int, error) {
	b, err := json.Marshal(pd.p)
	if err != nil {
		return 0, err
	}
	copy(p, b)
	return len(p), io.EOF
}

func main() {
	p := Person{
		Name: "Danilo Marques",
		Email: "danilomarques.cristo@gmail.com",
		IsWorking: true,
	}

	pd := PersonAdapter{p}

	f, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}

	SendKafka(f)
	SendKafka(pd)
}
