// GENERATED FILE: DO NOT EDIT!

package audiences

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client represents an API client.
type Client struct {
	service string
	APIKey  string
	client  *http.Client
}

// NewClient creates an API client.
func NewClient(service string, c *http.Client) *Client {
	client := &Client{}
	client.service = service
	if c != nil {
		client.client = c
	} else {
		client.client = http.DefaultClient
	}
	return client
}

//
func (client *Client) ListRecommendedAudiences(
	limit int64,
) (
	response *ListRecommendedAudiencesResponses,
	err error,
) {
	path := client.service + "/audiences/recommended"
	v := url.Values{}
	if client.APIKey != "" {
		v.Set("key", client.APIKey)
	}
	if len(v) > 0 {
		path = path + "?" + v.Encode()
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &ListRecommendedAudiencesResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audiences{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) ListAudiences(
	limit int64,
) (
	response *ListAudiencesResponses,
	err error,
) {
	path := client.service + "/audiences"
	v := url.Values{}
	if client.APIKey != "" {
		v.Set("key", client.APIKey)
	}
	if len(v) > 0 {
		path = path + "?" + v.Encode()
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &ListAudiencesResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audiences{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) CreateAudience(
	audience Audience,
) (
	response *CreateAudienceResponses,
	err error,
) {
	path := client.service + "/audiences"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(audience)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &CreateAudienceResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audience{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) ShowAudienceById(
	audienceId string,
) (
	response *ShowAudienceByIdResponses,
	err error,
) {
	path := client.service + "/audience/{audienceId}"
	path = strings.Replace(path, "{audienceId}", fmt.Sprintf("%v", audienceId), 1)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &ShowAudienceByIdResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audience{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) FilterAudienceByDemographic(
	audienceId string,
	demographic Demographic,
) (
	response *FilterAudienceByDemographicResponses,
	err error,
) {
	path := client.service + "/audiences/{audienceId}/demographic"
	path = strings.Replace(path, "{audienceId}", fmt.Sprintf("%v", audienceId), 1)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(demographic)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &FilterAudienceByDemographicResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audience{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) FilterAudienceByInterests(
	audienceId string,
	interests Interests,
) (
	response *FilterAudienceByInterestsResponses,
	err error,
) {
	path := client.service + "/audiences/{audienceId}/interests"
	path = strings.Replace(path, "{audienceId}", fmt.Sprintf("%v", audienceId), 1)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(interests)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &FilterAudienceByInterestsResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audience{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) FilterAudienceByUserStatus(
	audienceId string,
	userstatus UserStatus,
) (
	response *FilterAudienceByUserStatusResponses,
	err error,
) {
	path := client.service + "/audiences/{audienceId}/behaviours/userstatus"
	path = strings.Replace(path, "{audienceId}", fmt.Sprintf("%v", audienceId), 1)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(userstatus)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &FilterAudienceByUserStatusResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audience{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}

//
func (client *Client) FilterAudienceByCampaign(
	audienceId string,
	campaign Campaign,
) (
	response *FilterAudienceByCampaignResponses,
	err error,
) {
	path := client.service + "/audiences/{audienceId}/behaviours/campaign"
	path = strings.Replace(path, "{audienceId}", fmt.Sprintf("%v", audienceId), 1)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(campaign)
	req, err := http.NewRequest("POST", path, body)
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	req.Header = reqHeaders
	if err != nil {
		return
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	response = &FilterAudienceByCampaignResponses{}
	switch {
	case resp.StatusCode == 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Audience{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.OK = result
	case resp.StatusCode == 401:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &UnauthorizedError{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Code401 = result
	default:
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return nil, err
		}
		response.Default = result
	}
	return
}
