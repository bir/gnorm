// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// TriggeredUpdateColumnTable is the database name for the table.
const TriggeredUpdateColumnTable = "information_schema.triggered_update_columns"

// TriggeredUpdateColumn represents a row from 'information_schema.triggered_update_columns'.
type TriggeredUpdateColumn struct {
	TriggerCatalog     SQLIdentifier `yaml:"trigger_catalog,omitempty"`      // trigger_catalog
	TriggerSchema      SQLIdentifier `yaml:"trigger_schema,omitempty"`       // trigger_schema
	TriggerName        SQLIdentifier `yaml:"trigger_name,omitempty"`         // trigger_name
	EventObjectCatalog SQLIdentifier `yaml:"event_object_catalog,omitempty"` // event_object_catalog
	EventObjectSchema  SQLIdentifier `yaml:"event_object_schema,omitempty"`  // event_object_schema
	EventObjectTable   SQLIdentifier `yaml:"event_object_table,omitempty"`   // event_object_table
	EventObjectColumn  SQLIdentifier `yaml:"event_object_column,omitempty"`  // event_object_column
}

// Constants defining each column in the table.
const (
	TriggeredUpdateColumnTriggerCatalogField     = "trigger_catalog"
	TriggeredUpdateColumnTriggerSchemaField      = "trigger_schema"
	TriggeredUpdateColumnTriggerNameField        = "trigger_name"
	TriggeredUpdateColumnEventObjectCatalogField = "event_object_catalog"
	TriggeredUpdateColumnEventObjectSchemaField  = "event_object_schema"
	TriggeredUpdateColumnEventObjectTableField   = "event_object_table"
	TriggeredUpdateColumnEventObjectColumnField  = "event_object_column"
)

// WhereClauses for every type in TriggeredUpdateColumn.
var (
	TriggeredUpdateColumnTriggerCatalogWhere     SQLIdentifierField = "trigger_catalog"
	TriggeredUpdateColumnTriggerSchemaWhere      SQLIdentifierField = "trigger_schema"
	TriggeredUpdateColumnTriggerNameWhere        SQLIdentifierField = "trigger_name"
	TriggeredUpdateColumnEventObjectCatalogWhere SQLIdentifierField = "event_object_catalog"
	TriggeredUpdateColumnEventObjectSchemaWhere  SQLIdentifierField = "event_object_schema"
	TriggeredUpdateColumnEventObjectTableWhere   SQLIdentifierField = "event_object_table"
	TriggeredUpdateColumnEventObjectColumnWhere  SQLIdentifierField = "event_object_column"
)

// QueryOneTriggeredUpdateColumn retrieves a row from 'information_schema.triggered_update_columns' as a TriggeredUpdateColumn.
func QueryOneTriggeredUpdateColumn(db XODB, where WhereClause, order OrderBy) (*TriggeredUpdateColumn, error) {
	const origsqlstr = `SELECT ` +
		`trigger_catalog, trigger_schema, trigger_name, event_object_catalog, event_object_schema, event_object_table, event_object_column ` +
		`FROM information_schema.triggered_update_columns WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	tuc := &TriggeredUpdateColumn{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&tuc.TriggerCatalog, &tuc.TriggerSchema, &tuc.TriggerName, &tuc.EventObjectCatalog, &tuc.EventObjectSchema, &tuc.EventObjectTable, &tuc.EventObjectColumn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tuc, nil
}

// QueryTriggeredUpdateColumn retrieves rows from 'information_schema.triggered_update_columns' as a slice of TriggeredUpdateColumn.
func QueryTriggeredUpdateColumn(db XODB, where WhereClause, order OrderBy) ([]*TriggeredUpdateColumn, error) {
	const origsqlstr = `SELECT ` +
		`trigger_catalog, trigger_schema, trigger_name, event_object_catalog, event_object_schema, event_object_table, event_object_column ` +
		`FROM information_schema.triggered_update_columns WHERE (`

	
	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*TriggeredUpdateColumn
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		tuc := TriggeredUpdateColumn{}

		err = q.Scan(&tuc.TriggerCatalog, &tuc.TriggerSchema, &tuc.TriggerName, &tuc.EventObjectCatalog, &tuc.EventObjectSchema, &tuc.EventObjectTable, &tuc.EventObjectColumn)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &tuc)
	}
	return vals, nil
}

// AllTriggeredUpdateColumn retrieves all rows from 'information_schema.triggered_update_columns' as a slice of TriggeredUpdateColumn.
func AllTriggeredUpdateColumn(db XODB, order OrderBy) ([]*TriggeredUpdateColumn, error) {
	const origsqlstr = `SELECT ` +
		`trigger_catalog, trigger_schema, trigger_name, event_object_catalog, event_object_schema, event_object_table, event_object_column ` +
		`FROM information_schema.triggered_update_columns`

	sqlstr := origsqlstr + order.String()

	var vals []*TriggeredUpdateColumn
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		tuc := TriggeredUpdateColumn{}

		err = q.Scan(&tuc.TriggerCatalog, &tuc.TriggerSchema, &tuc.TriggerName, &tuc.EventObjectCatalog, &tuc.EventObjectSchema, &tuc.EventObjectTable, &tuc.EventObjectColumn)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &tuc)
	}
	return vals, nil
}
