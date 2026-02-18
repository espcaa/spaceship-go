package spaceship

import "fmt"

func (c *Client) GetDomainDNSRecords(domainName string, take, skip int, orderBy string) (*[]DNSRecord, error) {
	path := fmt.Sprintf("/dns/records/%s?take=%d&skip=%d", domainName, take, skip)

	if orderBy != "" {
		path += "&orderBy=" + orderBy
	}

	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var result ListDNSRecordsResponse
	if err := c.do(req, &result); err != nil {
		return nil, err
	}

	return &result.Items, nil
}

func (c *Client) SaveDNSRecords(domain string, force bool, records []DNSRecord) error {
	path := fmt.Sprintf("/dns/records/%s", domain)

	body := SaveDNSRecordsRequest{
		Force: force,
		Items: records,
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

	body := SaveDNSRecordsRequest{
		Items: records,
	}

	req, err := c.newRequest("DELETE", path, body)
	if err != nil {
		return err
	}

	if err := c.do(req, nil); err != nil {
		return err
	}

	return nil
}
