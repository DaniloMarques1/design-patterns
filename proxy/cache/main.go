package main

import (
	"log"

	"github.com/danilomarques1/design_patterns/proxy/proxy"
	"github.com/danilomarques1/design_patterns/proxy/service"
)

func main() {
	us := proxy.NewUserServiceProxy()
	//us := service.NewUserService()
	us.Add(&service.User{
		Id:   "1",
		Name: "danilo",
	})
	us.Add(&service.User{
		Id:   "2",
		Name: "Fitz",
	})
	us.Add(&service.User{
		Id:   "3",
		Name: "Man",
	})

	u, err := us.FindById("3")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Name = %v\n", u.Name)

	u, err = us.FindById("3")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Name = %v\n", u.Name)

	u, err = us.FindById("3")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Name = %v\n", u.Name)
}
