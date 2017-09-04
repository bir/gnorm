// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// ForeignDataWrapperTable is the database name for the table.
const ForeignDataWrapperTable = "information_schema.foreign_data_wrappers"

// ForeignDataWrapper represents a row from 'information_schema.foreign_data_wrappers'.
type ForeignDataWrapper struct {
	ForeignDataWrapperCatalog  SQLIdentifier `yaml:"foreign_data_wrapper_catalog,omitempty"`  // foreign_data_wrapper_catalog
	ForeignDataWrapperName     SQLIdentifier `yaml:"foreign_data_wrapper_name,omitempty"`     // foreign_data_wrapper_name
	AuthorizationIdentifier    SQLIdentifier `yaml:"authorization_identifier,omitempty"`      // authorization_identifier
	LibraryName                CharacterData `yaml:"library_name,omitempty"`                  // library_name
	ForeignDataWrapperLanguage CharacterData `yaml:"foreign_data_wrapper_language,omitempty"` // foreign_data_wrapper_language
}

// Constants defining each column in the table.
const (
	ForeignDataWrapperForeignDataWrapperCatalogField  = "foreign_data_wrapper_catalog"
	ForeignDataWrapperForeignDataWrapperNameField     = "foreign_data_wrapper_name"
	ForeignDataWrapperAuthorizationIdentifierField    = "authorization_identifier"
	ForeignDataWrapperLibraryNameField                = "library_name"
	ForeignDataWrapperForeignDataWrapperLanguageField = "foreign_data_wrapper_language"
)

// WhereClauses for every type in ForeignDataWrapper.
var (
	ForeignDataWrapperForeignDataWrapperCatalogWhere  SQLIdentifierField = "foreign_data_wrapper_catalog"
	ForeignDataWrapperForeignDataWrapperNameWhere     SQLIdentifierField = "foreign_data_wrapper_name"
	ForeignDataWrapperAuthorizationIdentifierWhere    SQLIdentifierField = "authorization_identifier"
	ForeignDataWrapperLibraryNameWhere                CharacterDataField = "library_name"
	ForeignDataWrapperForeignDataWrapperLanguageWhere CharacterDataField = "foreign_data_wrapper_language"
)

// QueryOneForeignDataWrapper retrieves a row from 'information_schema.foreign_data_wrappers' as a ForeignDataWrapper.
func QueryOneForeignDataWrapper(db XODB, where WhereClause, order OrderBy) (*ForeignDataWrapper, error) {
	const origsqlstr = `SELECT ` +
		`foreign_data_wrapper_catalog, foreign_data_wrapper_name, authorization_identifier, library_name, foreign_data_wrapper_language ` +
		`FROM information_schema.foreign_data_wrappers WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	fdw := &ForeignDataWrapper{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&fdw.ForeignDataWrapperCatalog, &fdw.ForeignDataWrapperName, &fdw.AuthorizationIdentifier, &fdw.LibraryName, &fdw.ForeignDataWrapperLanguage)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return fdw, nil
}

// QueryForeignDataWrapper retrieves rows from 'information_schema.foreign_data_wrappers' as a slice of ForeignDataWrapper.
func QueryForeignDataWrapper(db XODB, where WhereClause, order OrderBy) ([]*ForeignDataWrapper, error) {
	const origsqlstr = `SELECT ` +
		`foreign_data_wrapper_catalog, foreign_data_wrapper_name, authorization_identifier, library_name, foreign_data_wrapper_language ` +
		`FROM information_schema.foreign_data_wrappers WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*ForeignDataWrapper
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		fdw := ForeignDataWrapper{}

		err = q.Scan(&fdw.ForeignDataWrapperCatalog, &fdw.ForeignDataWrapperName, &fdw.AuthorizationIdentifier, &fdw.LibraryName, &fdw.ForeignDataWrapperLanguage)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &fdw)
	}
	return vals, nil
}

// AllForeignDataWrapper retrieves all rows from 'information_schema.foreign_data_wrappers' as a slice of ForeignDataWrapper.
func AllForeignDataWrapper(db XODB, order OrderBy) ([]*ForeignDataWrapper, error) {
	const origsqlstr = `SELECT ` +
		`foreign_data_wrapper_catalog, foreign_data_wrapper_name, authorization_identifier, library_name, foreign_data_wrapper_language ` +
		`FROM information_schema.foreign_data_wrappers`

	sqlstr := origsqlstr + order.String()

	var vals []*ForeignDataWrapper
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		fdw := ForeignDataWrapper{}

		err = q.Scan(&fdw.ForeignDataWrapperCatalog, &fdw.ForeignDataWrapperName, &fdw.AuthorizationIdentifier, &fdw.LibraryName, &fdw.ForeignDataWrapperLanguage)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &fdw)
	}
	return vals, nil
}
