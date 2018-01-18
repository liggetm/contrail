package models

// VirtualDnsType

import "encoding/json"

// VirtualDnsType
type VirtualDnsType struct {
	RecordOrder              DnsRecordOrderType    `json:"record_order,omitempty"`
	FloatingIPRecord         FloatingIpDnsNotation `json:"floating_ip_record,omitempty"`
	DomainName               string                `json:"domain_name,omitempty"`
	ExternalVisible          bool                  `json:"external_visible,omitempty"`
	NextVirtualDNS           string                `json:"next_virtual_DNS,omitempty"`
	DynamicRecordsFromClient bool                  `json:"dynamic_records_from_client,omitempty"`
	ReverseResolution        bool                  `json:"reverse_resolution,omitempty"`
	DefaultTTLSeconds        int                   `json:"default_ttl_seconds,omitempty"`
}

// String returns json representation of the object
func (model *VirtualDnsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualDnsType makes VirtualDnsType
func MakeVirtualDnsType() *VirtualDnsType {
	return &VirtualDnsType{
		//TODO(nati): Apply default
		DynamicRecordsFromClient: false,
		ReverseResolution:        false,
		DefaultTTLSeconds:        0,
		RecordOrder:              MakeDnsRecordOrderType(),
		FloatingIPRecord:         MakeFloatingIpDnsNotation(),
		DomainName:               "",
		ExternalVisible:          false,
		NextVirtualDNS:           "",
	}
}

// MakeVirtualDnsTypeSlice() makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
	return []*VirtualDnsType{}
}
