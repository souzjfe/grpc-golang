package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "path/to/event"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEventServiceClient(conn)

	event := &pb.Event{
		IdDispositivo:    123,
		NCanal:           456,
		ObjetoDetectado:  "objeto",
		CodRegraBurlada:  "codigo",
		Horario:          time.Now().Format(time.RFC3339),
		Reenvio:           0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for i := 0; i < 5; i++ {
		resp, err := client.EmitEvent(ctx, event)
		if err == nil {
			log.Printf("Event received by server: %v", event)
			break
		}

		// Se ocorrer um erro, incrementa o contador de reenvios e tenta enviar novamente.
		event.Reenvio++
		log.Printf("Error sending event: %v. Retrying in %v seconds...", err, rand.Intn(5)+1)
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}

	if event.Reenvio == 5 {
		log.Printf("Failed to send event after 5 retries: %v", event)
	}
}