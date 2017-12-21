package models

// IpamSubnetType

import "encoding/json"

// IpamSubnetType
type IpamSubnetType struct {
	HostRoutes       *RouteTableType       `json:"host_routes"`
	AllocUnit        int                   `json:"alloc_unit"`
	Created          string                `json:"created"`
	DNSNameservers   []string              `json:"dns_nameservers"`
	SubnetUUID       string                `json:"subnet_uuid"`
	SubnetName       string                `json:"subnet_name"`
	Subnet           *SubnetType           `json:"subnet"`
	AddrFromStart    bool                  `json:"addr_from_start"`
	EnableDHCP       bool                  `json:"enable_dhcp"`
	DNSServerAddress IpAddressType         `json:"dns_server_address"`
	DHCPOptionList   *DhcpOptionsListType  `json:"dhcp_option_list"`
	LastModified     string                `json:"last_modified"`
	DefaultGateway   IpAddressType         `json:"default_gateway"`
	AllocationPools  []*AllocationPoolType `json:"allocation_pools"`
}

// String returns json representation of the object
func (model *IpamSubnetType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeIpamSubnetType makes IpamSubnetType
func MakeIpamSubnetType() *IpamSubnetType {
	return &IpamSubnetType{
		//TODO(nati): Apply default
		AllocUnit:        0,
		Created:          "",
		DNSNameservers:   []string{},
		SubnetUUID:       "",
		HostRoutes:       MakeRouteTableType(),
		Subnet:           MakeSubnetType(),
		AddrFromStart:    false,
		EnableDHCP:       false,
		DNSServerAddress: MakeIpAddressType(),
		SubnetName:       "",
		DHCPOptionList:   MakeDhcpOptionsListType(),
		LastModified:     "",
		DefaultGateway:   MakeIpAddressType(),

		AllocationPools: MakeAllocationPoolTypeSlice(),
	}
}

// InterfaceToIpamSubnetType makes IpamSubnetType from interface
func InterfaceToIpamSubnetType(iData interface{}) *IpamSubnetType {
	data := iData.(map[string]interface{})
	return &IpamSubnetType{
		HostRoutes: InterfaceToRouteTableType(data["host_routes"]),

		//{"description":"Host routes to be sent via DHCP for VM(s) in this subnet, Next hop for these routes is always default gateway","type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}}
		AllocUnit: data["alloc_unit"].(int),

		//{"description":"allocation unit for this subnet to allocate bulk ip addresses","type":"integer"}
		Created: data["created"].(string),

		//{"description":"timestamp when subnet object gets created","type":"string"}
		DNSNameservers: data["dns_nameservers"].([]string),

		//{"description":"Tenant DNS servers ip address in tenant DNS method","type":"array","item":{"type":"string"}}
		SubnetUUID: data["subnet_uuid"].(string),

		//{"description":"Subnet UUID is auto generated by the system","type":"string"}
		SubnetName: data["subnet_name"].(string),

		//{"description":"User provided name for this subnet","type":"string"}
		Subnet: InterfaceToSubnetType(data["subnet"]),

		//{"description":"ip prefix and length for the subnet","type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}
		AddrFromStart: data["addr_from_start"].(bool),

		//{"description":"Start address allocation from start or from end of address range.","type":"boolean"}
		EnableDHCP: data["enable_dhcp"].(bool),

		//{"description":"Enable DHCP for the VM(s) in this subnet","type":"boolean"}
		DNSServerAddress: InterfaceToIpAddressType(data["dns_server_address"]),

		//{"description":"DNS server ip address in the subnet, if not provided one is auto generated by the system.","type":"string"}
		DHCPOptionList: InterfaceToDhcpOptionsListType(data["dhcp_option_list"]),

		//{"description":"DHCP options list to be sent via DHCP for  VM(s) in this subnet","type":"object","properties":{"dhcp_option":{"type":"array","item":{"type":"object","properties":{"dhcp_option_name":{"type":"string"},"dhcp_option_value":{"type":"string"},"dhcp_option_value_bytes":{"type":"string"}}}}}}
		LastModified: data["last_modified"].(string),

		//{"description":"timestamp when subnet object gets updated","type":"string"}
		DefaultGateway: InterfaceToIpAddressType(data["default_gateway"]),

		//{"description":"default-gateway ip address in the subnet, if not provided one is auto generated by the system.","type":"string"}

		AllocationPools: InterfaceToAllocationPoolTypeSlice(data["allocation_pools"]),

		//{"description":"List of ranges of ip address within the subnet from which to allocate ip address. default is entire prefix","type":"array","item":{"type":"object","properties":{"end":{"type":"string"},"start":{"type":"string"},"vrouter_specific_pool":{"type":"boolean"}}}}

	}
}

// InterfaceToIpamSubnetTypeSlice makes a slice of IpamSubnetType from interface
func InterfaceToIpamSubnetTypeSlice(data interface{}) []*IpamSubnetType {
	list := data.([]interface{})
	result := MakeIpamSubnetTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamSubnetType(item))
	}
	return result
}

// MakeIpamSubnetTypeSlice() makes a slice of IpamSubnetType
func MakeIpamSubnetTypeSlice() []*IpamSubnetType {
	return []*IpamSubnetType{}
}
