package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/geltsch/myaktion-go/src/myaktion/db"
	"github.com/geltsch/myaktion-go/src/myaktion/model"
)

func AddDonation(campaignId uint, donation *model.Donation) error {
	donation.CampaignID = campaignId
	result := db.DB.Create(donation)
	if result.Error != nil {
		return result.Error
	}
	entry := log.WithField("ID", campaignId)
	entry.Info("Successfully added new donation to campaign in database.")
	entry.Tracef("Stored: %v", donation)
	return nil
}
