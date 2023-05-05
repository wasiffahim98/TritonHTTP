package tritonhttp

import (
	"log"
	"net"
)

/**
	Initialize the tritonhttp server by populating HttpServer structure
**/
func NewHttpdServer(port string, docRoot map[string]string, mimePath string) (*HttpServer, error) {
	//panic("todo - NewHttpdServer")

	// Initialize mimeMap for server to refer

	// Return pointer to HttpServer

	mimeMap, err := ParseMIME(mimePath)

	if err != nil {
		log.Panicln(err)
	}
	result := HttpServer{port, docRoot, mimePath, mimeMap}

	return &result, nil
}

/**
	Start the tritonhttp server
**/
func (hs *HttpServer) Start() (err error) {
	//panic("todo - StartServer")

	// Start listening to the server port

	// Accept connection from client

	// Spawn a go routine to handle request

	// port := flag.Int("port", strconv.Atoi(hs.ServerPort), "port to accept connections on.")
	// //host := flag.String("host", "127.0.0.1", "Host or IP to bind to")
	// flag.Parse()
	l, err := net.Listen("tcp", hs.ServerPort)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Listening to connections at '"+hs.ServerPort+"' on port", hs.ServerPort)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go hs.handleConnection(conn)
	}

}
