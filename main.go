package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response[T any] struct {
	Success bool `json:"success"`
	Data T `json:"data"`
}

type PostPasteData struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Content string `json:"content"`
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeNamedPipe) == 0 {
		log.Println("Input must be through a pipe")
		return
	}

	instance := flag.String("instance", "https://zer0b.in", "The instance to use")
	markdown := flag.Bool("md", false, "Markdown mode")

	flag.Parse()

	buf, _ := ioutil.ReadAll(os.Stdin)
	content := string(buf)

	if *markdown {
		content = "md " + content
	}

	resp, err := Post(*instance, content)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response Response[PostPasteData]
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !response.Success {
		fmt.Println(response.Data.Message)
		os.Exit(1)
	}

	url := fmt.Sprintf("%s/%s", *instance, response.Data.Id)

	fmt.Println(url)
}

func Post(instance string, content string) (*http.Response, error) {
	values := map[string]string{"content": content}
	payload, err := json.Marshal(values)

	if err != nil {
		return nil, err
	}
	
	return http.Post(instance + "/api/p/n", "application/json",
		bytes.NewBuffer(payload))
}