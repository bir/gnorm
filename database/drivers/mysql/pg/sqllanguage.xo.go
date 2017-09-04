// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// SQLLanguageTable is the database name for the table.
const SQLLanguageTable = "information_schema.sql_languages"

// SQLLanguage represents a row from 'information_schema.sql_languages'.
type SQLLanguage struct {
	SQLLanguageSource              CharacterData `yaml:"sql_language_source,omitempty"`               // sql_language_source
	SQLLanguageYear                CharacterData `yaml:"sql_language_year,omitempty"`                 // sql_language_year
	SQLLanguageConformance         CharacterData `yaml:"sql_language_conformance,omitempty"`          // sql_language_conformance
	SQLLanguageIntegrity           CharacterData `yaml:"sql_language_integrity,omitempty"`            // sql_language_integrity
	SQLLanguageImplementation      CharacterData `yaml:"sql_language_implementation,omitempty"`       // sql_language_implementation
	SQLLanguageBindingStyle        CharacterData `yaml:"sql_language_binding_style,omitempty"`        // sql_language_binding_style
	SQLLanguageProgrammingLanguage CharacterData `yaml:"sql_language_programming_language,omitempty"` // sql_language_programming_language
}

// Constants defining each column in the table.
const (
	SQLLanguageSQLLanguageSourceField              = "sql_language_source"
	SQLLanguageSQLLanguageYearField                = "sql_language_year"
	SQLLanguageSQLLanguageConformanceField         = "sql_language_conformance"
	SQLLanguageSQLLanguageIntegrityField           = "sql_language_integrity"
	SQLLanguageSQLLanguageImplementationField      = "sql_language_implementation"
	SQLLanguageSQLLanguageBindingStyleField        = "sql_language_binding_style"
	SQLLanguageSQLLanguageProgrammingLanguageField = "sql_language_programming_language"
)

// WhereClauses for every type in SQLLanguage.
var (
	SQLLanguageSQLLanguageSourceWhere              CharacterDataField = "sql_language_source"
	SQLLanguageSQLLanguageYearWhere                CharacterDataField = "sql_language_year"
	SQLLanguageSQLLanguageConformanceWhere         CharacterDataField = "sql_language_conformance"
	SQLLanguageSQLLanguageIntegrityWhere           CharacterDataField = "sql_language_integrity"
	SQLLanguageSQLLanguageImplementationWhere      CharacterDataField = "sql_language_implementation"
	SQLLanguageSQLLanguageBindingStyleWhere        CharacterDataField = "sql_language_binding_style"
	SQLLanguageSQLLanguageProgrammingLanguageWhere CharacterDataField = "sql_language_programming_language"
)

// QueryOneSQLLanguage retrieves a row from 'information_schema.sql_languages' as a SQLLanguage.
func QueryOneSQLLanguage(db XODB, where WhereClause, order OrderBy) (*SQLLanguage, error) {
	const origsqlstr = `SELECT ` +
		`sql_language_source, sql_language_year, sql_language_conformance, sql_language_integrity, sql_language_implementation, sql_language_binding_style, sql_language_programming_language ` +
		`FROM information_schema.sql_languages WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	sl := &SQLLanguage{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&sl.SQLLanguageSource, &sl.SQLLanguageYear, &sl.SQLLanguageConformance, &sl.SQLLanguageIntegrity, &sl.SQLLanguageImplementation, &sl.SQLLanguageBindingStyle, &sl.SQLLanguageProgrammingLanguage)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return sl, nil
}

// QuerySQLLanguage retrieves rows from 'information_schema.sql_languages' as a slice of SQLLanguage.
func QuerySQLLanguage(db XODB, where WhereClause, order OrderBy) ([]*SQLLanguage, error) {
	const origsqlstr = `SELECT ` +
		`sql_language_source, sql_language_year, sql_language_conformance, sql_language_integrity, sql_language_implementation, sql_language_binding_style, sql_language_programming_language ` +
		`FROM information_schema.sql_languages WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*SQLLanguage
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		sl := SQLLanguage{}

		err = q.Scan(&sl.SQLLanguageSource, &sl.SQLLanguageYear, &sl.SQLLanguageConformance, &sl.SQLLanguageIntegrity, &sl.SQLLanguageImplementation, &sl.SQLLanguageBindingStyle, &sl.SQLLanguageProgrammingLanguage)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &sl)
	}
	return vals, nil
}

// AllSQLLanguage retrieves all rows from 'information_schema.sql_languages' as a slice of SQLLanguage.
func AllSQLLanguage(db XODB, order OrderBy) ([]*SQLLanguage, error) {
	const origsqlstr = `SELECT ` +
		`sql_language_source, sql_language_year, sql_language_conformance, sql_language_integrity, sql_language_implementation, sql_language_binding_style, sql_language_programming_language ` +
		`FROM information_schema.sql_languages`

	sqlstr := origsqlstr + order.String()

	var vals []*SQLLanguage
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		sl := SQLLanguage{}

		err = q.Scan(&sl.SQLLanguageSource, &sl.SQLLanguageYear, &sl.SQLLanguageConformance, &sl.SQLLanguageIntegrity, &sl.SQLLanguageImplementation, &sl.SQLLanguageBindingStyle, &sl.SQLLanguageProgrammingLanguage)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &sl)
	}
	return vals, nil
}
