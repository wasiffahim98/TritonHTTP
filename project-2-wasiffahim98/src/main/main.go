package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-ini/ini"
	"github.com/ucsd-cse-124-fa21/project-2-wasiffahim98/tritonhttp"
)

type docRootEntry struct {
	Hostname string `json:"hostname"`
	Docroot  string `json:"docroot"`
}

// Server startup configuration constants
const ARG_COUNT int = 3
const CONFIG_FILE_INDEX int = 1
const HOSTS_FILE_INDEX int = 2

// Exit flags
const EX_USAGE int = 64
const EX_CONFIG int = 78

// Usage string
const USAGE_STRING string = "Usage: ./run-server [config_file] [hosts_file]"

// Configuration constant labels
const HTTPD string = "httpd"
const USE_DEFAULT_SERVER string = "use_default_server"
const SERVER_PORT string = "port"
const DOC_ROOT_PATH string = "doc_root"
const MIME_TYPE_PATH string = "mime_types"

func main() {
	var err error

	// When the input argument is less than 2
	if len(os.Args) != ARG_COUNT {
		log.Println(USAGE_STRING)
		os.Exit(EX_USAGE)
	}

	// Load the configuration file
	configFilePath := os.Args[CONFIG_FILE_INDEX]
	configContent, err := ini.Load(configFilePath)
	if err != nil {
		log.Println(err)
		log.Println("Failed to load config file:", configFilePath)
		log.Println(USAGE_STRING)
		os.Exit(EX_CONFIG)
	}

	// Load the detailed configuration from section "httpd"
	httpdConfigs := configContent.Section(HTTPD)
	useDefaultServer, _ := httpdConfigs.Key(USE_DEFAULT_SERVER).Bool()
	serverPort := httpdConfigs.Key(SERVER_PORT).String()
	DocRoot := parseDocRootJSON()
	mimeTypes := httpdConfigs.Key(MIME_TYPE_PATH).String()

	fmt.Println("Done loading configurations")

	// If useDefaultServer is true, start Go's in-built FileServer
	if useDefaultServer {
		docRoot := httpdConfigs.Key(DOC_ROOT_PATH).String()
		log.Println("Starting default server on port:", serverPort)
		log.Println("Server has doc root as:", docRoot)
		log.Println("Server has mime types file at:", mimeTypes)

		// Create a file server handler and let it handle all request start with an "/"
		http.Handle("/", http.FileServer(http.Dir(docRoot)))

		s := &http.Server{
			Addr:         ":" + serverPort,
			Handler:      http.FileServer(http.Dir(docRoot)),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}

		// Once there's an error in listen and serve, it will be caught by log Fatal and terminate the process
		log.Fatal(s.ListenAndServe())

	} else {
		log.Println("Starting TritonHTTP Server on port:", serverPort)

		// Initialize tritonhttp server
		httpdServer, err := tritonhttp.NewHttpdServer(":"+serverPort, DocRoot, mimeTypes)
		if err != nil {
			log.Fatal(err)
		}

		// Start tritonhttp server
		log.Fatal(httpdServer.Start())
	}
}

func parseDocRootJSON() map[string]string {

	docRoot := make(map[string]string, 2)
	docRootEntryList := make([]docRootEntry, 2)
	hostsFilePath := os.Args[HOSTS_FILE_INDEX]
	data, err := ioutil.ReadFile(hostsFilePath)
	if err != nil {
		log.Println(err)
		log.Println("Failed to hosts file:", hostsFilePath)
		log.Println(USAGE_STRING)
		os.Exit(EX_CONFIG)
	}
	err = json.Unmarshal(data, &docRootEntryList)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range docRootEntryList {
		absPath, err := filepath.Abs(entry.Docroot)
		if err != nil {
			log.Fatal(err)
		}
		docRoot[entry.Hostname] = absPath
	}
	return docRoot
}
