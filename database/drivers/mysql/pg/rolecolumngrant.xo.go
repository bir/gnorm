// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// RoleColumnGrantTable is the database name for the table.
const RoleColumnGrantTable = "information_schema.role_column_grants"

// RoleColumnGrant represents a row from 'information_schema.role_column_grants'.
type RoleColumnGrant struct {
	Grantor       SQLIdentifier `yaml:"grantor,omitempty"`        // grantor
	Grantee       SQLIdentifier `yaml:"grantee,omitempty"`        // grantee
	TableCatalog  SQLIdentifier `yaml:"table_catalog,omitempty"`  // table_catalog
	TableSchema   SQLIdentifier `yaml:"table_schema,omitempty"`   // table_schema
	TableName     SQLIdentifier `yaml:"table_name,omitempty"`     // table_name
	ColumnName    SQLIdentifier `yaml:"column_name,omitempty"`    // column_name
	PrivilegeType CharacterData `yaml:"privilege_type,omitempty"` // privilege_type
	IsGrantable   YesOrNo       `yaml:"is_grantable,omitempty"`   // is_grantable
}

// Constants defining each column in the table.
const (
	RoleColumnGrantGrantorField       = "grantor"
	RoleColumnGrantGranteeField       = "grantee"
	RoleColumnGrantTableCatalogField  = "table_catalog"
	RoleColumnGrantTableSchemaField   = "table_schema"
	RoleColumnGrantTableNameField     = "table_name"
	RoleColumnGrantColumnNameField    = "column_name"
	RoleColumnGrantPrivilegeTypeField = "privilege_type"
	RoleColumnGrantIsGrantableField   = "is_grantable"
)

// WhereClauses for every type in RoleColumnGrant.
var (
	RoleColumnGrantGrantorWhere       SQLIdentifierField = "grantor"
	RoleColumnGrantGranteeWhere       SQLIdentifierField = "grantee"
	RoleColumnGrantTableCatalogWhere  SQLIdentifierField = "table_catalog"
	RoleColumnGrantTableSchemaWhere   SQLIdentifierField = "table_schema"
	RoleColumnGrantTableNameWhere     SQLIdentifierField = "table_name"
	RoleColumnGrantColumnNameWhere    SQLIdentifierField = "column_name"
	RoleColumnGrantPrivilegeTypeWhere CharacterDataField = "privilege_type"
	RoleColumnGrantIsGrantableWhere   YesOrNoField       = "is_grantable"
)

// QueryOneRoleColumnGrant retrieves a row from 'information_schema.role_column_grants' as a RoleColumnGrant.
func QueryOneRoleColumnGrant(db XODB, where WhereClause, order OrderBy) (*RoleColumnGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_column_grants WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	rcg := &RoleColumnGrant{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&rcg.Grantor, &rcg.Grantee, &rcg.TableCatalog, &rcg.TableSchema, &rcg.TableName, &rcg.ColumnName, &rcg.PrivilegeType, &rcg.IsGrantable)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rcg, nil
}

// QueryRoleColumnGrant retrieves rows from 'information_schema.role_column_grants' as a slice of RoleColumnGrant.
func QueryRoleColumnGrant(db XODB, where WhereClause, order OrderBy) ([]*RoleColumnGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_column_grants WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*RoleColumnGrant
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		rcg := RoleColumnGrant{}

		err = q.Scan(&rcg.Grantor, &rcg.Grantee, &rcg.TableCatalog, &rcg.TableSchema, &rcg.TableName, &rcg.ColumnName, &rcg.PrivilegeType, &rcg.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &rcg)
	}
	return vals, nil
}

// AllRoleColumnGrant retrieves all rows from 'information_schema.role_column_grants' as a slice of RoleColumnGrant.
func AllRoleColumnGrant(db XODB, order OrderBy) ([]*RoleColumnGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_column_grants`

	sqlstr := origsqlstr + order.String()

	var vals []*RoleColumnGrant
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		rcg := RoleColumnGrant{}

		err = q.Scan(&rcg.Grantor, &rcg.Grantee, &rcg.TableCatalog, &rcg.TableSchema, &rcg.TableName, &rcg.ColumnName, &rcg.PrivilegeType, &rcg.IsGrantable)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &rcg)
	}
	return vals, nil
}
