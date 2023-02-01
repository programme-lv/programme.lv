package main

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/protofiles"
	"gorm.io/gorm"
	"log"
	"strings"
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

	executable, err := NewExecutable(job.GetSrcCode(), job.GetLangId())
	if err != nil {
		return err
	}

	evaluation := models.TaskSubmEvaluation{}
	e.database.Preload("TaskSubmission.Task.Tests").First(&evaluation, job.GetJobId())
	log.Printf("evaluation: %+v", evaluation)

	tests := evaluation.TaskSubmission.Task.Tests
	for _, test := range tests {
		stdout, stderr, err := executable.Execute(strings.NewReader(test.Input))
		if err != nil {
			return err
		}
		log.Println("test:", test.ID, stdout, stderr)

		taskTestResult := protofiles.TaskTestResult{TestId: int32(test.ID), TestStatus: protofiles.TaskTestStatusCode_TT_OK, Stdout: stdout, Stderr: stderr}
		taskTestStatus := protofiles.TaskEvalStatus_TestRes{TestRes: &taskTestResult}
		err = resStream.Send(&protofiles.TaskEvalStatus{JobId: job.GetJobId(), Status: &taskTestStatus})
		if err != nil {
			return err
		}

	}

	taskEvalResult := protofiles.TaskEvalResult{EvalStatus: protofiles.TaskEvalStatusCode_TE_OK, Score: 100}
	taskEvalStatus := protofiles.TaskEvalStatus_TaskRes{TaskRes: &taskEvalResult}
	err = resStream.Send(&protofiles.TaskEvalStatus{JobId: job.GetJobId(), Status: &taskEvalStatus})
	if err != nil {
		return err
	}

	return nil
}
