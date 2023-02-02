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

func (e *EvaluationService) EvaluateTaskSubmission(job *pb.TaskEvalJob, resStream ResStream) (err error) {
	log.Println("evaluating task submission ", job.GetJobId())

	evaluation := models.TaskSubmEvaluation{}
	err = e.database.Preload("TaskSubmission.Task.Tests").First(&evaluation, job.GetJobId()).Error
	if err != nil {
		return
	}
	task := evaluation.TaskSubmission.Task

	err = resStream.Send(NewEvalStatus(job.GetJobId(), evalICS, 0))
	if err != nil {
		return
	}

	executable, _, err := NewExecutable(job.GetSrcCode(), job.GetLangId())
	if err != nil {
		return
	}

	err = resStream.Send(NewEvalStatus(job.GetJobId(), evalITS, 0))
	if err != nil {
		return
	}

	for _, test := range task.Tests {
		res, err := executable.Execute(strings.NewReader(test.Input))
		if err != nil {
			return err
		}
		log.Println("stdout: ", res.Stdout)
		err = resStream.Send(NewTestStatus(job.GetJobId(), test.ID, testOK, res.Stdout, res.Stderr))
		if err != nil {
			return err
		}

	}

	err = resStream.Send(NewEvalStatus(job.GetJobId(), pb.TaskEvalStatusCode_TE_OK, 100))
	if err != nil {
		return
	}

	return nil
}

func NewTestStatus(jobId, testId uint64, testStatus pb.TaskTestStatusCode, stdout, stderr string) *pb.TaskEvalStatus {
	taskTestResult := pb.TaskTestResult{
		TestId: int32(testId), TestStatus: testStatus, Stdout: stdout, Stderr: stderr}
	taskTestStatus := pb.TaskEvalStatus_TestRes{TestRes: &taskTestResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &taskTestStatus}
}

func NewEvalStatus(jobId uint64, evalStatus pb.TaskEvalStatusCode, score int32) *pb.TaskEvalStatus {
	taskEvalResult := pb.TaskEvalResult{EvalStatus: evalStatus, Score: score}
	taskEvalStatus := pb.TaskEvalStatus_TaskRes{TaskRes: &taskEvalResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &taskEvalStatus}
}
