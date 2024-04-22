package model

type Campaign struct {
	ID                 uint
	Name               string
	OrganizerName      string
	TargetAmount       float64
	DonationMinimum    float64
	AmountDonatedSoFar float64
	Donations          []Donation
	Account            Account
}
