package evaluation

import (
	pb "github.com/KrisjanisP/deikstra/service/protofiles"
	"github.com/KrisjanisP/deikstra/service/worker/execution"
)

func NewTestStatus(jobId, testId uint64, testStatus pb.TaskTestStatusCode, stdout, stderr string) *pb.TaskEvalStatus {
	taskTestResult := pb.TaskTestResult{
		TestId: int32(testId), TestStatus: testStatus, Stdout: stdout, Stderr: stderr}
	taskTestStatus := pb.TaskEvalStatus_TestRes{TestRes: &taskTestResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &taskTestStatus}
}

func NewEvalStatus(jobId uint64, evalStatus pb.TaskEvalStatusCode, score int32) *pb.TaskEvalStatus {
	taskEvalResult := pb.TaskEvalResult{EvalStatus: evalStatus, Score: score}
	taskEvalStatus := pb.TaskEvalStatus_EvalRes{EvalRes: &taskEvalResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &taskEvalStatus}
}

func NewCompStatus(jobId uint64, result *execution.ExtendedResult) *pb.TaskEvalStatus {
	submCompResult := pb.CompilationRes{Stdout: result.Stdout, Stderr: result.Stderr, ExitCode: int64(result.ExitCode)}
	submCompStatus := pb.TaskEvalStatus_CompRes{CompRes: &submCompResult}
	return &pb.TaskEvalStatus{JobId: jobId, Status: &submCompStatus}
}
