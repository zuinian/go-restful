package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"srm/srm-apiserver/models"
)

type SRMCClient struct {
	Host string
}

func (c *SRMCClient) SchoolAdd(school models.School) error {
	schoolJson, err := json.Marshal(school)
	if err != nil {
		return err
	}

	_, err = resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(schoolJson).
		Post(fmt.Sprintf("%s/api/v1/schools", c.Host))

	return err
}

func (c *SRMCClient) SchoolsGet() (*resty.Response, error) {
	res, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(nil).
		Get(fmt.Sprintf("%s/api/v1/schools", c.Host))
	return res, err
}