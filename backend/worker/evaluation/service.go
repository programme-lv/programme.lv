package evaluation

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"gorm.io/gorm"
)

type Service struct {
	database *gorm.DB
}

func NewEvaluationService(database *gorm.DB) *Service {
	return &Service{database: database}
}

func (e *Service) DownloadTask(jobId uint64) (task *models.Task, err error) {
	evaluation := models.TaskSubmEvaluation{}
	err = e.database.Preload("TaskSubmission.Task.Tests").First(&evaluation, jobId).Error
	if err != nil {
		return
	}
	task = evaluation.TaskSubmission.Task
	return
}
