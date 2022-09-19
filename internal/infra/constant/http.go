package constants

import (
	"net/http"

	"schneider.vip/problem"
)

const (
	SampleRessource string = "/sample"
)

var (
	ProblemInvalidContract = problem.New(
		problem.Title("INVALID_CONTRACT"),
		problem.Detail("Invalid Payload"),
		problem.Status(http.StatusBadRequest),
	)

	ProblemInternalServerError = problem.New(problem.Title(http.StatusText(http.StatusInternalServerError)),
		problem.Status(http.StatusInternalServerError),
		problem.Detail("The server encountered an unexpected condition which prevented it for fulfilling the request."),
	)
)
