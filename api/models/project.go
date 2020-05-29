package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Project struct {
	ID          int64      `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	URL         NullString `db:"url" json:"url"`
	Description NullString `db:"description" json:"description"`
	Category    string     `db:"category" json:"category"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updatedAt"`
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}
