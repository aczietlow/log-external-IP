package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type IP struct {
	Query string
}

func main() {
	for i := 0; i < 3; i++ {
		//time.Sleep(15 * time.Minute)
		time.Sleep(1 * time.Second)
		fmt.Println(getIP())
		data := []string{time.Now().Format("2006-01-02 15:04:05"), getIP()}
		writeToFile("data.csv", data)
	}
}

func getIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func writeToFile(fileName string, value []string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error opening the file.")
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(value)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}
}
