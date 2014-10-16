package minfraud

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	MinfraudURL = "https://minfraud.maxmind.com/app/ccv2r"
)

func Do(query Query) (*QueryResult, error) {
	if query.Verbose {
		log.Println("minfraud.Do(...)")
		log.Printf("body => %s\n", query.Values().Encode())
	}

	body := query.Values().Encode()
	req, err := http.NewRequest("POST", MinfraudURL, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	// set additional headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if query.Verbose {
		log.Printf("received => %s\n", string(data))
	}

	return ParseQueryResult(string(data))
}
