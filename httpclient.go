package hackernews

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func get(url string, value interface{}) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	err = json.Unmarshal(body, value)
	if err != nil {
		fmt.Println("unmarshal body failure:", err)
		return err
	}
	return nil
}
