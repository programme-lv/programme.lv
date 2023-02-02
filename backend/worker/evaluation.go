package main

import (
	"github.com/KrisjanisP/deikstra/service/models"
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"gorm.io/gorm"
	"log"
	"strings"
)

type ResStream pb.Scheduler_ReportTaskEvalStatusClient

type EvaluationService struct {
	database *gorm.DB
}

func NewEvaluationService(database *gorm.DB) *EvaluationService {
	return &EvaluationService{database: database}
}

func (e *EvaluationService) EvaluateTaskSubmission(job *pb.TaskEvalJob, resStream ResStream) error {
	log.Println("evaluating task submission ", job.GetJobId())

	executable, err := NewExecutable(job.GetSrcCode(), job.GetLangId())
	if err != nil {
		return err
	}

	evaluation := models.TaskSubmEvaluation{}
	e.database.Preload("TaskSubmission.Task.Tests").First(&evaluation, job.GetJobId())

	task := evaluation.TaskSubmission.Task

	for _, test := range task.Tests {
		stdout, stderr, err := executable.Execute(strings.NewReader(test.Input))
		if err != nil {
			return err
		}

		err = resStream.Send(NewTestStatus(job.GetJobId(), test.ID, testOK, stdout, stderr))
		if err != nil {
			return err
		}

	}

	err = resStream.Send(NewTaskStatus(job.GetJobId(), pb.TaskEvalStatusCode_TE_OK, 100))
	if err != nil {
		return err
	}

	return nil
}

func NewTestStatus(jobId, testId uint64, testStatus pb.TaskTestStatusCode, stdout, stderr string) *pb.TaskEvalStatus {
	taskTestResult := pb.TaskTestResult{
		TestId: int32(testId), TestStatus: testStatus, Stdout: stdout, Stderr: stderr}
	taskTestStatus := pb.TaskEvalStatus_TestRes{TestRes: &taskTestResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &taskTestStatus}
}

func NewTaskStatus(jobId uint64, evalStatus pb.TaskEvalStatusCode, score int32) *pb.TaskEvalStatus {
	taskEvalResult := pb.TaskEvalResult{EvalStatus: evalStatus, Score: score}
	taskEvalStatus := pb.TaskEvalStatus_TaskRes{TaskRes: &taskEvalResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &taskEvalStatus}
}
