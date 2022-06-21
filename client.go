package pureapi

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
)

//StorageClient ... Contstructor required Variables
type StorageClient struct {
	Hostname string
	Username string
	Password string
	APIToken string
	Version  string

	Client *http.Client
}

//NewStorageClient ... Initializes the Constructor with the above variables
func NewStorageClient(target string, username string, password string, apiToken string,
	version string) (*StorageClient, error) {

	// Create a new Client instance
	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &StorageClient{Hostname: target, Username: username, Password: password, APIToken: apiToken, Version: version}
	c.Client = &http.Client{Transport: tr, Jar: cookieJar}

	// Get an API Token if not provided
	if apiToken == "" {
		c.getAPIToken()
	}

	// Authenticate to the API and store the session
	err := c.login()
	if err != nil {
		return nil, err
	}

	return c, err
}
