package evaluation

import (
	"github.com/KrisjanisP/deikstra/service/models"
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"github.com/KrisjanisP/deikstra/service/worker/constants"
	"github.com/KrisjanisP/deikstra/service/worker/executable"
	"log"
	"strings"
)

type ResStream pb.Scheduler_ReportTaskEvalStatusClient

func (e *Service) EvaluateTaskSubmission(job *pb.TaskEvalJob, resStream ResStream) (err error) {
	log.Println("evaluating task submission ", job.GetJobId())

	var task *models.Task
	task, err = e.DownloadTask(job.GetJobId())
	if err != nil {
		return
	}

	var srcCode *executable.SrcCode
	srcCode, err = e.DownloadSourceCode(job.GetJobId())
	if err != nil {
		return
	}

	err = resStream.Send(NewEvalStatus(job.GetJobId(), constants.EvalICS, 0))
	if err != nil {
		return
	}

	box, err := e.isolateController.NewIsolateBox()
	if err != nil {
		return
	}

	defer box.Release()

	exe, compileOutput, err := executable.NewIsolatedExecutable(box, srcCode)
	if err != nil {
		return
	}

	err = resStream.Send(NewCompStatus(job.GetJobId(), compileOutput))
	err = resStream.Send(NewEvalStatus(job.GetJobId(), constants.EvalITS, 0))
	if err != nil {
		return
	}

	for _, test := range task.Tests {
		res, err := exe.Execute(strings.NewReader(test.Input))
		if err != nil {
			return err
		}
		log.Printf("test exec result %+v", res)
		if res.Stdout == test.Answer {
			err = resStream.Send(NewTestStatus(job.GetJobId(), test.ID, constants.TestOK, res.Stdout, res.Stderr))
		} else {
			err = resStream.Send(NewTestStatus(job.GetJobId(), test.ID, constants.TestWA, res.Stdout, res.Stderr))
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
