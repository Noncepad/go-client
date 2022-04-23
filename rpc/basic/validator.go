package basic

import (
	"encoding/json"
	"errors"
	"log"

	pbb "github.com/noncepad/client/proto/base"
	pbjob "github.com/noncepad/client/proto/job"
	pbsol "github.com/noncepad/client/proto/solana"
	"github.com/noncepad/client/util"
)

type ValidatorLSArgs struct {
	RequestWithJwt
}

type ValidatorInstance struct {
	Id       string
	Hostv4   string
	Hostv6   string
	Port     uint32
	User     string
	Password string
}
type ValidatorLSResults struct {
	Instance []ValidatorInstance
}

func (e1 Basic) ValidatorLS(args ValidatorLSArgs, results *ValidatorLSResults) error {

	ans, err := e1.solanaClient.ListValidator(e1.Ctx(args.Token), &pbb.Empty{})
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
	RequestWithJwt
	Type util.ValidatorType
}

func (e1 Basic) ValidatorCreate(args ValidatorCreateArgs, results *Job) error {
	valType, err := util.ValidatorTypeStringToProto(args.Type)
	if err != nil {
		return err
	}
	ans, err := e1.solanaClient.CreateValidator(e1.Ctx(args.Token), &pbsol.ValidatorCreateRequest{
		Type: valType,
	})
	if err != nil {
		return err
	}

	*results = Job{Id: ans.Id, Status: "new", Data: "", Type: ""}
	return nil
}

type ValidatorStatusArgs struct {
	RequestWithJwt
	JobId string
}

func (e1 Basic) ValidatorStatus(args ValidatorStatusArgs, results *ValidatorInstance) error {

	ans, err := e1.jobClient.GetStatus(e1.Ctx(args.Token), &pbjob.StatusRequest{Id: args.JobId})
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
	RequestWithJwt
	Id string
}

type ValidatorDestroyResults bool

func (e1 Basic) ValidatorDestroy(args ValidatorDestroyArgs, results *ValidatorDestroyResults) error {
	log.Printf("destroying id=%s", args.Id)
	_, err := e1.solanaClient.DestroyValidator(e1.Ctx(args.Token), &pbsol.ValidatorDeleteRequest{Id: args.Id})
	if err != nil {
		return err
	}
	*results = true
	return nil
}
