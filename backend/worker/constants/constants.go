package constants

import "github.com/KrisjanisP/deikstra/service/protofiles"

const (
	TestOK  = protofiles.TaskTestStatusCode_TT_OK
	TestWA  = protofiles.TaskTestStatusCode_TT_WA
	EvalICS = protofiles.TaskEvalStatusCode_TE_ICS
	EvalITS = protofiles.TaskEvalStatusCode_TE_ITS
)
