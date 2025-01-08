package main

import (
	"context"
	"log"
	"time"

	"github.com/Edu58/Learn-Go/GRPCTutorial/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := invoicer.NewInvoicerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := c.Create(ctx, &invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:  345345,
			Curreny: "KES",
		},
		From: "edwin",
		To:   "Jeff Bezos",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("RECEIVED RESPONSE pdf: %s and docx: %s", string(response.GetPdf()), string(response.GetDocx()))
}
