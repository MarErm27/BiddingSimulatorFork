package models

import (
	"github.com/MarErm27/uadmin"
)

// Participants model ...
type Participant struct {
	uadmin.Model
	Name    string `uadmin:"required;search;multilingual"`
	Virtual bool
}
