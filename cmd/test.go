package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var TestRestCmd = &cobra.Command{
	Use:   "test-ticket",
	Short: "start test-ticket",
	RunE:  testTicket,
}

func testTicket(cmd *cobra.Command, args []string) error {
	url := "http://localhost:5001/get-ticket"
	ticketID := "1001"
	userID1 := "1"
	userID2 := "3"

	client := &http.Client{Timeout: 10 * time.Second}

	var wg sync.WaitGroup
	wg.Add(2)

	makeRequest := func(userID string) {
		defer wg.Done()
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Request creation error:", err)
			return
		}

		req.Header.Set("ticket-id", ticketID)
		req.Header.Set("user-id", userID)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[%s] Request failed: %v\n", userID, err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("[%s] Response status: %s\n", userID, resp.Status)
	}

	go makeRequest(userID1)
	go makeRequest(userID2)

	wg.Wait()

	return nil
}
