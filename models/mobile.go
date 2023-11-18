package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Mobile represents the structure of the mobile data
type Mobile struct {
	ID          uint     `gorm:"primary_key" json:"id"`
	Name        string   `json:"name"`
	Company     string   `json:"company"`
	Price       int      `json:"price"`
	Colors      Colors   `gorm:"type:json" json:"colors"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Featured    bool     `json:"featured,omitempty"`
	Shipping    bool     `json:"shipping,omitempty"`
}

// Colors is a custom type to handle the Colors field in the database
type Colors []string

// Value implements the driver Valuer interface
func (c Colors) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan implements the sql Scanner interface
func (c *Colors) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var data string
	switch v := value.(type) {
	case []byte:
		data = string(v)
	case string:
		data = v
	default:
		return errors.New("unsupported Colors type")
	}

	return json.Unmarshal([]byte(data), c)
}
