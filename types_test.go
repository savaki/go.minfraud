package minfraud

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValues(t *testing.T) {
	Convey("Given a Query", t, func() {
		query := Query{
			IpAddr:      "1.2.3.4",
			LicenseKey:  "license-key",
			Email:       "joe.public@gmail.com",
			OrderAmount: 10.50,
		}

		Convey("When I retrieve the encoded values", func() {
			values := query.Values().Encode()

			Convey("Then I expect i=1.2.3.4", func() {
				So(values, ShouldContainSubstring, "i=1.2.3.4")
			})
		})
	})
}

func TestParseQueryResult(t *testing.T) {
	var result *QueryResult
	var err error

	Convey("Given a sample query result", t, func() {
		output := "riskScore=13.2;distance=6;countryMatch=Yes;countryCode=US;freeMail=Yes"

		Convey("When I call #ParseQueryResult", func() {
			result, err = ParseQueryResult(output)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect riskScore to be set", func() {
				So(result.RiskScore, ShouldEqual, 13.2)
			})

			Convey("And I expect distance to be set", func() {
				So(result.Distance, ShouldEqual, 6)
			})

			Convey("And I expect countryMatch to be set", func() {
				So(result.CountryMatch, ShouldEqual, "Yes")
			})

			Convey("And I expect countryCode to be set", func() {
				So(result.CountryCode, ShouldEqual, "US")
			})

			Convey("And I expect freeMail to be set", func() {
				So(result.FreeMail, ShouldEqual, "Yes")
			})
		})
	})
}
