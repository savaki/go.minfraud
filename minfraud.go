package minfraud

import (
	"log"

	"github.com/savaki/httpctx"
	"golang.org/x/net/context"
)

const (
	MinfraudURL = "https://minfraud.maxmind.com/app/ccv2r"
)

type Client struct {
	HttpClient httpctx.HttpClient
	LicenseKey string
}

func New(licenseKey string) *Client {
	return &Client{
		HttpClient: httpctx.NewClient(),
		LicenseKey: licenseKey,
	}
}

func (c *Client) Do(query Query) (*QueryResult, error) {
	ctx := context.Background()
	return c.DoWithContext(ctx, query)
}

func (c *Client) DoWithContext(ctx context.Context, query Query) (*QueryResult, error) {
	params := query.Values()
	params.Add("license_key", c.LicenseKey)

	if query.Verbose {
		log.Println("minfraud.Do(...)")
		log.Printf("body => %s\n", params.Encode())
	}

	var data []byte
	err := c.HttpClient.Get(ctx, MinfraudURL, &params, &data)
	if err != nil {
		return nil, err
	}

	if query.Verbose {
		log.Printf("received => %s\n", string(data))
	}

	return ParseQueryResult(string(data))
}
