package services

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

	"services-work-team-update/database"
	"services-work-team-update/models"
)

// UpdateWorkTeamService updates a work team based on teamCode
func UpdateWorkTeamService(teamCode string, updatedTeam *models.WorkTeam) (*models.WorkTeam, error) {
	var existingTeam models.WorkTeam

	if err := database.DB.Where("team_code = ?", teamCode).First(&existingTeam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("team with teamCode %s not found", teamCode)
		}
		return nil, fmt.Errorf("error finding team: %w", err)
	}

	existingTeam.TeamName = updatedTeam.TeamName
	existingTeam.Description = updatedTeam.Description
	existingTeam.Members = pq.StringArray(updatedTeam.Members)
	existingTeam.Status = true
	existingTeam.UpdatedAt = time.Now()

	if err := database.DB.Save(&existingTeam).Error; err != nil {
		return nil, fmt.Errorf("error updating team: %w", err)
	}

	return &existingTeam, nil
}
