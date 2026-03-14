package spaceship

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetDomainDNSRecords(domainName string, take, skip int, orderBy string) (ListDNSRecordsResponse, error) {
	path := fmt.Sprintf("/dns/records/%s?take=%d&skip=%d", domainName, take, skip)

	if orderBy != "" {
		path += "&orderBy=" + orderBy
	}

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return ListDNSRecordsResponse{}, err
	}

	var result ListDNSRecordsResponse
	if err := c.do(req, &result); err != nil {
		return ListDNSRecordsResponse{}, err
	}

	return result, nil
}

func (c *Client) SaveDNSRecords(domain string, force bool, records []DNSRecord) error {
	path := fmt.Sprintf("/dns/records/%s", domain)

	wrapped := make([]json.RawMessage, len(records))
	for i, r := range records {
		b, err := json.Marshal(r)
		if err != nil {
			return err
		}
		var m map[string]any
		json.Unmarshal(b, &m)
		m["type"] = r.GetType()
		wrapped[i], _ = json.Marshal(m)
	}

	body := struct {
		Force bool               `json:"force,omitempty"`
		Items []json.RawMessage  `json:"items"`
	}{
		Force: force,
		Items: wrapped,
	}

	req, err := c.newRequest("PUT", path, body)
	if err != nil {
		return err
	}

	if err := c.do(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteDNSRecords(domain string, records []DNSRecord) error {
	path := fmt.Sprintf("/dns/records/%s", domain)

	wrapped := make([]json.RawMessage, len(records))
	for i, r := range records {
		b, err := json.Marshal(r)
		if err != nil {
			return err
		}
		var m map[string]any
		json.Unmarshal(b, &m)
		m["type"] = r.GetType()
		wrapped[i], _ = json.Marshal(m)
	}

	req, err := c.newRequest("DELETE", path, wrapped)
	if err != nil {
		return err
	}

	if err := c.do(req, nil); err != nil {
		return err
	}

	return nil
}
