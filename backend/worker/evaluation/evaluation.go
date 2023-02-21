package evaluation

import (
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"github.com/KrisjanisP/deikstra/service/worker/constants"
	"log"
	"strings"
)

type ResStream pb.Scheduler_ReportTaskEvalStatusClient

func (e *Service) EvaluateTaskSubmission(job *pb.TaskEvalJob, resStream ResStream) (err error) {
	log.Println("evaluating task submission ", job.GetJobId())

	task, err := e.DownloadTask(job.GetJobId())
	if err != nil {
		return
	}

	err = resStream.Send(NewEvalStatus(job.GetJobId(), constants.EvalICS, 0))
	if err != nil {
		return
	}

	executable, _, err := NewExecutable(job.GetSrcCode(), job.GetLangId())
	if err != nil {
		return
	}

	err = resStream.Send(NewEvalStatus(job.GetJobId(), constants.EvalITS, 0))
	if err != nil {
		return
	}

	for _, test := range task.Tests {
		res, err := executable.Execute(strings.NewReader(test.Input))
		if err != nil {
			return err
		}
		log.Printf("test exec result %+v", res)
		if res.Stdout == test.Answer {
			err = resStream.Send(NewTestStatus(job.GetJobId(), test.ID, main.testOK, res.Stdout, res.Stderr))
		} else {
			err = resStream.Send(NewTestStatus(job.GetJobId(), test.ID, main.testWA, res.Stdout, res.Stderr))
		}
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
