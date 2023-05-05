package tritonhttp

import (
	"log"
	"net"
	"strings"
	"time"
)

/*
For a connection, keep handling requests until
	1. a timeout occurs or
	2. client closes connection or
	3. client sends a bad request
*/
func (hs *HttpServer) handleConnection(conn net.Conn) {

	//panic("todo - handleConnection")

	// Start a loop for reading requests continuously

	// Set a timeout for read operation

	// Read from the connection socket into a buffer

	// Validate the request lines that were read

	// Handle any complete requests

	// Update any ongoing requests

	// If reusing read buffer, truncate it before next read
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")
	timeOut := 5 * time.Second
	readerRequest := make([]byte, 128)
	completeString := []string{}
	log.Println("Checking here")
	iteration := 0

	for {
		conn.SetReadDeadline(time.Now().Add(timeOut))
		data, err := conn.Read(readerRequest)
		log.Println("Read Info: ", string(readerRequest))

		if err != nil || data == 0 {
			break
		}
		log.Println("Data Info: ", data)
		completeString = append(completeString, string(readerRequest))
		log.Println("Complete String: ", completeString)
		iteration += 1
		log.Println("Counter: ", iteration)

	}

	// conn.SetReadDeadline(time.Now().Add(timeOut))
	// completeString = append(completeString, string(readerRequest))
	// log.Println("Complete String: ", completeString)

	newString := strings.Join(completeString, "")
	log.Println("newString: ", newString)
	newString2 := strings.Split(newString, "\r\n\r\n")
	log.Println("newString2: ", newString2)

	for i := 0; i < len(newString2)-1; i++ {
		hs.handleResponse(newString2[i], conn)
	}
}
