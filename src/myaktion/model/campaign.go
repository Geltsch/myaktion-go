package model

import "gorm.io/gorm"

/*type Campaign struct {
	ID                 uint
	Name               string
	OrganizerName      string
	TargetAmount       float64
	DonationMinimum    float64
	AmountDonatedSoFar float64
	Donations          []Donation
	Account            Account
}*/

type Campaign struct {
	gorm.Model
	Name               string     `gorm:"notNull;size:30"`
	OrganizerName      string     `gorm:"notNull"`
	TargetAmount       float64    `gorm:"notNull;check:target_amount >= 10.0"`
	DonationMinimum    float64    `gorm:"notNull;check:donation_minimum >= 1.0"`
	Donations          []Donation `gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AmountDonatedSoFar float64    `gorm:"-"`
	Account            Account    `gorm:"embedded;embeddedPrefix:account_"`
}
