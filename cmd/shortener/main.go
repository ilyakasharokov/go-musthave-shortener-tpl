package main
import (
	"ilyakasharokov/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
