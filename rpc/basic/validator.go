package basic

import (
	"encoding/json"
	"errors"

	pbb "github.com/noncepad/client/proto/base"
	pbjob "github.com/noncepad/client/proto/job"
	pbsol "github.com/noncepad/client/proto/solana"
	"github.com/noncepad/client/util"
)

type ValidatorLSArgs struct {
}

type ValidatorInstance struct {
	Id       string `json:"id"`
	Hostv4   string `json:"hostv4"`
	Hostv6   string `json:"hostv6"`
	Port     uint32 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
type ValidatorLSResults struct {
	Instance []ValidatorInstance `json:"instance"`
}

func (e1 Basic) ValidatorLS(args ValidatorLSArgs, results *ValidatorLSResults) error {
	ctx, err := e1.Ctx()
	if err != nil {
		return err
	}
	ans, err := e1.solanaClient.ListValidator(ctx, &pbb.Empty{})
	if err != nil {
		return err
	}
	list := ans.GetInstance()

	a := make([]ValidatorInstance, len(list))
	for i := 0; i < len(list); i++ {
		a[i] = instanceToResult(list[i])
	}

	*results = ValidatorLSResults{Instance: a}

	return nil
}

type ValidatorCreateArgs struct {
	Type util.ValidatorType `json:"type"`
}

func (e1 Basic) ValidatorCreate(args ValidatorCreateArgs, results *Job) error {
	valType, err := util.ValidatorTypeStringToProto(args.Type)
	if err != nil {
		return err
	}
	ctx, err := e1.Ctx()
	if err != nil {
		return err
	}
	ans, err := e1.solanaClient.CreateValidator(ctx, &pbsol.ValidatorCreateRequest{
		Type: valType,
	})
	if err != nil {
		return err
	}

	*results = Job{Id: ans.Id, Status: "new", Data: "", Type: ""}
	return nil
}

type ValidatorStatusArgs struct {
	JobId string `json:"job_id"`
}

func (e1 Basic) ValidatorStatus(args ValidatorStatusArgs, results *ValidatorInstance) error {
	ctx, err := e1.Ctx()
	if err != nil {
		return err
	}
	ans, err := e1.jobClient.GetStatus(ctx, &pbjob.StatusRequest{Id: args.JobId})
	if err != nil {
		return err
	}
	s := ans.GetStatus()
	if s == nil {
		return errors.New("no status")
	}
	if s.Status == pbjob.Status_NEW {
		return errors.New("new")
	}
	if s.Status == pbjob.Status_STARTED {
		return errors.New("started")
	}
	if s.Status == pbjob.Status_FAILED {
		return errors.New("failed")
	}
	if s.Status != pbjob.Status_FINISHED {
		return errors.New("unknown status")
	}

	instance := new(pbsol.ValidatorInstance)
	err = json.Unmarshal(ans.GetData(), instance)
	if err != nil {
		return err
	}
	*results = instanceToResult(instance)
	return nil
}

func instanceToResult(instance *pbsol.ValidatorInstance) ValidatorInstance {
	return ValidatorInstance{
		Id:       instance.Id,
		Hostv4:   instance.Hostv4,
		Hostv6:   instance.Hostv6,
		Port:     instance.Port,
		User:     instance.User,
		Password: instance.Password,
	}
}

type ValidatorDestroyArgs struct {
	Id string `json:"id"`
}

type ValidatorDestroyResults bool

func (e1 Basic) ValidatorDestroy(args ValidatorDestroyArgs, results *ValidatorDestroyResults) error {

	ctx, err := e1.Ctx()
	if err != nil {
		return err
	}
	_, err = e1.solanaClient.DestroyValidator(ctx, &pbsol.ValidatorDeleteRequest{Id: args.Id})
	if err != nil {
		return err
	}
	*results = true
	return nil
}
