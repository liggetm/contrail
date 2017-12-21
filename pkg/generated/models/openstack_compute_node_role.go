package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	UUID                        string         `json:"uuid"`
	ParentType                  string         `json:"parent_type"`
	DisplayName                 string         `json:"display_name"`
	Perms2                      *PermType2     `json:"perms2"`
	ProvisioningStartTime       string         `json:"provisioning_start_time"`
	DefaultGateway              string         `json:"default_gateway"`
	ParentUUID                  string         `json:"parent_uuid"`
	Annotations                 *KeyValuePairs `json:"annotations"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage"`
	ProvisioningState           string         `json:"provisioning_state"`
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members"`
	VrouterType                 string         `json:"vrouter_type"`
	IDPerms                     *IdPermsType   `json:"id_perms"`
	ProvisioningProgress        int            `json:"provisioning_progress"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface"`
	FQName                      []string       `json:"fq_name"`
	ProvisioningLog             string         `json:"provisioning_log"`
}

// String returns json representation of the object
func (model *OpenstackComputeNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackComputeNodeRole makes OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRole() *OpenstackComputeNodeRole {
	return &OpenstackComputeNodeRole{
		//TODO(nati): Apply default
		VrouterBondInterfaceMembers: "",
		VrouterType:                 "",
		IDPerms:                     MakeIdPermsType(),
		ProvisioningProgress:        0,
		ProvisioningState:           "",
		VrouterBondInterface:        "",
		FQName:                      []string{},
		ProvisioningLog:             "",
		UUID:                        "",
		ParentType:                  "",
		DisplayName:                 "",
		Perms2:                      MakePermType2(),
		DefaultGateway:              "",
		ParentUUID:                  "",
		Annotations:                 MakeKeyValuePairs(),
		ProvisioningProgressStage:   "",
		ProvisioningStartTime:       "",
	}
}

// InterfaceToOpenstackComputeNodeRole makes OpenstackComputeNodeRole from interface
func InterfaceToOpenstackComputeNodeRole(iData interface{}) *OpenstackComputeNodeRole {
	data := iData.(map[string]interface{})
	return &OpenstackComputeNodeRole{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		DefaultGateway: data["default_gateway"].(string),

		//{"title":"Default Gateway","description":"Default Gateway","default":"","type":"string","permission":["create","update"]}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		VrouterBondInterfaceMembers: data["vrouter_bond_interface_members"].(string),

		//{"title":"vRouter Bond Interface Members","description":"vRouter Bond Interface Members","default":"ens7f0,ens7f1","type":"string","permission":["create","update"]}
		VrouterType: data["vrouter_type"].(string),

		//{"title":"vRouter Type","default":"kernel","type":"string","permission":["create","update"],"enum":["kernel","dpdk","smartNiC"]}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		VrouterBondInterface: data["vrouter_bond_interface"].(string),

		//{"title":"vRouter Bond Interface","description":"vRouter Bond Interface","default":"bond0","type":"string","permission":["create","update"]}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}

	}
}

// InterfaceToOpenstackComputeNodeRoleSlice makes a slice of OpenstackComputeNodeRole from interface
func InterfaceToOpenstackComputeNodeRoleSlice(data interface{}) []*OpenstackComputeNodeRole {
	list := data.([]interface{})
	result := MakeOpenstackComputeNodeRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToOpenstackComputeNodeRole(item))
	}
	return result
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
