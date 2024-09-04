package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	port       = ":8080"
	bufferSize = 1024
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	defer listener.Close()
	log.Printf("Server started on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Connection established with %s", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	// Read file name
	fileName, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading file name: %s", err)
		return
	}
	// Removes the newline character from the end of the file name
	fileName = fileName[:len(fileName)-1]

	// Read file size
	fileSizeStr, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading file size: %s", err)
		return
	}

	fileSize, err := strconv.ParseInt(fileSizeStr[:len(fileSizeStr)-1], 10, 64)
	if err != nil {
		log.Printf("Error parsing file size: %s", err)
		return
	}

	log.Printf("File name: %s", fileName)
	log.Printf("File size: %d bytes", fileSize)

	// Create a file to save the received data
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating file %s: %s", fileName, err)
		return
	}
	defer file.Close()

	buffer := make([]byte, bufferSize)
	var totalBytes int64

	// Read the file data from the client and write it to the file
	for totalBytes < fileSize {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading file data: %s", err)
			return
		}
		if bytesRead > 0 {
			_, err = file.Write(buffer[:bytesRead])
			if err != nil {
				log.Printf("Error writing to file: %s", err)
				return
			}
			totalBytes += int64(bytesRead)
			printProgress(totalBytes, fileSize)
		}
	}

	log.Printf("File received successfully, total bytes: %d", totalBytes)
}

func printProgress(bytesReceived, totalSize int64) {
	percentage := float64(bytesReceived) / float64(totalSize) * 100
	fmt.Printf("\rProgress: %.2f%%", percentage)
}
