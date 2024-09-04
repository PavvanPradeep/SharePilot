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
	serverAddr = "localhost:8080"
	bufferSize = 1024
)

type ProgressWriter struct {
	writer    io.Writer
	totalSize int64
	bytesSent int64
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.writer.Write(p)
	pw.bytesSent += int64(n)
	pw.printProgress()
	return n, err
}

func (pw *ProgressWriter) printProgress() {
	percentage := float64(pw.bytesSent) / float64(pw.totalSize) * 100
	fmt.Printf("\rProgress: %.2f%%", percentage)
}

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

	// Send file content with progress tracking
	progressWriter := &ProgressWriter{writer: conn, totalSize: size}
	_, err = io.Copy(progressWriter, file)
	if err != nil {
		log.Fatalf("Error sending file data: %s", err)
	}

	fmt.Println("\nFile sent successfully")
}
