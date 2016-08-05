package cloudinary

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var domain string = "https://api.cloudinary.com/v1_1/"

// Cloudinary main struct
type Cloudinary struct {
	publicKey string
	secretKey string
	name      string
	urls      Option
}

// Option is the optional parameters custom struct
type Option map[string]string

// ErrorResp is the failed api request main struct
type ErrorResp struct {
	Message string `json:"message"`
}

// Create is creating a new cloudinary instance
func Create(public string, secret string, name string) *Cloudinary {
	return &Cloudinary{
		publicKey: public,
		secretKey: secret,
		name:      name,
		urls: Option{
			"upload": domain + name + "/image/upload",
		},
	}
}

func (c *Cloudinary) checkOptionsAreValid(options Option, validator []string) {
	for key := range options {
		if !validOption(key, validator) {
			panic("Upload paramater not valid: " + key)
		}
	}
}

var validOption = func(optionName string, val []string) bool {
	for _, param := range val {
		if param == optionName {
			return true
		}
	}
	return false
}

func (c *Cloudinary) sortParamsByKey(options Option) Option {
	options["timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	result := Option{}
	sortedKeys := []string{}

	for key := range options {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		result[key] = options[key]
	}

	return result
}

func (c *Cloudinary) createSignature(options Option) string {
	signature := ""
	i := 0

	for key, value := range options {
		if i != 0 {
			signature += "&"
		}
		signature += key + "=" + value
		i++
	}

	hash := sha1.New()
	hash.Write([]byte(signature + c.secretKey))

	return hex.EncodeToString(hash.Sum(nil))
}

func (c *Cloudinary) send(url string, postParams url.Values, options Option) []byte {
	postParams.Add("api_key", c.publicKey)
	postParams.Add("signature", c.createSignature(options))

	req, err := http.NewRequest("POST", c.urls[url], strings.NewReader(postParams.Encode()))

	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
