package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeDashboard makes Dashboard
func MakeDashboard() *Dashboard {
	return &Dashboard{
		//TODO(nati): Apply default
		UUID:            "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
		IDPerms:         MakeIdPermsType(),
		DisplayName:     "",
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		ContainerConfig: "",
	}
}

// MakeDashboard makes Dashboard
func InterfaceToDashboard(i interface{}) *Dashboard {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Dashboard{
		//TODO(nati): Apply default
		UUID:            schema.InterfaceToString(m["uuid"]),
		ParentUUID:      schema.InterfaceToString(m["parent_uuid"]),
		ParentType:      schema.InterfaceToString(m["parent_type"]),
		FQName:          schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:     schema.InterfaceToString(m["display_name"]),
		Annotations:     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:          InterfaceToPermType2(m["perms2"]),
		ContainerConfig: schema.InterfaceToString(m["container_config"]),
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}

// InterfaceToDashboardSlice() makes a slice of Dashboard
func InterfaceToDashboardSlice(i interface{}) []*Dashboard {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Dashboard{}
	for _, item := range list {
		result = append(result, InterfaceToDashboard(item))
	}
	return result
}