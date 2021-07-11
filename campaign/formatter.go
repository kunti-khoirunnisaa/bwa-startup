package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignformatter := CampaignFormatter{}
	campaignformatter.ID = campaign.ID
	campaignformatter.UserID = campaign.userID
	campaignformatter.Name = campaign.Name
	campaignformatter.ShortDescription = campaign.ShortDescription
	campaignformatter.GoalAmount = campaign.GoalAmount
	campaignformatter.CurrentAmount = campaign.CurrentAmount

	if len(campaign.CampaignImages) > 0 {
		campaignformatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignformatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
