package main

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/protofiles"
	"gorm.io/gorm"
	"log"
)

type ResStream protofiles.Scheduler_ReportTaskEvalStatusClient

type EvaluationService struct {
	database *gorm.DB
}

func NewEvaluationService(database *gorm.DB) *EvaluationService {
	return &EvaluationService{database: database}
}

func (e *EvaluationService) EvaluateTaskSubmission(job *protofiles.TaskEvalJob, resStream ResStream) error {
	log.Println("evaluating task submission ", job.GetJobId())

	if err := resStream.Send(&protofiles.TaskEvalStatus{JobId: job.GetJobId()}); err != nil {
		return err
	}

	evaluation := models.TaskSubmEvaluation{}
	e.database.Preload("TaskSubmission.Task.Tests").First(&evaluation, job.GetJobId())
	log.Printf("evaluation: %+v", evaluation)

	tests := evaluation.TaskSubmission.Task.Tests
	for _, test := range tests {
		log.Println("test:", test.ID)
	}

	taskEvalResult := protofiles.TaskEvalResult{EvalStatus: protofiles.TaskEvalStatusCode_TE_OK, Score: 100}
	taskEvalStatus := protofiles.TaskEvalStatus_TaskRes{TaskRes: &taskEvalResult}
	err := resStream.Send(&protofiles.TaskEvalStatus{JobId: job.GetJobId(), Status: &taskEvalStatus})
	if err != nil {
		return err
	}
	return nil
}
