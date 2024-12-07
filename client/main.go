package main

import (
	"context"
	"io"
	"log"
	"nicolascastro/go/grpc/testpb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	// DoUnary(c)

	// DoClientStreaming(c)

	DoServerStreaming(c)

}

func DoUnary(c testpb.TestServiceClient) {
	req := &testpb.GetTestRequest{
		Id: "t1",
	}

	res, err := c.GetTest(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GetTest: %v", err)
	}
	log.Printf("Response from GetTest: %v", res)

}

func DoClientStreaming(c testpb.TestServiceClient) {
	questions := []*testpb.Question{
		{
			Id:       "q8t1",
			Answer:   "Azul",
			Question: "Color asociado a Golang",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "Google",
			Question: "Empresa que desarrollo Golang",
			TestId:   "t1",
		},
		{
			Id:       "q10t1",
			Answer:   "Back-end",
			Question: "Tipo de lenguaje de programacion es Golang",
			TestId:   "t1",
		},
	}
	stream, err := c.SetQuestion(context.Background())
	if err != nil {
		log.Fatalf("error while calling SetQuestion: %v", err)
	}
	for _, question := range questions {
		log.Println("Sending question: ", question.Id)
		stream.Send(question)
		time.Sleep(2 * time.Second)

	}
	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from SetQuestion: %v", err)
	}
	log.Printf("Response from SetQuestion: %v", msg)
}

func DoServerStreaming(c testpb.TestServiceClient) {
	req := &testpb.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GetStudentsPerTest: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GetStudentsPerTest: %v", msg)

	}
}
