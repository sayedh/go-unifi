package unifi

import (
	"context"
	"fmt"
)

// just to fix compile issues with the import.
var (
	_ fmt.Formatter
	_ context.Context
)

// This is a v2 API object, so manually coded for now, need to figure out generation...

// Example packet from api sniffing
// {
// 	"_id": "12345abcdef12345abcdef11",
// 	"enabled": true,
// 	"key": "example.com",
// 	"port": 0,
// 	"priority": 0,
// 	"record_type": "A",
// 	"ttl": 0,
// 	"value": "192.168.1.1",
// 	"weight": 0
// }

type StaticDns struct {
	ID 					string	`json:"_id,omitempty"`
	Enabled             bool	`json:"enabled,omitempty"`
	Key   				string	`json:"key"`
	Port                int		`json:"port,omitempty"`
	Priority            int		`json:"priority,omitempty"`
	RecordType			string	`json:"record_type"`
	TTL					int		`json:"ttl,omitempty"`
	Value   			string	`json:"value"`
	Weight				int		`json:"weight,omitempty"`
}

func (c *Client) CreateStaticDns(ctx context.Context, site string, d *StaticDns) (*StaticDns, error) {
	var respBody StaticDns

	err := c.do(ctx, "POST", fmt.Sprintf("%s/site/%s/static-dns", c.apiV2Path, site), d, &respBody)
	if err != nil {
		return nil, err
	}

	// if len(respBody.Data) != 1 {
	// 	return nil, &NotFoundError{}
	// }

	return &respBody, nil
}

func (c *Client) ListStaticDNS(ctx context.Context, site string) ([]StaticDns, error) {
	var respBody []StaticDns

	err := c.do(ctx, "GET", fmt.Sprintf("%s/site/%s/static-dns", c.apiV2Path, site), nil, &respBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (c *Client) GetStaticDns(ctx context.Context, site, id string) (*StaticDns, error) {
	respBody, err := c.ListStaticDNS(ctx, site)
	if err != nil {
		return nil, err
	}

	if len(respBody) == 0 {
		return nil, &NotFoundError{}
	}

	for _, dns := range respBody {
		if dns.ID == id {
			return &dns, nil
		}
	}

	return nil, &NotFoundError{}
}

func (c *Client) UpdateStaticDns(ctx context.Context, site string, d *StaticDns) (*StaticDns, error) {
	var respBody StaticDns

	err := c.do(ctx, "PUT", fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiV2Path, site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	// if len(respBody.Data) != 1 {
	// 	return nil, &NotFoundError{}
	// }

	return &respBody, nil
}

func (c *Client) DeleteStaticDns(ctx context.Context, site, id string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("%s/site/%s/static-dns/%s", c.apiV2Path, site, id), struct{}{}, nil)
	if err != nil {
		return err
	}
	return nil
}
