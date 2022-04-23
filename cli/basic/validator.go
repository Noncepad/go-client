package basic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/google/uuid"
	pbb "github.com/noncepad/client/proto/base"
	pbc "github.com/noncepad/client/proto/customer"
	pbjob "github.com/noncepad/client/proto/job"
	pbsol "github.com/noncepad/client/proto/solana"
	"github.com/noncepad/client/util"
)

func (e1 *external) validator(args []string) error {
	if len(args) < 1 {
		return errors.New("bad argument for service")
	}
	subcmd := args[0]
	var err error
	switch subcmd {
	case VALIDATOR_CMD_LS:
		err = e1.validator_ls(args[1:])
	case VALIDATOR_CMD_CREATE:
		err = e1.validator_create(args[1:])
	case VALIDATOR_CMD_DESTORY:
		err = e1.validator_destroy(args[1:])
	case VALIDATOR_CMD_STATUS:
		err = e1.validator_status(args[1:])
	default:
		e1.validator_help()
		err = errors.New("bad command")
	}
	if err != nil {
		return err
	}
	return nil
}

func (e1 *external) validator_ls(args []string) error {

	client := pbc.NewSolanaClient(e1.conn)

	ans, err := client.ListValidator(e1.Ctx(), &pbb.Empty{})
	if err != nil {
		return err
	}
	list := ans.GetInstance()

	e1.print_str(fmt.Sprintf("Total running instances:   %d", len(list)))
	for i := 0; i < len(list); i++ {
		e1.print_json(list[i])
	}
	return nil
}

func (e1 *external) validator_create(args []string) error {
	if len(args) < 1 {
		return errors.New("no validator type specified")
	}
	client := pbc.NewSolanaClient(e1.conn)

	valType, err := util.ValidatorTypeStringToProto(args[0])
	if err != nil {
		return err
	}

	ans, err := client.CreateValidator(e1.Ctx(), &pbsol.ValidatorCreateRequest{
		Type: valType,
	})
	if err != nil {
		return err
	}
	e1.print_str("One instance created")

	jobClient := pbjob.NewOwnerClient(e1.conn)
	stream, err := jobClient.GetStatusStream(e1.Ctx(), &pbjob.StatusRequest{Id: ans.Id})
	if err != nil {
		return err
	}
	return e1.validator_status_stream(stream)
}

func (e1 external) validator_status_stream(stream pbjob.Owner_GetStatusStreamClient) error {
	var ans *pbjob.Job
	var err error
out:
	for {
		ans, err = stream.Recv()
		if err == io.EOF {
			err = nil
			break out
		}
		if err != nil {
			break out
		}
		if ans.Status == nil {
			err = errors.New("blank status")
			break out
		}
		if ans.Status.Status == pbjob.Status_FINISHED {
			instance := new(pbsol.ValidatorInstance)
			err = json.Unmarshal(ans.Data, instance)
			if err != nil {
				break out
			}
			e1.print_instance(instance)
			err = nil
			break out
		} else if ans.Status.Status == pbjob.Status_FAILED {
			err = errors.New("job failed")
			break out
		} else {
			switch ans.Status.Status {
			case pbjob.Status_STARTED:
				e1.print_str("status=STARTED")
			default:
			}
		}
	}

	if err != nil {
		return err
	}
	return nil
}

func (e1 *external) print_instance(instance *pbsol.ValidatorInstance) {
	e1.print_str(fmt.Sprintf(`
Instance Information

==============

Id  %s

Connection: 
   IPV4  %s
   IPV6  %s
   Port  %d

Http Auth Basic
   User       %s
   Password   %s
	
	`, instance.Id, instance.Hostv4, instance.Hostv6, instance.Port, instance.User, instance.Password))
}

func (e1 *external) validator_destroy(args []string) error {
	if len(args) < 1 {
		return errors.New("instance id not specified")
	}
	instanceId, err := uuid.Parse(args[0])
	if err != nil {
		return err
	}
	client := pbc.NewSolanaClient(e1.conn)
	_, err = client.DestroyValidator(e1.Ctx(), &pbsol.ValidatorDeleteRequest{Id: instanceId.String()})
	if err != nil {
		return err
	}
	return nil
}

func (e1 *external) validator_status(args []string) error {
	if len(args) < 1 {
		return errors.New("job id not specified")
	}
	jobId, err := uuid.Parse(args[0])
	if err != nil {
		return err
	}
	client := pbjob.NewOwnerClient(e1.conn)
	stream, err := client.GetStatusStream(e1.Ctx(), &pbjob.StatusRequest{Id: jobId.String()})
	if err != nil {
		return err
	}
	return e1.validator_status_stream(stream)
}

type ValidatorSubCommand = string

const (
	VALIDATOR_CMD_LS      ServiceSubCommand = "ls"
	VALIDATOR_CMD_CREATE  ServiceSubCommand = "create"
	VALIDATOR_CMD_DESTORY ServiceSubCommand = "destroy"
	VALIDATOR_CMD_STATUS  ServiceSubCommand = "status"
)

func (e1 *external) validator_help() {
	fmt.Print(`
Command: service

==============

ls - list currently running instance

create - create a validator instance

destroy - destroy a validator instance

status - get the status of a create job
	
	`)
}
