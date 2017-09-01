// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// UserMappingOptionTable is the database name for the table.
const UserMappingOptionTable = "information_schema.user_mapping_options"

// UserMappingOption represents a row from 'information_schema.user_mapping_options'.
type UserMappingOption struct {
	AuthorizationIdentifier SQLIdentifier `yaml:"authorization_identifier,omitempty"` // authorization_identifier
	ForeignServerCatalog    SQLIdentifier `yaml:"foreign_server_catalog,omitempty"`   // foreign_server_catalog
	ForeignServerName       SQLIdentifier `yaml:"foreign_server_name,omitempty"`      // foreign_server_name
	OptionName              SQLIdentifier `yaml:"option_name,omitempty"`              // option_name
	OptionValue             CharacterData `yaml:"option_value,omitempty"`             // option_value
}

// Constants defining each column in the table.
const (
	UserMappingOptionAuthorizationIdentifierField = "authorization_identifier"
	UserMappingOptionForeignServerCatalogField    = "foreign_server_catalog"
	UserMappingOptionForeignServerNameField       = "foreign_server_name"
	UserMappingOptionOptionNameField              = "option_name"
	UserMappingOptionOptionValueField             = "option_value"
)

// WhereClauses for every type in UserMappingOption.
var (
	UserMappingOptionAuthorizationIdentifierWhere SQLIdentifierField = "authorization_identifier"
	UserMappingOptionForeignServerCatalogWhere    SQLIdentifierField = "foreign_server_catalog"
	UserMappingOptionForeignServerNameWhere       SQLIdentifierField = "foreign_server_name"
	UserMappingOptionOptionNameWhere              SQLIdentifierField = "option_name"
	UserMappingOptionOptionValueWhere             CharacterDataField = "option_value"
)

// QueryOneUserMappingOption retrieves a row from 'information_schema.user_mapping_options' as a UserMappingOption.
func QueryOneUserMappingOption(db XODB, where WhereClause, order OrderBy) (*UserMappingOption, error) {
	const origsqlstr = `SELECT ` +
		`authorization_identifier, foreign_server_catalog, foreign_server_name, option_name, option_value ` +
		`FROM information_schema.user_mapping_options WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	umo := &UserMappingOption{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&umo.AuthorizationIdentifier, &umo.ForeignServerCatalog, &umo.ForeignServerName, &umo.OptionName, &umo.OptionValue)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return umo, nil
}

// QueryUserMappingOption retrieves rows from 'information_schema.user_mapping_options' as a slice of UserMappingOption.
func QueryUserMappingOption(db XODB, where WhereClause, order OrderBy) ([]*UserMappingOption, error) {
	const origsqlstr = `SELECT ` +
		`authorization_identifier, foreign_server_catalog, foreign_server_name, option_name, option_value ` +
		`FROM information_schema.user_mapping_options WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*UserMappingOption
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		umo := UserMappingOption{}

		err = q.Scan(&umo.AuthorizationIdentifier, &umo.ForeignServerCatalog, &umo.ForeignServerName, &umo.OptionName, &umo.OptionValue)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &umo)
	}
	return vals, nil
}

// AllUserMappingOption retrieves all rows from 'information_schema.user_mapping_options' as a slice of UserMappingOption.
func AllUserMappingOption(db XODB, order OrderBy) ([]*UserMappingOption, error) {
	const origsqlstr = `SELECT ` +
		`authorization_identifier, foreign_server_catalog, foreign_server_name, option_name, option_value ` +
		`FROM information_schema.user_mapping_options`

	sqlstr := origsqlstr + order.String()

	var vals []*UserMappingOption
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		umo := UserMappingOption{}

		err = q.Scan(&umo.AuthorizationIdentifier, &umo.ForeignServerCatalog, &umo.ForeignServerName, &umo.OptionName, &umo.OptionValue)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &umo)
	}
	return vals, nil
}