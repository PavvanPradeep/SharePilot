package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

const (
	serverAddr = "localhost:8080"
	bufferSize = 1024
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run client.go <filename>")
	}
	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening file %s: %s", filepath, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %s", err)
	}
	size := stat.Size()
	log.Printf("File size: %d bytes", size)

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Error connecting to server: %s", err)
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)

	// Send file name followed by a newline
	_, err = writer.WriteString(filepath + "\n")
	if err != nil {
		log.Fatalf("Error sending file name: %s", err)
	}

	// Send file size followed by a newline
	_, err = writer.WriteString(strconv.FormatInt(size, 10) + "\n")
	if err != nil {
		log.Fatalf("Error sending file size: %s", err)
	}

	writer.Flush()

	// Send file content
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Fatalf("Error sending file data: %s", err)
	}

	log.Println("File sent successfully")
}
