# SharePilot
SharePilot is a robust and efficient file-sharing solution built using Go (Golang). Designed to simplify the process of transferring files over a network, SharePilot offers a seamless way to send and receive files between clients and a server. Its core functionality includes real-time progress tracking, error handling, and reliable data transmission.

Key Features:
- Effortless File Transfer: SharePilot facilitates the transfer of files from a client to a server using TCP connections. Simply specify the file you want to send, and SharePilot handles the rest.

- Error Handling: The application includes comprehensive error handling to manage issues such as connection failures, file read/write errors, and more. This ensures a smooth user experience and minimal disruptions.

- Server-Client Architecture: The server listens for incoming connections on port 8080, and the client connects to this server to initiate the file transfer. The server handles multiple connections concurrently, making it suitable for various use cases.

- Simple Command-Line Interface: The client is operated via a straightforward command-line interface. Just run the client with the file path as an argument, and SharePilot takes care of sending the file to the designated server.

Example Usage:
go run client.go <name_of_file>

go run server.go