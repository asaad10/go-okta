package okta

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client        *http.Client
	org           string
	Url           string
	ApiToken      string
}


func NewClient(org string, apitoken string) *Client {
	client := Client{
		client: &http.Client{},
		org:    org,
		Url:    "okta.com",
		ApiToken: apitoken,
	}

	return &client
}

func (c *Client) UserByUsername(username string) (*UserByUsername, error) {

	var response = &UserByUsername{}
	err := c.call("users/?q="+username, "GET", nil, response)
	return response, err
}

func (c *Client) User(userID string) (*User, error) {

	var response = &User{}
	err := c.call("users/"+userID, "GET", nil, response)
	return response, err
}


func (c *Client) GetAppsForUser(userID string) (*AppList, error) {

	var response = &AppList{}
	err := c.call("apps?filter=user.id+eq+" + "%22"+userID+"%22", "GET", nil, response)
	return response, err
}

func (c *Client) call(endpoint, method string, request, response interface{}) error {
	data, _ := json.Marshal(request)

	var url = "https://" + c.org + "." + c.Url + "/api/v1/" + endpoint
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "SSWS "+ c.ApiToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		err := json.Unmarshal(body, &response)
		if err != nil {
			return err
		}
	} else {
		var errors ErrorResponse
		err = json.Unmarshal(body, &errors)

		return &errorResponse{
			HTTPCode: resp.StatusCode,
			Response: errors,
			Endpoint: url,
		}
	}

	return nil
}
