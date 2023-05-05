package tritonhttp

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/**
	Load and parse the mime.types file
**/
func ParseMIME(MIMEPath string) (MIMEMap map[string]string, err error) {
	//panic("todo - ParseMIME")

	file1, err := os.Open((MIMEPath))

	if err != nil {
		log.Panicln(err)
	}
	defer file1.Close()
	scanner1 := bufio.NewScanner(file1)
	mimeMap := make(map[string]string)
	for scanner1.Scan() {
		text1 := scanner1.Text()
		arr := strings.Split(text1, " ")
		i, j := arr[0], arr[1]
		mimeMap[i] = j
	}

	return mimeMap, err
}
