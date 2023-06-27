package models

import uuid "github.com/google/uuid"

// State struct defining the state model
type State struct {
	ExternalID uuid.UUID `gorm:"column:external_id;type:uuid;default:gen_random_uuid();primarykey" json:"external_id"`
	Name       string    `gorm:"column:name;unique" json:"name" binding:"required"`
}

// District struct defining the district model
type District struct {
	ExternalID      uuid.UUID `gorm:"column:external_id;type:uuid;default:gen_random_uuid();primarykey" json:"external_id"`
	Name            string    `gorm:"column:name;unique" json:"name" binding:"required"`
	StateExternalID uuid.UUID `gorm:"column:state_external_id;type:uuid" json:"state_external_id"`
	State           State     `gorm:"foreignkey:StateExternalID" json:"state"`
}
