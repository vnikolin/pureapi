package pureapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

//basicAuth ... will create the basicauth encoded string for the credentials
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *StorageClient) formatPath(path string) string {
	return fmt.Sprintf("https://%s/api/%s/%s", c.Hostname, c.Version, path)
}

func (c *StorageClient) getAPIToken() error {

	authURL, err := url.Parse(c.formatPath("auth/apitoken"))
	if err != nil {
		return err
	}

	data := map[string]string{"username": c.Username, "password": c.Password}
	jsonValue, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", authURL.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	r, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	t := &Auth{}
	err = json.NewDecoder(r.Body).Decode(t)
	c.APIToken = t.Token

	return err
}

// Authenticate to the API and store the session
func (c *StorageClient) login() error {
	authURL := c.formatPath("auth/session")

	data := map[string]string{"api_token": c.APIToken}
	jsonValue, _ := json.Marshal(data)
	_, err := c.Client.Post(authURL, "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}
	return nil
}

//queryData ... will make REST verbs based on the url
func (c *StorageClient) QueryData(call string, link string, data []byte) ([]byte, http.Header, int, error) {

	client, _ := NewStorageClient(c.Hostname, c.Username, c.Password, c.APIToken, c.Version)
	req, err := http.NewRequest(call, link, bytes.NewBuffer(data))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := errors.New(StatusBadRequest)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 403 {
			err := errors.New(StatusForbidden)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 404 {
			err := errors.New(StatusNotFound)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 409 {
			err := errors.New(StatusConflict)
			return nil, resp.Header, resp.StatusCode, err
		}

	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, resp.StatusCode, err
	}

	return _body, resp.Header, resp.StatusCode, nil
}
