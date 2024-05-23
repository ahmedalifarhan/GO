package main

import (
	"log"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "192.168.0.35:8080")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	log.Println("Connected to server.")

	// Create a file to write the received data
	file, err := os.Create("Abdelrahman.txt")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	// Specify the number of bytes to receive
	bytesToReceive := 150  // Example: Receive 1024 bytes

	// Read data from connection and write it to file
	buffer := make([]byte, bytesToReceive)
	bytesReceived, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Error receiving data:", err)
	}

	// Write received bytes to file
	_, err = file.Write(buffer[:bytesReceived])
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	log.Println("File received successfully.")
}

