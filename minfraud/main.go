package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/savaki/go.minfraud"
	"log"
	"os"
)

const (
	ipFlag         = "ip"
	licenseKeyFlag = "key"
	verboseFlag    = "verbose"
)

func main() {
	app := cli.NewApp()
	app.Name = "minfraud"
	app.Usage = "command line interface to minfraud service"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{licenseKeyFlag, "", "the license key", "MINFRAUD_LICENSE_KEY"},
		cli.StringFlag{ipFlag, "", "the ip to check against", ""},
		cli.BoolFlag{verboseFlag, "verbose logging", ""},
	}
	app.Action = Run
	app.Run(os.Args)
}

func Run(c *cli.Context) {
	license := c.String(licenseKeyFlag)
	ip := c.String(ipFlag)
	verbose := c.Bool(verboseFlag)

	query := minfraud.Query{
		IpAddr:  ip,
		Verbose: verbose,
	}

	result, err := minfraud.New(license).Do(query)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
