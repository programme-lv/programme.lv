package evaluation

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/worker/executable"
	"github.com/KrisjanisP/deikstra/service/worker/execution"
	"gorm.io/gorm"
)

type Service struct {
	database          *gorm.DB
	isolateController *execution.IsolateController
}

func NewEvaluationService(database *gorm.DB, isolateBoxes int) *Service {
	return &Service{database: database, isolateController: execution.NewIsolateController(isolateBoxes)}
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

func (e *Service) DownloadSourceCode(jobId uint64) (srcCode *executable.SrcCode, err error) {

	return
}
