package main

import (
	"fmt"
	"net/http"
	"os"
)

const healthURL = "http://127.0.0.1:8080/health"

func main() {
	if err := checkHealth(healthURL); err != nil {
		os.Exit(1)
	}
}

func checkHealth(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health check returned HTTP status %d", resp.StatusCode)
	}

	return nil
}