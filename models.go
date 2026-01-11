package spaceship

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

type ListDNSRecordsResponse struct {
	Items []DNSRecord `json:"items"`
	Total int         `json:"total"`
}

type SaveDNSRecordsRequest struct {
	Force bool        `json:"force,omitempty"`
	Items []DNSRecord `json:"items"`
}
