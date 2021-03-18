package campaign

import "time"

type Campaign struct {
	ID               int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	perks            string
	BeckerCount      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
