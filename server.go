package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "pb"
)

type eventServer struct{}

func (s *eventServer) EmitEvent(ctx context.Context, event *pb.Event) (*pb.Empty, error) {
    log.Printf("Event received: id_dispositivo=%d, n_canal=%d, objeto_detectado=%s, cod_regra_burlada=%s, horario=%s",
        event.IdDispositivo, event.NCanal, event.ObjetoDetectado, event.CodRegraBurlada, event.Horario)

    // Implemente aqui a lógica para lidar com o evento recebido.

    return &pb.Empty{}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterEventServiceServer(s, &eventServer{})

    log.Println("Starting gRPC server on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}