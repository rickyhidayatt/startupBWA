package campaign

import "time"

// langkah pertama kita bikin entity buat nampung database campaign dan images

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdateAt         time.Time
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsiPrimary int
	CreatedAt  time.Time
	UpdateAt   time.Time
}
