package main

func main() {
	grpcServer := NewGRPCServer(":9090")
	grpcServer.Run()
}
