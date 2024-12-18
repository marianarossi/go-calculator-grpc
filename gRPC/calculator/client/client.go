package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ibilalkayy/Small-Projects/gRPC/calculator/proto" //biblioteca de calculadora do GO
	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:50051"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	for {
		fmt.Println("\nSelect an operation:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		var choice string
		fmt.Scanln(&choice)

		if choice == "5" {
			fmt.Println("Exiting...")
			break
		}

		var num1, num2 float64

		fmt.Print("Enter first number: ")
		_, err := fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		fmt.Print("Enter second number: ")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		var result *pb.Response

		switch choice {
		case "1":
			result, err = client.Add(context.Background(), &pb.AddRequest{Num1: num1, Num2: num2})
		case "2":
			result, err = client.Subtract(context.Background(), &pb.SubtractRequest{Num1: num1, Num2: num2})
		case "3":
			result, err = client.Multiply(context.Background(), &pb.MultiplyRequest{Num1: num1, Num2: num2})
		case "4":
			result, err = client.Divide(context.Background(), &pb.DivideRequest{Num1: num1, Num2: num2})
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Result: %.2f\n", result.Response)
		}
	}
}
