package spaceship

import (
	"encoding/json"
	"fmt"
)

type DomainInfo struct {
	Name               string            `json:"name"`
	UnicodeName        string            `json:"unicodeName"`
	IsPremium          bool              `json:"isPremium"`
	AutoRenew          bool              `json:"autoRenew"`
	RegistrationDate   string            `json:"registrationDate"`
	ExpirationDate     string            `json:"expirationDate"`
	LifecycleStatus    string            `json:"lifecycleStatus"`
	VerificationStatus string            `json:"verificationStatus"`
	EppStatuses        []string          `json:"eppStatuses"`
	Suspensions        []string          `json:"suspensions"`
	PrivacyProtection  PrivacyProtection `json:"privacyProtection"`
	Nameservers        Nameservers       `json:"nameservers"`
	Contacts           ContactInfo       `json:"contacts"`
}

type ContactInfo struct {
	Registrant string   `json:"registrant"`
	Admin      string   `json:"admin"`
	Tech       string   `json:"tech"`
	Billing    string   `json:"billing"`
	Attributes []string `json:"attributes"`
}

type ListDomainsResponse struct {
	Items []DomainInfo `json:"items"`
	Total int          `json:"total"`
}

type PrivacyProtection struct {
	ContactForm bool   `json:"contactForm"`
	Level       string `json:"level"`
}

type Nameservers struct {
	Provider string   `json:"provider"`
	Hosts    []string `json:"hosts"`
}

type DNSRecord interface {
	GetType() string
}

type MXRecord struct {
	Name       string         `json:"name"`
	TTL        int            `json:"ttl"`
	Group      DNSRecordGroup `json:"group"`
	Exchange   string         `json:"exchange"`
	Preference uint16         `json:"preference"`
}

type CNAMERecord struct {
	Name  string         `json:"name"`
	TTL   int            `json:"ttl"`
	Group DNSRecordGroup `json:"group"`
	CNAME string         `json:"cname"`
}

type TXTRecord struct {
	Value string         `json:"value"`
	Name  string         `json:"name"`
	TTL   int            `json:"ttl"`
	Group DNSRecordGroup `json:"group"`
}

type TLSARecord struct {
	Port            string         `json:"port"`
	Protocol        string         `json:"protocol"`
	Usage           uint16         `json:"usage"`
	Selector        uint16         `json:"selector"`
	Matching        uint16         `json:"matching"`
	AssociationData string         `json:"associationData"`
	Name            string         `json:"name"`
	TTL             int            `json:"ttl"`
	Group           DNSRecordGroup `json:"group"`
}

type SVCBRecord struct {
	Port        string         `json:"port"`
	Scheme      string         `json:"scheme"`
	SvcPriority uint16         `json:"svcPriority"`
	TargetName  string         `json:"targetName"`
	SvcParams   string         `json:"svcParams"`
	Name        string         `json:"name"`
	TTL         int            `json:"ttl"`
	Group       DNSRecordGroup `json:"group"`
}

type SRVRecord struct {
	Service  string         `json:"service"`
	Protocol string         `json:"protocol"`
	Priority uint16         `json:"priority"`
	Weight   uint16         `json:"weight"`
	Port     uint16         `json:"port"`
	Target   string         `json:"target"`
	Name     string         `json:"name"`
	TTL      int            `json:"ttl"`
	Group    DNSRecordGroup `json:"group"`
}

type PTRRecord struct {
	Pointer string         `json:"pointer"`
	Name    string         `json:"name"`
	TTL     int            `json:"ttl"`
	Group   DNSRecordGroup `json:"group"`
}

type NSRecord struct {
	Nameserver string         `json:"nameserver"`
	Name       string         `json:"name"`
	TTL        int            `json:"ttl"`
	Group      DNSRecordGroup `json:"group"`
}

type HTTPSRecord struct {
	Port        string         `json:"port"`
	Scheme      string         `json:"scheme"`
	SvcPriority uint16         `json:"svcPriority"`
	TargetName  string         `json:"targetName"`
	SvcParams   string         `json:"svcParams"`
	Name        string         `json:"name"`
	TTL         int            `json:"ttl"`
	Group       DNSRecordGroup `json:"group"`
}

type CAARecord struct {
	Flag  uint8          `json:"flag"`
	Tag   string         `json:"tag"`
	Value string         `json:"value"`
	Name  string         `json:"name"`
	TTL   int            `json:"ttl"`
	Group DNSRecordGroup `json:"group"`
}

type ARecord struct {
	Adress string         `json:"address"`
	Name   string         `json:"name"`
	TTL    int            `json:"ttl"`
	Group  DNSRecordGroup `json:"group"`
}

type AliasRecord struct {
	AliasTarget string         `json:"aliasTarget"`
	Name        string         `json:"name"`
	TTL         int            `json:"ttl"`
	Group       DNSRecordGroup `json:"group"`
}

type AAAARecord struct {
	Adress string         `json:"address"`
	Name   string         `json:"name"`
	TTL    int            `json:"ttl"`
	Group  DNSRecordGroup `json:"group"`
}

type DNSRecordGroup struct {
	Type string `json:"type"`
}

func (r MXRecord) GetType() string {
	return "MX"
}

func (r CNAMERecord) GetType() string {
	return "CNAME"
}

func (r TXTRecord) GetType() string {
	return "TXT"
}

func (r TLSARecord) GetType() string {
	return "TLSA"
}

func (r SVCBRecord) GetType() string {
	return "SVCB"
}

func (r SRVRecord) GetType() string {
	return "SRV"
}

func (r PTRRecord) GetType() string {
	return "PTR"
}

func (r NSRecord) GetType() string {
	return "NS"
}

func (r HTTPSRecord) GetType() string {
	return "HTTPS"
}

func (r CAARecord) GetType() string {
	return "CAA"
}

func (r ARecord) GetType() string {
	return "A"
}

func (r AliasRecord) GetType() string {
	return "ALIAS"
}

func (r AAAARecord) GetType() string {
	return "AAAA"
}

type ListDNSRecordsResponse struct {
	Items []DNSRecord `json:"-"`
	Total int         `json:"total"`
}

func (r *ListDNSRecordsResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Items []json.RawMessage `json:"items"`
		Total int               `json:"total"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	r.Total = raw.Total
	r.Items = make([]DNSRecord, 0, len(raw.Items))

	for _, item := range raw.Items {
		var peek struct {
			Group DNSRecordGroup `json:"group"`
		}
		if err := json.Unmarshal(item, &peek); err != nil {
			return err
		}

		var rec DNSRecord
		switch peek.Group.Type {
		case "A":
			var v ARecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "AAAA":
			var v AAAARecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "ALIAS":
			var v AliasRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "CAA":
			var v CAARecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "CNAME":
			var v CNAMERecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "HTTPS":
			var v HTTPSRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "MX":
			var v MXRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "NS":
			var v NSRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "PTR":
			var v PTRRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "SRV":
			var v SRVRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "SVCB":
			var v SVCBRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "TLSA":
			var v TLSARecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		case "TXT":
			var v TXTRecord
			if err := json.Unmarshal(item, &v); err != nil {
				return err
			}
			rec = v
		default:
			return fmt.Errorf("unknown DNS record type: %s", peek.Group.Type)
		}

		r.Items = append(r.Items, rec)
	}

	return nil
}

type SaveDNSRecordsRequest struct {
	Force bool        `json:"force,omitempty"`
	Items []DNSRecord `json:"items"`
}
