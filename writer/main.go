package main

import (
	userpb "go-grpc-simple/gen/go/user/v1"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	u := userpb.User{
		Uuid: "123",
		Fullname: "John Doe",
		BirthYear: 1990,
	}

	b, err := proto.Marshal(&u)
	if err != nil {
		log.Fatalln("Failed to encode user:", err)
	}

	if err := os.WriteFile("user.bin", b, 0666); err != nil {
		log.Fatalln("Failed to write user:", err)
	}

}