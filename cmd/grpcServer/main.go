package main

import (
	"database/sql"
	"net"

	"github.com/vinicius3g/gRPC.git/internal/database"
	"github.com/vinicius3g/gRPC.git/internal/pb"
	"github.com/vinicius3g/gRPC.git/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// conecta o banco d dados
	db, err := sql.Open("sqlite3", "./db.sqlite")
	// verifica se ha erros
	if err != nil {
		panic(err)
	}
	// fecha a conexão
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	// cria o servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	// abrir conexão tcp
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
