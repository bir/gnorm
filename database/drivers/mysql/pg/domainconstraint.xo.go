// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// DomainConstraintTable is the database name for the table.
const DomainConstraintTable = "information_schema.domain_constraints"

// DomainConstraint represents a row from 'information_schema.domain_constraints'.
type DomainConstraint struct {
	ConstraintCatalog SQLIdentifier `yaml:"constraint_catalog,omitempty"` // constraint_catalog
	ConstraintSchema  SQLIdentifier `yaml:"constraint_schema,omitempty"`  // constraint_schema
	ConstraintName    SQLIdentifier `yaml:"constraint_name,omitempty"`    // constraint_name
	DomainCatalog     SQLIdentifier `yaml:"domain_catalog,omitempty"`     // domain_catalog
	DomainSchema      SQLIdentifier `yaml:"domain_schema,omitempty"`      // domain_schema
	DomainName        SQLIdentifier `yaml:"domain_name,omitempty"`        // domain_name
	IsDeferrable      YesOrNo       `yaml:"is_deferrable,omitempty"`      // is_deferrable
	InitiallyDeferred YesOrNo       `yaml:"initially_deferred,omitempty"` // initially_deferred
}

// Constants defining each column in the table.
const (
	DomainConstraintConstraintCatalogField = "constraint_catalog"
	DomainConstraintConstraintSchemaField  = "constraint_schema"
	DomainConstraintConstraintNameField    = "constraint_name"
	DomainConstraintDomainCatalogField     = "domain_catalog"
	DomainConstraintDomainSchemaField      = "domain_schema"
	DomainConstraintDomainNameField        = "domain_name"
	DomainConstraintIsDeferrableField      = "is_deferrable"
	DomainConstraintInitiallyDeferredField = "initially_deferred"
)

// WhereClauses for every type in DomainConstraint.
var (
	DomainConstraintConstraintCatalogWhere SQLIdentifierField = "constraint_catalog"
	DomainConstraintConstraintSchemaWhere  SQLIdentifierField = "constraint_schema"
	DomainConstraintConstraintNameWhere    SQLIdentifierField = "constraint_name"
	DomainConstraintDomainCatalogWhere     SQLIdentifierField = "domain_catalog"
	DomainConstraintDomainSchemaWhere      SQLIdentifierField = "domain_schema"
	DomainConstraintDomainNameWhere        SQLIdentifierField = "domain_name"
	DomainConstraintIsDeferrableWhere      YesOrNoField       = "is_deferrable"
	DomainConstraintInitiallyDeferredWhere YesOrNoField       = "initially_deferred"
)

// QueryOneDomainConstraint retrieves a row from 'information_schema.domain_constraints' as a DomainConstraint.
func QueryOneDomainConstraint(db XODB, where WhereClause, order OrderBy) (*DomainConstraint, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, domain_catalog, domain_schema, domain_name, is_deferrable, initially_deferred ` +
		`FROM information_schema.domain_constraints WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	dc := &DomainConstraint{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&dc.ConstraintCatalog, &dc.ConstraintSchema, &dc.ConstraintName, &dc.DomainCatalog, &dc.DomainSchema, &dc.DomainName, &dc.IsDeferrable, &dc.InitiallyDeferred)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return dc, nil
}

// QueryDomainConstraint retrieves rows from 'information_schema.domain_constraints' as a slice of DomainConstraint.
func QueryDomainConstraint(db XODB, where WhereClause, order OrderBy) ([]*DomainConstraint, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, domain_catalog, domain_schema, domain_name, is_deferrable, initially_deferred ` +
		`FROM information_schema.domain_constraints WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*DomainConstraint
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		dc := DomainConstraint{}

		err = q.Scan(&dc.ConstraintCatalog, &dc.ConstraintSchema, &dc.ConstraintName, &dc.DomainCatalog, &dc.DomainSchema, &dc.DomainName, &dc.IsDeferrable, &dc.InitiallyDeferred)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &dc)
	}
	return vals, nil
}

// AllDomainConstraint retrieves all rows from 'information_schema.domain_constraints' as a slice of DomainConstraint.
func AllDomainConstraint(db XODB, order OrderBy) ([]*DomainConstraint, error) {
	const origsqlstr = `SELECT ` +
		`constraint_catalog, constraint_schema, constraint_name, domain_catalog, domain_schema, domain_name, is_deferrable, initially_deferred ` +
		`FROM information_schema.domain_constraints`

	sqlstr := origsqlstr + order.String()

	var vals []*DomainConstraint
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		dc := DomainConstraint{}

		err = q.Scan(&dc.ConstraintCatalog, &dc.ConstraintSchema, &dc.ConstraintName, &dc.DomainCatalog, &dc.DomainSchema, &dc.DomainName, &dc.IsDeferrable, &dc.InitiallyDeferred)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &dc)
	}
	return vals, nil
}
