package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
)

func TestOpenstackStorageNodeRole(t *testing.T) {
	t.Parallel()
	db := testDB
	common.UseTable(db, "openstack_storage_node_role")
	defer func() {
		common.ClearTable(db, "openstack_storage_node_role")
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeOpenstackStorageNodeRole()
	model.UUID = "dummy_uuid"

	err := common.DoInTransaction(db, func(tx *sql.Tx) error {
		return CreateOpenstackStorageNodeRole(tx, model)
	})
	if err != nil {
		t.Fatal("create failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListOpenstackStorageNodeRole(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 1 {
			return fmt.Errorf("expected one element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		return DeleteOpenstackStorageNodeRole(tx, model.UUID, nil)
	})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	err = common.DoInTransaction(db, func(tx *sql.Tx) error {
		models, err := ListOpenstackStorageNodeRole(tx, &common.ListSpec{Limit: 1})
		if err != nil {
			return err
		}
		if len(models) != 0 {
			return fmt.Errorf("expected no element")
		}
		return nil
	})
	if err != nil {
		t.Fatal("list failed", err)
	}
	return
}
