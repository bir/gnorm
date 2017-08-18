// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// SQLFeatureTable is the database name for the table.
const SQLFeatureTable = "information_schema.sql_features"

// SQLFeature represents a row from 'information_schema.sql_features'.
type SQLFeature struct {
	FeatureID      CharacterData `json:"feature_id"`       // feature_id
	FeatureName    CharacterData `json:"feature_name"`     // feature_name
	SubFeatureID   CharacterData `json:"sub_feature_id"`   // sub_feature_id
	SubFeatureName CharacterData `json:"sub_feature_name"` // sub_feature_name
	IsSupported    YesOrNo       `json:"is_supported"`     // is_supported
	IsVerifiedBy   CharacterData `json:"is_verified_by"`   // is_verified_by
	Comments       CharacterData `json:"comments"`         // comments
}

// Constants defining each column in the table.
const (
	SQLFeatureFeatureIDField      = "feature_id"
	SQLFeatureFeatureNameField    = "feature_name"
	SQLFeatureSubFeatureIDField   = "sub_feature_id"
	SQLFeatureSubFeatureNameField = "sub_feature_name"
	SQLFeatureIsSupportedField    = "is_supported"
	SQLFeatureIsVerifiedByField   = "is_verified_by"
	SQLFeatureCommentsField       = "comments"
)

// WhereClauses for every type in SQLFeature.
var (
	SQLFeatureFeatureIDWhere      CharacterDataField = "feature_id"
	SQLFeatureFeatureNameWhere    CharacterDataField = "feature_name"
	SQLFeatureSubFeatureIDWhere   CharacterDataField = "sub_feature_id"
	SQLFeatureSubFeatureNameWhere CharacterDataField = "sub_feature_name"
	SQLFeatureIsSupportedWhere    YesOrNoField       = "is_supported"
	SQLFeatureIsVerifiedByWhere   CharacterDataField = "is_verified_by"
	SQLFeatureCommentsWhere       CharacterDataField = "comments"
)

// QueryOneSQLFeature retrieves a row from 'information_schema.sql_features' as a SQLFeature.
func QueryOneSQLFeature(db XODB, where WhereClause, order OrderBy) (*SQLFeature, error) {
	const origsqlstr = `SELECT ` +
		`feature_id, feature_name, sub_feature_id, sub_feature_name, is_supported, is_verified_by, comments ` +
		`FROM information_schema.sql_features WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	sf := &SQLFeature{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&sf.FeatureID, &sf.FeatureName, &sf.SubFeatureID, &sf.SubFeatureName, &sf.IsSupported, &sf.IsVerifiedBy, &sf.Comments)
	if err != nil {
		return nil, err
	}
	return sf, nil
}

// QuerySQLFeature retrieves rows from 'information_schema.sql_features' as a slice of SQLFeature.
func QuerySQLFeature(db XODB, where WhereClause, order OrderBy) ([]*SQLFeature, error) {
	const origsqlstr = `SELECT ` +
		`feature_id, feature_name, sub_feature_id, sub_feature_name, is_supported, is_verified_by, comments ` +
		`FROM information_schema.sql_features WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*SQLFeature
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		sf := SQLFeature{}

		err = q.Scan(&sf.FeatureID, &sf.FeatureName, &sf.SubFeatureID, &sf.SubFeatureName, &sf.IsSupported, &sf.IsVerifiedBy, &sf.Comments)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &sf)
	}
	return vals, nil
}