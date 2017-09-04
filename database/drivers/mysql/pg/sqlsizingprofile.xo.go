// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// SQLSizingProfileTable is the database name for the table.
const SQLSizingProfileTable = "information_schema.sql_sizing_profiles"

// SQLSizingProfile represents a row from 'information_schema.sql_sizing_profiles'.
type SQLSizingProfile struct {
	SizingID      CardinalNumber `yaml:"sizing_id,omitempty"`      // sizing_id
	SizingName    CharacterData  `yaml:"sizing_name,omitempty"`    // sizing_name
	ProfileID     CharacterData  `yaml:"profile_id,omitempty"`     // profile_id
	RequiredValue CardinalNumber `yaml:"required_value,omitempty"` // required_value
	Comments      CharacterData  `yaml:"comments,omitempty"`       // comments
}

// Constants defining each column in the table.
const (
	SQLSizingProfileSizingIDField      = "sizing_id"
	SQLSizingProfileSizingNameField    = "sizing_name"
	SQLSizingProfileProfileIDField     = "profile_id"
	SQLSizingProfileRequiredValueField = "required_value"
	SQLSizingProfileCommentsField      = "comments"
)

// WhereClauses for every type in SQLSizingProfile.
var (
	SQLSizingProfileSizingIDWhere      CardinalNumberField = "sizing_id"
	SQLSizingProfileSizingNameWhere    CharacterDataField  = "sizing_name"
	SQLSizingProfileProfileIDWhere     CharacterDataField  = "profile_id"
	SQLSizingProfileRequiredValueWhere CardinalNumberField = "required_value"
	SQLSizingProfileCommentsWhere      CharacterDataField  = "comments"
)

// QueryOneSQLSizingProfile retrieves a row from 'information_schema.sql_sizing_profiles' as a SQLSizingProfile.
func QueryOneSQLSizingProfile(db XODB, where WhereClause, order OrderBy) (*SQLSizingProfile, error) {
	const origsqlstr = `SELECT ` +
		`sizing_id, sizing_name, profile_id, required_value, comments ` +
		`FROM information_schema.sql_sizing_profiles WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	ssp := &SQLSizingProfile{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&ssp.SizingID, &ssp.SizingName, &ssp.ProfileID, &ssp.RequiredValue, &ssp.Comments)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ssp, nil
}

// QuerySQLSizingProfile retrieves rows from 'information_schema.sql_sizing_profiles' as a slice of SQLSizingProfile.
func QuerySQLSizingProfile(db XODB, where WhereClause, order OrderBy) ([]*SQLSizingProfile, error) {
	const origsqlstr = `SELECT ` +
		`sizing_id, sizing_name, profile_id, required_value, comments ` +
		`FROM information_schema.sql_sizing_profiles WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*SQLSizingProfile
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		ssp := SQLSizingProfile{}

		err = q.Scan(&ssp.SizingID, &ssp.SizingName, &ssp.ProfileID, &ssp.RequiredValue, &ssp.Comments)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &ssp)
	}
	return vals, nil
}

// AllSQLSizingProfile retrieves all rows from 'information_schema.sql_sizing_profiles' as a slice of SQLSizingProfile.
func AllSQLSizingProfile(db XODB, order OrderBy) ([]*SQLSizingProfile, error) {
	const origsqlstr = `SELECT ` +
		`sizing_id, sizing_name, profile_id, required_value, comments ` +
		`FROM information_schema.sql_sizing_profiles`

	sqlstr := origsqlstr + order.String()

	var vals []*SQLSizingProfile
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		ssp := SQLSizingProfile{}

		err = q.Scan(&ssp.SizingID, &ssp.SizingName, &ssp.ProfileID, &ssp.RequiredValue, &ssp.Comments)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &ssp)
	}
	return vals, nil
}
