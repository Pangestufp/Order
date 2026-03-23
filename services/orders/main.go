package main

func main() {
	grpcServer := NewGRPCServer(":9090")
	grpcServer.Run()

	httpServer := NewHttpServer(":8080")
	httpServer.Run()
}
