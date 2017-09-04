{{ range (makeSlice "int" "hstore.Hstore" "Jsonb" "string" "sql.NullString" "int64" "sql.NullInt64" "float64" "sql.NullFloat64" "bool" "sql.NullBool" "time.Time" "pq.NullTime" "uuid.UUID" "uuid.NullUUID" "uint32" "YesOrNo") }}
{{ $fieldName := . | typeName | titleCase }}
// {{$fieldName}}Field is a component that returns a WhereClause that contains a
// comparison based on its field and a strongly typed value.
type {{$fieldName}}Field string

// Equals returns a WhereClause for this field.
func (f {{$fieldName}}Field) Equals(v {{.}}) WhereClause {
	return whereClause{
		field: string(f),
		comp:  compEqual,
		value: v,
	}
}

// GreaterThan returns a WhereClause for this field.
func (f {{$fieldName}}Field) GreaterThan(v {{.}}) WhereClause {
	return whereClause{
		field: string(f),
		comp:  compGreater,
		value: v,
	}
}

// LessThan returns a WhereClause for this field.
func (f {{$fieldName}}Field) LessThan(v {{.}}) WhereClause {
	return whereClause{
		field: string(f),
		comp:  compEqual,
		value: v,
	}
}

// GreaterOrEqual returns a WhereClause for this field.
func (f {{$fieldName}}Field) GreaterOrEqual(v {{.}}) WhereClause {
	return whereClause{
		field: string(f),
		comp:  compGTE,
		value: v,
	}
}

// LessOrEqual returns a WhereClause for this field.
func (f {{$fieldName}}Field) LessOrEqual(v {{.}}) WhereClause {
	return whereClause{
		field: string(f),
		comp:  compLTE,
		value: v,
	}
}

// NotEqual returns a WhereClause for this field.
func (f {{$fieldName}}Field) NotEqual(v {{.}}) WhereClause {
	return whereClause{
		field: string(f),
		comp:  compNE,
		value: v,
	}
}

// In returns a WhereClause for this field.
func (f {{$fieldName}}Field) In(vals []{{.}}) WhereClause {
	values := make([]interface{}, len(vals))
	for x := range vals {
		values[x] = vals[x]
	}
	return inClause{
		field: string(f),
		values: values,
	}
}

{{end}}
