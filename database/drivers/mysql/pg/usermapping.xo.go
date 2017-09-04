// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// UserMappingTable is the database name for the table.
const UserMappingTable = "information_schema.user_mappings"

// UserMapping represents a row from 'information_schema.user_mappings'.
type UserMapping struct {
	AuthorizationIdentifier SQLIdentifier `yaml:"authorization_identifier,omitempty"` // authorization_identifier
	ForeignServerCatalog    SQLIdentifier `yaml:"foreign_server_catalog,omitempty"`   // foreign_server_catalog
	ForeignServerName       SQLIdentifier `yaml:"foreign_server_name,omitempty"`      // foreign_server_name
}

// Constants defining each column in the table.
const (
	UserMappingAuthorizationIdentifierField = "authorization_identifier"
	UserMappingForeignServerCatalogField    = "foreign_server_catalog"
	UserMappingForeignServerNameField       = "foreign_server_name"
)

// WhereClauses for every type in UserMapping.
var (
	UserMappingAuthorizationIdentifierWhere SQLIdentifierField = "authorization_identifier"
	UserMappingForeignServerCatalogWhere    SQLIdentifierField = "foreign_server_catalog"
	UserMappingForeignServerNameWhere       SQLIdentifierField = "foreign_server_name"
)

// QueryOneUserMapping retrieves a row from 'information_schema.user_mappings' as a UserMapping.
func QueryOneUserMapping(db XODB, where WhereClause, order OrderBy) (*UserMapping, error) {
	const origsqlstr = `SELECT ` +
		`authorization_identifier, foreign_server_catalog, foreign_server_name ` +
		`FROM information_schema.user_mappings WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	um := &UserMapping{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&um.AuthorizationIdentifier, &um.ForeignServerCatalog, &um.ForeignServerName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return um, nil
}

// QueryUserMapping retrieves rows from 'information_schema.user_mappings' as a slice of UserMapping.
func QueryUserMapping(db XODB, where WhereClause, order OrderBy) ([]*UserMapping, error) {
	const origsqlstr = `SELECT ` +
		`authorization_identifier, foreign_server_catalog, foreign_server_name ` +
		`FROM information_schema.user_mappings WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*UserMapping
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		um := UserMapping{}

		err = q.Scan(&um.AuthorizationIdentifier, &um.ForeignServerCatalog, &um.ForeignServerName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &um)
	}
	return vals, nil
}

// AllUserMapping retrieves all rows from 'information_schema.user_mappings' as a slice of UserMapping.
func AllUserMapping(db XODB, order OrderBy) ([]*UserMapping, error) {
	const origsqlstr = `SELECT ` +
		`authorization_identifier, foreign_server_catalog, foreign_server_name ` +
		`FROM information_schema.user_mappings`

	sqlstr := origsqlstr + order.String()

	var vals []*UserMapping
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		um := UserMapping{}

		err = q.Scan(&um.AuthorizationIdentifier, &um.ForeignServerCatalog, &um.ForeignServerName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &um)
	}
	return vals, nil
}
