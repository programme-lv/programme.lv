package main

import (
	"github.com/KrisjanisP/deikstra/service/protofiles"
	"log"
)

type ResStream protofiles.Scheduler_ReportTaskEvalStatusClient

func evaluateTaskSubmission(job *protofiles.TaskEvalJob, resStream ResStream) error {
	log.Println("evaluating task submission ", job.GetJobId())

	if err := resStream.Send(&protofiles.TaskEvalStatus{JobId: job.GetJobId()}); err != nil {
		return err
	}
	return nil
}
