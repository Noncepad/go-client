package basic

import (
	"errors"

	pbjob "github.com/noncepad/client/proto/job"
)

type Job struct {
	Id     string
	Status string
	Data   string
	Type   string
}

type JobStatusArgs struct {
	Id string `json:"id"`
}

func typeToString(jobType pbjob.JobType) (string, error) {
	switch jobType {
	case pbjob.JobType_VALIDATOR_INSTANCE:
		return "validator_instance", nil
	default:
		return "", errors.New("unknown job type")
	}
}

func statusToString(jobStatus pbjob.Status) (string, error) {

	switch jobStatus {
	case pbjob.Status_NEW:
		return "new", nil
	case pbjob.Status_STARTED:
		return "started", nil
	case pbjob.Status_FAILED:
		return "failed", nil
	case pbjob.Status_FINISHED:
		return "finished", nil
	default:
		return "", errors.New("unknown status")
	}
}

func jobToResult(job *pbjob.Job) (*Job, error) {
	s := job.GetStatus()
	if s == nil {
		return nil, errors.New("no status")
	}
	status, err := statusToString(s.Status)
	if err != nil {
		return nil, err
	}
	jobType, err := typeToString(job.Type)
	if err != nil {
		return nil, err
	}
	data := string(job.Data[:])
	return &Job{
		Id: job.Id, Status: status, Type: jobType, Data: data,
	}, nil
}

func (e1 Basic) JobStatus(args JobStatusArgs, results *Job) error {
	ctx, err := e1.Ctx()
	if err != nil {
		return err
	}
	ans, err := e1.jobClient.GetStatus(ctx, &pbjob.StatusRequest{Id: args.Id})
	if err != nil {
		return err
	}

	j, err := jobToResult(ans)
	if err != nil {
		return err
	}

	*results = *j
	return nil
}
