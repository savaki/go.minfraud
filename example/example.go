package main

import (
	"encoding/json"
	"github.com/savaki/go.minfraud"
	"os"
)

func main() {
	licenseKey := "your-license-key"
	client := minfraud.New(licenseKey)

	// execute the query
	query := minfraud.Query{
		IpAddr: "1.2.3.4", // ip address to check
	}
	result, _ := client.Do(query)

	// dumps results to screen as json
	json.NewEncoder(os.Stdout).Encode(result)
}
