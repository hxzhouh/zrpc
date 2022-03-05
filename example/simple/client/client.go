package main

import (
	"context"
	"github.com/hxzhouh/zrpc/example/simple/hello"
	client2 "github.com/hxzhouh/zrpc/pkg/client"
	"log"
	"time"
)

func main() {

	helloClient := client2.NewClient("hello")
	helloClient.Init()
	conn, err := helloClient.Run()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := hello.NewHelloClient(conn)
	name := "nosixtools"

	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := client.SayHelloStream(ctx, &hello.HelloReq{Name: name})
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Hello: %s", r.Name)
		}
		time.Sleep(time.Second)
	}
}
