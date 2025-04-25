package models

import (
	"gorm.io/datatypes"
	"time"
)

type ScanFileInfo struct {
	ScanId       string    `json:"scan_id"`
	Timestamp    time.Time `json:"timestamp"`
	ScanStatus   string    `json:"scan_status"`
	ResourceType string    `json:"resource_type"`
	ResourceName string    `json:"resource_name"`
	FileName     string    `json:"file_name"`
}

type Vulnerability struct {
	Id             string                      `gorm:"primaryKey" json:"id"`
	Severity       string                      `gorm:"size:20" json:"severity"`
	Cvss           float64                     `json:"cvss"`
	Status         string                      `gorm:"size:30" json:"status"`
	PackageName    string                      `gorm:"size:50" json:"package_name"`
	CurrentVersion string                      `gorm:"size:20" json:"current_version"`
	FixedVersion   string                      `gorm:"size:20" json:"fixed_version"`
	Description    string                      `gorm:"size:90" json:"description"`
	PublishedDate  time.Time                   `json:"published_date"`
	Link           string                      `json:"link"`
	RiskFactors    datatypes.JSONSlice[string] `gorm:"type:json" json:"risk_factors"`
	Metadata       ScanFileInfo                `gorm:"embedded"`
}
