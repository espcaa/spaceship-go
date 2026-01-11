package spaceship

import (
	"fmt"
)

func (c *Client) ListDomains(take, skip int, orderBy string) (*ListDomainsResponse, error) {
	path := fmt.Sprintf("/domains?take=%d&skip=%d", take, skip)
	if orderBy != "" {
		path += "&orderBy=" + orderBy
	}

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result ListDomainsResponse
	if err := c.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetDomainInfo(domainName string) (*DomainInfo, error) {
	path := fmt.Sprintf("/domains/%s", domainName)

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result DomainInfo
	if err := c.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
