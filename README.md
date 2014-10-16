go.minfraud
===========

[![GoDoc](https://godoc.org/github.com/savaki/go.minfraud?status.svg)](https://godoc.org/github.com/savaki/go.minfraud) [![Build Status](https://snap-ci.com/savaki/go.minfraud/branch/master/build_image)](https://snap-ci.com/savaki/go.minfraud/branch/master)

golang implementation of maxmind's fraud library


## Example

```
package main

import (
	"encoding/json"
	"github.com/savaki/go.minfraud"
	"os"
)

func main() {
	licenseKey := "your-license-key"
	client := minfraud.New(licenseKey)

	// minfraud.Query supports for the full range of query options as defined by
	// http://dev.maxmind.com/minfraud/
	query := minfraud.Query{
		IpAddr: "1.2.3.4", // ip address to check
	}
	result, _ := client.Do(query)

	// dumps results to screen as json
	json.NewEncoder(os.Stdout).Encode(result)
}
```

## See Also

* [minFraud Web Service API](http://dev.maxmind.com/minfraud/)
