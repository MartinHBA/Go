package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func readMyToken() string {
	// Open the file.
	f, _ := os.Open("token.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	var lineNumber int
	for scanner.Scan() {
		if lineNumber == 0 {
			return scanner.Text()
		}
		lineNumber++
	}
	return ""
}

func translateToSK(mystring string) string {
	// Set up the request to the DeepL API.
	url := "https://api-free.deepl.com/v2/translate"

	// Generate the JSON payload using json.Marshal.
	payload, err := json.Marshal(map[string]interface{}{
		"text":        []string{mystring},
		"target_lang": "en",
		"source_lang": "sk",
	})
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	token := readMyToken()
	req.Header.Set("Authorization", token)

	// Send the request and get the response.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}

	// Parse the JSON response.
	var data map[string][]map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}

	// Access the translated text from the response.
	translations := data["translations"]
	if len(translations) == 0 {
		fmt.Println("No translations found in response.")
		return "No translations found in response"
	}
	translatedText := translations[0]["text"]
	return translatedText
}
