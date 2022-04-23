package basic

import (
	"errors"
	"fmt"

	pbb "github.com/noncepad/client/proto/base"
	pbc "github.com/noncepad/client/proto/customer"
	pbsaas "github.com/noncepad/client/proto/saas"
)

func init_service() (map[pbsaas.Service]string, map[string]pbsaas.Service) {
	serviceToString := map[pbsaas.Service]string{
		pbsaas.Service_VALIDATOR_RPC:      "validator_rpc",
		pbsaas.Service_VALIDATOR_INSTANCE: "validator_instance",
	}
	serviceToEnum := make(map[string]pbsaas.Service)
	for k, v := range serviceToString {
		serviceToEnum[v] = k
	}
	return serviceToString, serviceToEnum
}

type ServicePortalArgs struct {
	RequestWithJwt
}

type ServicePortalResults struct {
	PortalUrl string
}

func (e1 Basic) ServicePortal(args ServicePortalArgs, results *ServicePortalResults) error {

	ans, err := e1.saasClient.LinkToStripe(e1.Ctx(args.Token), &pbb.Empty{})
	if err != nil {
		return err
	}
	if ans == nil {
		return errors.New("blank response")
	}

	*results = ServicePortalResults{PortalUrl: ans.StripeUrl}

	return nil
}

type ServiceMenuArgs struct {
	RequestWithJwt
}

type Value struct {
	Amount  int64
	Decimal uint8
	Ticker  string
}

type Rate struct {
	Value Value
	Unit  string
}

type Tier struct {
	Fixed    Rate
	Floating Rate
	UpTo     int64
}

type Product struct {
	Name        string
	Description string
	Tier        []Tier
}

type ServiceMenuResults struct {
	Product []Product
}

func (e1 Basic) ServiceMenu(args ServiceMenuArgs, results *ServiceMenuResults) error {

	list, err := e1.saasClient.ListProduct(e1.Ctx(args.Token), &pbc.SaasListProductRequest{})
	if err != nil {
		return err
	}

	x := make([]Product, len(list.Product))

	for i := 0; i < len(list.Product); i++ {
		newP := new(Product)

		p := list.Product[i]
		if p.Description == nil {
			return errors.New("bad description")
		}
		if p.Name == nil {
			return errors.New("bad name")
		}

		newP.Tier = make([]Tier, len(p.Tier))
		newP.Name = p.Name.En_US
		newP.Description = p.Description.En_US

		for j := 0; j < len(p.Tier); j++ {

			fixed, err := e1.rate_to_string(args.Token, p.Tier[j].Fixed)
			if err != nil {
				return err
			}
			floating, err := e1.rate_to_string(args.Token, p.Tier[j].Float)
			if err != nil {
				return err
			}
			upto := p.Tier[j].UpTo

			newP.Tier[j] = Tier{Fixed: *fixed, Floating: *floating, UpTo: upto}
		}
		x[i] = *newP
	}

	*results = ServiceMenuResults{Product: x}

	return nil
}

func (e1 Basic) rate_to_string(token string, rate *pbb.Rate) (*Rate, error) {
	var unit string
	switch rate.Unit {
	case pbb.Rate_DATA_PER_KB:
		unit = "per KB"
	case pbb.Rate_DATA_PER_MB:
		unit = "per MB"
	case pbb.Rate_TIME_PER_DAY:
		unit = "per Day"
	case pbb.Rate_TIME_PER_YEAR:
		unit = "per Year"
	}
	var err error
	r := new(Rate)
	r.Unit = unit
	v, err := e1.value_to_string(token, rate.Amount)
	r.Value = *v
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Divide(a int64, b int64) (int64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	c := b
	if b < 0 {
		c = -1 * b
	}
	r := a % c
	return (a - r) / b, nil
}

// return amount, decimal, string, error
func (e1 Basic) value_to_string(token string, value *pbb.Value) (*Value, error) {

	short, err := e1.get_short_commodity(token, value.CommodityId)
	if err != nil {
		return nil, err
	}

	if short.Decimal < 0 || 255 < short.Decimal {
		return nil, errors.New("decimal out of bounds")
	}

	return &Value{
		Amount: value.Amount, Decimal: uint8(short.Decimal), Ticker: fmt.Sprintf("%s_%s", short.Ticker[0], short.Ticker[1]),
	}, nil
}

func (e1 Basic) get_short_commodity(token string, id int64) (*pbb.ShortCommodity, error) {

	if e1.conn == nil {
		return nil, errors.New("no connection")
	}
	client := pbc.NewLedgerClient(e1.conn)
	ans, err := client.GetCommodity(e1.Ctx(token), &pbc.CommodityRequest{
		CommodityId: id,
	})
	if err != nil {
		return nil, err
	}

	return ans, nil
}

type ServiceSubArgs struct {
	RequestWithJwt
	Service string
}

func (e1 Basic) ServiceSub(args ServiceSubArgs, results *ServicePortalResults) error {

	service, present := e1.serviceToEnum[args.Service]
	if !present {
		return errors.New("no such service")
	}

	resp, err := e1.saasClient.Subscribe(e1.Ctx(args.Token), &pbc.SaasSubscribeRequest{
		Service: service,
	})
	if err != nil {
		return err
	}

	*results = ServicePortalResults{PortalUrl: resp.GetStripeUrl()}

	return errors.New("not implemented yet")
}
