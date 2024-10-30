// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type answersTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *answersTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("answers").
func (v *answersTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *answersTableType) Columns() []string {
	return []string{
		"id",
		"text",
		"question_id",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *answersTableType) NewStruct() reform.Struct {
	return new(Answers)
}

// NewRecord makes a new record for that table.
func (v *answersTableType) NewRecord() reform.Record {
	return new(Answers)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *answersTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AnswersTable represents answers view or table in SQL database.
var AnswersTable = &answersTableType{
	s: parse.StructInfo{
		Type:    "Answers",
		SQLName: "answers",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int32", Column: "id"},
			{Name: "Text", Type: "string", Column: "text"},
			{Name: "QuestionID", Type: "*int32", Column: "question_id"},
		},
		PKFieldIndex: 0,
	},
	z: new(Answers).Values(),
}

// String returns a string representation of this struct or record.
func (s Answers) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Text: " + reform.Inspect(s.Text, true)
	res[2] = "QuestionID: " + reform.Inspect(s.QuestionID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Answers) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Text,
		s.QuestionID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Answers) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Text,
		&s.QuestionID,
	}
}

// View returns View object for that struct.
func (s *Answers) View() reform.View {
	return AnswersTable
}

// Table returns Table object for that record.
func (s *Answers) Table() reform.Table {
	return AnswersTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Answers) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Answers) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Answers) HasPK() bool {
	return s.ID != AnswersTable.z[AnswersTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Answers) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = AnswersTable
	_ reform.Struct = (*Answers)(nil)
	_ reform.Table  = AnswersTable
	_ reform.Record = (*Answers)(nil)
	_ fmt.Stringer  = (*Answers)(nil)
)

func init() {
	parse.AssertUpToDate(&AnswersTable.s, new(Answers))
}
