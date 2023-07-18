package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	hellopb "sample/pkg/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
)

func main() {
	fmt.Println("start gRPC Client.")

	scanner = bufio.NewScanner(os.Stdin)

	conn, err := grpc.Dial(
		"localhost:8080",

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	grpcClient := hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello(grpcClient)

		case "2":
			fmt.Println("bye.")
			return
		}
	}
}

func Hello(grpcClient hellopb.GreetingServiceClient) {
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Please enter your age.")
	scanner.Scan()
	age, err := strconv.Atoi(scanner.Text())

	req := &hellopb.HelloRequest{
		Name: name,
		Age:  int32(age),
	}
	res, err := grpcClient.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
