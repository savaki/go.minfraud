package minfraud

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	IpAddr         string
	City           string
	Region         string
	Postal         string
	Country        string
	ShipAddr       string
	ShipCity       string
	ShipRegion     string
	ShipPostal     string
	ShipCountry    string
	Domain         string
	Phone          string
	Email          string
	Username       string
	Bin            string
	BinName        string
	BinPhone       string
	SessionId      string
	UserAgent      string
	AcceptLanguage string
	TxnId          string
	OrderAmount    float32
	OrderCurrency  string
	ShopId         string
	TxnType        string
	AvsResult      string
	CvvResult      string
	RequestedType  string

	// Verbose logging
	Verbose bool
}

func (q Query) Values() url.Values {
	values := url.Values{}

	set := func(key, value string) {
		if value != "" {
			values.Set(key, value)
		}
	}

	setMD5 := func(key, value string) {
		if value != "" {
			hasher := md5.New()
			hasher.Write([]byte(value))
			hashed := hex.EncodeToString(hasher.Sum(nil))
			values.Set(key, hashed)
		}
	}

	setDecimal := func(key string, value float32) {
		if value != 0.0 {
			text := fmt.Sprintf("3.3f", value)
			values.Set(key, text)
		}
	}

	// Required Fields
	set("i", q.IpAddr)

	// Billing Address
	set("city", q.City)
	set("region", q.Region)
	set("postal", q.Postal)
	set("country", q.Country)

	// Shipping Address
	set("shipAddr", q.ShipAddr)
	set("shipCity", q.ShipCity)
	set("shopRegion", q.ShipRegion)
	set("shopPostal", q.ShipPostal)
	set("shipCountry", q.ShipCountry)

	// User Data
	set("domain", q.Domain)
	set("custPhone", q.Phone)
	setMD5("email", q.Email)
	setMD5("username", q.Username)

	// Bin Related
	set("bin", q.Bin)
	set("binName", q.BinName)
	set("binPhone", q.BinPhone)

	// Transaction Linking
	set("sessionID", q.SessionId)
	set("user_agent", q.UserAgent)
	set("accept_language", q.AcceptLanguage)

	// Transaction Information
	set("txnID", q.TxnId)
	setDecimal("order_amount", q.OrderAmount)
	set("order_current", q.OrderCurrency)
	set("shopID", q.ShopId)
	set("txn_type", q.TxnType)

	// Credit Card Check
	set("avs_result", q.AvsResult)
	set("cvv_result", q.CvvResult)

	// Miscellanious
	set("requested_type", q.RequestedType)

	return values
}

type QueryResult struct {
	// Risk Score
	RiskScore float32 `json:"riskScore,omitempty"`

	// GeoIP Location Checks
	CountryMatch     string  `json:"countryMatch,omitempty"`
	HighRiskCountry  string  `json:"highRiskCountry,omitempty"`
	Distance         int     `json:"distance,omitempty"`
	IpAccuracyRadius int     `json:"ip_accuracyRadius,omitempty"`
	IpCity           string  `json:"ip_city,omitempty"`
	IpRegion         string  `json:"ip_region,omitempty"`
	IpRegionName     string  `json:"ip_regionName,omitempty"`
	IpPostalCode     string  `json:"ip_postalCode,omitempty"`
	IpMetroCode      int     `json:"ip_metroCode,omitempty"`
	IpAreaCode       string  `json:"ip_areaCode,omitempty"`
	CountryCode      string  `json:"countryCode,omitempty"`
	IpCountryName    string  `json:"ip_countryName,omitempty"`
	IpContinentCode  string  `json:"ip_continentCode,omitempty"`
	IpLatitude       float32 `json:"ip_latitude,omitempty"`
	IpLongitude      float32 `json:"ip_longitude,omitempty"`
	IpTimeZone       string  `json:"ip_timeZone,omitempty"`
	IpAsnum          string  `json:"ip_asnum,omitempty"`
	IpUserType       string  `json:"ip_userType,omitempty"`
	IpNetSpeedCell   string  `json:"ip_netSpeedCell,omitempty"`
	IpDomain         string  `json:"ip_domain,omitempty"`
	IpIsp            string  `json:"ip_isp,omitempty"`
	IpOrg            string  `json:"ip_org,omitempty"`
	IpCityConf       float32 `json:"ip_cityConf,omitempty"`
	IpRegionConf     float32 `json:"ip_regionConf,omitempty"`
	IpPostalConf     float32 `json:"ip_postalConf,omitempty"`
	IpCountryConf    float32 `json:"ip_countryConf,omitempty"`

	// Proxy Detection
	AnonymousProxy   string  `json:"anonymousProxy,omitempty"`
	ProxyScore       float32 `json:"proxyScore,omitempty"`
	IpCorporateProxy string  `json:"ip_corporateProxy,omitempty"`

	// Emails and Login Checks
	FreeMail    string `json:"freeMail,omitempty"`
	CarderEmail string `json:"carderEmail,omitempty"`

	// Bank Checks
	BinMatch      string `json:"binMatch,omitempty"`
	BinCountry    string `json:"binCountry,omitempty"`
	BinNameMatch  string `json:"binNameMatch,omitempty"`
	BinName       string `json:"binName,omitempty"`
	BinPhoneMatch string `json:"binPhoneMatch,omitempty"`
	BinPhone      string `json:"binPhone,omitempty"`
	Prepaid       string `json:"prepaid,omitempty"`

	// Address and Phone Number Checks
	CustPhoneInBillingLoc string `json:"custPhoneInBillingLoc,omitempty"`
	ShipForward           string `json:"shipForward,omitempty"`
	CityPostalMatch       string `json:"CityPostalMatch,omitempty"`
	ShipCitPostalMatch    string `json:"ShipCityPostalMatch,omitempty"`

	// Account Information Fields
	QueriesRemaining int    `json:"queriesRemaining,omitempty"`
	MaxmindID        string `json:"maxmindID,omitempty"`
	MinfraudVersion  string `json:"minfraud_version,omitempty"`
	ServiceLevel     string `json:"service_level,omitempty"`

	// Error Reporting
	Err string `json:"err,omitempty"`
}

func ParseQueryResult(output string) (*QueryResult, error) {
	v := make(map[string]interface{})
	for _, section := range strings.Split(output, ";") {
		parts := strings.SplitN(section, "=", 2)
		key := parts[0]
		value := parts[1]

		if value == "" {
			continue
		}

		switch key {
		case "distance", "ip_accuracyRadius", "ip_metroCode", "queriesRemaining":
			if i, err := strconv.Atoi(value); err == nil {
				v[key] = i
			} else {
				return nil, err
			}
		case "riskScore", "ip_latitude", "ip_longitude", "ip_cityConf", "ip_regionConf", "ip_postalConf", "ip_countryConf", "proxyScore":
			if f, err := strconv.ParseFloat(value, 32); err == nil {
				v[key] = float32(f)
			} else {
				return nil, err
			}
		default:
			v[key] = value
		}
	}

	data, _ := json.Marshal(v)

	result := &QueryResult{}
	err := json.Unmarshal(data, result)
	return result, err
}
