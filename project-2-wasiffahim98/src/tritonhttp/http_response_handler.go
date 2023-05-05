package tritonhttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path"
	"strconv"
	"strings"
	"time"
)

func (hs *HttpServer) handleBadRequest(conn net.Conn) {
	//panic("todo - handleBadRequest")
	errorHeader := "HTTP/1.1 400 Bad Request"
	errorDate := time.Now().String()
	errorLastModified := time.Now().String()
	errorCon := "close"
	crlf := "\r\n"
	erroResult :=
		errorHeader + crlf +
			"Date: " + errorDate + crlf +
			"Last-Modified: " + errorLastModified + crlf +
			"Connection: " + errorCon + crlf + crlf
	conn.Write([]byte(erroResult))
}

func (hs *HttpServer) handleFileNotFoundRequest(requestHeader string, conn net.Conn) {
	//panic("todo - handleFileNotFoundRequest")
	errorHeader := "HTTP/1.1 404 Not Found"
	errorDate := time.Now().String()
	errorLastModified := time.Now().String()
	errorCon := "close"
	crlf := "\r\n"
	erroResult :=
		errorHeader + crlf +
			"Date: " + errorDate + crlf +
			"Last-Modified: " + errorLastModified + crlf +
			"Connection: " + errorCon + crlf + crlf
	conn.Write([]byte(erroResult))
}

func (hs *HttpServer) handleResponse(request string, conn net.Conn) {
	modified := strings.Split(request, "\r\n")
	log.Println("Mofiefied: ", modified)
	requestBeginArr := modified[0] //GET <URL> HTTP/1.1
	url2 := strings.Split(requestBeginArr, " ")
	log.Println("URL2: ", len(url2), url2)

	actualURL := url2[1]

	//Create hash/dictionary
	tritonHash := make(map[string]string)

	//Loop through Modified
	for i := 1; i < len(modified); i++ {

		//Add in every key value pair into the hash/dictionary
		newVal := strings.SplitN(modified[i], ":", 2)
		newVal2 := strings.TrimSpace(newVal[1])
		tritonHash[newVal[0]] = newVal2
	}

	log.Println("Triton Hash Map: ", tritonHash)
	log.Println("DOC ROOT: ", hs.DocRoot)

	//get host through key "host"
	actualHost, presentActualHost := tritonHash["Host"]
	if !presentActualHost {
		hs.handleBadRequest(conn)
	}
	log.Println("The HOST is: ", actualHost)

	//get connection through key "connection"
	actualConnection, presentActualConnection := tritonHash["Connection"]
	if !presentActualConnection {
		actualConnection = "keep-alive"
	}
	log.Println("The CONNECTION is: ", actualConnection)

	//hs.DocRoot[] will return the HTML
	htmlPathOg, presenthtmlPathOg := hs.DocRoot[actualHost]
	if !presenthtmlPathOg {
		hs.handleBadRequest(conn)
	}
	log.Println("Original HTML File Path: ", htmlPathOg)

	//Once path given, append the URL
	filePath := path.Clean(htmlPathOg + actualURL)
	log.Println("Finalized File Path: ", filePath)

	htmlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		hs.handleFileNotFoundRequest(request, conn)
	}

	getType := strings.Split(actualURL, ".")
	setType := getType[len(getType)-1]
	finalType := "." + setType

	status := "HTTP/1.1 200 OK"
	crlf := "\r\n"
	date := time.Now()
	lastModified := time.Now()
	contentType := hs.MIMEMap[finalType]
	mainHtmlBody := string(htmlFile)
	contentLength := len(mainHtmlBody)
	connection := actualConnection

	result2 := status + crlf +
		"Date: " + date.String() + crlf +
		"Last-Modified: " + lastModified.String() + crlf +
		"Content-Type: " + contentType + crlf +
		"Content-Length: " + strconv.Itoa(contentLength) + crlf +
		connection + crlf +
		crlf +
		mainHtmlBody

	log.Println("Final CONN Write: ", result2)
	conn.Write(([]byte(result2)))

	// You get full path
	// Clean up the path
	// Make sure it doesnt access other file
	// Open file from full PATH, use ioutill GO Package -> used to open file
	// Find length of data and then assign it content-length
	// Get data from file
	// Figure out content-type from URL
	// Use hs.MimeMap[.html]

}

func (hs *HttpServer) sendResponse(responseHeader HttpResponseHeader, conn net.Conn) {
	panic("todo - sendResponse")

	// Send headers

	// Send file if required

	// Hint - Use the bufio package to write response
}
