package basic

import (
	"errors"
	"fmt"
	"math"

	pbb "github.com/noncepad/client/proto/base"
	pbc "github.com/noncepad/client/proto/customer"
	pbsaas "github.com/noncepad/client/proto/saas"
)

func (e1 *external) service(args []string) error {
	if len(args) < 1 {
		return errors.New("bad argument for service")
	}
	subcmd := args[0]
	var err error
	switch subcmd {
	case SERVICE_CMD_MENU:
		err = e1.service_menu(args[1:])
	case SERVICE_CMD_PORTAL:
		err = e1.service_portal(args[1:])
	case SERVICE_CMD_SUB:
		err = e1.service_sub(args[1:])
	case SERVICE_CMD_UNSUB:
		err = e1.service_unsub(args[1:])
	default:
		e1.service_help()
		err = errors.New("bad command")
	}
	if err != nil {
		return err
	}
	return nil
}

func (e1 *external) init_service() {
	e1.serviceToString = map[pbsaas.Service]string{
		pbsaas.Service_VALIDATOR_RPC:      "validator_rpc",
		pbsaas.Service_VALIDATOR_INSTANCE: "validator_instance",
	}
	e1.serviceToEnum = make(map[string]pbsaas.Service)
	for k, v := range e1.serviceToString {
		e1.serviceToEnum[v] = k
	}
}

func (e1 *external) service_portal(args []string) error {

	client := pbc.NewSaasClient(e1.conn)

	ans, err := client.LinkToStripe(e1.Ctx(), &pbb.Empty{})
	if err != nil {
		return err
	}
	e1.print_str(fmt.Sprintf("Stripe URL:   %s", ans.StripeUrl))
	return nil
}

func (e1 *external) service_menu(args []string) error {

	client := pbc.NewSaasClient(e1.conn)

	list, err := client.ListProduct(e1.Ctx(), &pbc.SaasListProductRequest{})
	if err != nil {
		return err
	}

	for i := 0; i < len(list.Product); i++ {
		p := list.Product[i]
		if p.Description == nil {
			return errors.New("bad description")
		}
		if p.Name == nil {
			return errors.New("bad name")
		}
		e1.print_str(fmt.Sprintf("Name: %s\n\tDescription: %s", p.Name.En_US, p.Description.En_US))
		e1.print_str("\tTiers (Fixed/Floating)")
		for j := 0; j < len(p.Tier); j++ {
			fixed, err := e1.rate_to_string(p.Tier[j].Fixed)
			if err != nil {
				return err
			}
			float, err := e1.rate_to_string(p.Tier[j].Float)
			if err != nil {
				return err
			}
			e1.print_str(fmt.Sprintf("\t\t%s   /  %s", fixed, float))
		}
	}

	//e1.print_json(list.Product)

	return nil
}
func (e1 *external) rate_to_string(rate *pbb.Rate) (string, error) {
	var unit string
	switch rate.Unit {
	case pbb.Rate_DATA_PER_KB:
		unit = " per KB "
	case pbb.Rate_DATA_PER_MB:
		unit = " per MB "
	case pbb.Rate_TIME_PER_DAY:
		unit = " per Day "
	case pbb.Rate_TIME_PER_YEAR:
		unit = " per Year "
	}
	valueStr, err := e1.value_to_string(rate.Amount)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", valueStr, unit), nil
}

func (e1 *external) value_to_string(value *pbb.Value) (string, error) {

	short, err := e1.get_short_commodity(value.CommodityId)
	if err != nil {
		return "", err
	}

	v := float64(value.Amount)
	d := float64(math.Pow10(int(short.Decimal)))

	if len(short.Ticker) < 3 {
		return "", errors.New("bad ticker")
	}

	return fmt.Sprintf("%f %s_%s", v/d, short.Ticker[0], short.Ticker[1]), nil
}

func (e1 *external) get_short_commodity(id int64) (*pbb.ShortCommodity, error) {

	if e1.conn == nil {
		return nil, errors.New("no connection")
	}
	client := pbc.NewLedgerClient(e1.conn)
	ans, err := client.GetCommodity(e1.Ctx(), &pbc.CommodityRequest{
		CommodityId: id,
	})
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func (e1 *external) service_sub(args []string) error {
	if len(args) < 1 {
		return errors.New("service name not specified")
	}
	service, present := e1.serviceToEnum[args[0]]
	if !present {
		return errors.New("no such service")
	}

	client := pbc.NewSaasClient(e1.conn)

	resp, err := client.Subscribe(e1.Ctx(), &pbc.SaasSubscribeRequest{
		Service: service,
	})
	if err != nil {
		return err
	}

	e1.print_str(fmt.Sprintf("Stripe URL for billing: %s", resp.GetStripeUrl()))

	return errors.New("not implemented yet")
}

func (e1 *external) service_unsub(args []string) error {
	return errors.New("not implemented yet")
}

type ServiceSubCommand = string

const (
	SERVICE_CMD_MENU   ServiceSubCommand = "menu"
	SERVICE_CMD_PORTAL ServiceSubCommand = "portal"
	SERVICE_CMD_LS     ServiceSubCommand = "ls"
	SERVICE_CMD_SUB    ServiceSubCommand = "sub"
	SERVICE_CMD_UNSUB  ServiceSubCommand = "unsub"
)

func (e1 *external) service_help() {
	fmt.Print(`
Command: service

==============

portal - get a link to visit the Stripe web portal

ls - list services currently subscribed to

sub - subscribe to a new service

unsub - unsubscribe from a service
	
	`)
}
