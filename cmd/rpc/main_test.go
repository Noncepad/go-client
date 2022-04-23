package main_test

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/rpc"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	solmateRpc "github.com/noncepad/client/rpc"
	basic "github.com/noncepad/client/rpc/basic"
	test "github.com/noncepad/client/rpc/test"
)

// test against api.test.noncepad.com
func TestRemoteRpc(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	tcpConn, err := net.Dial("tcp", "api.test.noncepad.com:443")
	if err != nil {
		t.Fatal(err)
	}
	conn := tls.Client(tcpConn, &tls.Config{
		ServerName: "api.test.noncepad.com",
	})
	client := rpc.NewClient(conn)
	//make connection to rpc server
	//client, err := rpc.DialHTTP("tcp", "api.test.noncepad.com")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}

	testTest(t, client)
	tokenC := make(chan string, 1)
	testSession(t, client, tokenC)
	token := <-tokenC

	testTxSubmit(t, client, token)
	//testPortal(t, client, token)
	testMenu(t, client, token)
	//idC := make(chan string, 1)
	//testTpuCreate(t, client, token, idC)
	//testValidatorLS(t, client, token)

	//t.Logf("id=%s", <-idC)
	//idList := []string{"1c3eebb9-6003-4bfa-a204-52a204bcf68f"}
	//for i := 0; i < len(idList); i++ {
	//	testValidatorDestroy(t, client, token, idList[i])
	//}
	//testValidatorDestroy(t, client, token, <-idC)
}

// test a locally run RPC server
func TestRpc(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}
	config, err := basic.ConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	listenUrl := ":34001"
	serverErrorC := make(chan error, 1)
	err = solmateRpc.Run(ctx, config, listenUrl, serverErrorC)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cancel()
		<-serverErrorC
	})

	time.Sleep(2 * time.Second)

	//make connection to rpc server
	client, err := rpc.DialHTTP("tcp", listenUrl)
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}

	testTest(t, client)
	tokenC := make(chan string, 1)
	testSession(t, client, tokenC)
	token := <-tokenC

	testTxSubmit(t, client, token)
	//testPortal(t, client, token)
	//testMenu(t, client, token)
	//idC := make(chan string, 1)
	//testTpuCreate(t, client, token, idC)
	//testValidatorLS(t, client, token)

	//t.Logf("id=%s", <-idC)
	//idList := []string{"1c3eebb9-6003-4bfa-a204-52a204bcf68f"}
	//for i := 0; i < len(idList); i++ {
	//	testValidatorDestroy(t, client, token, idList[i])
	//}
	//testValidatorDestroy(t, client, token, <-idC)
}

func testTxSubmit(t *testing.T, client *rpc.Client, token string) {
	var err error
	createArgs := &basic.SendBatchArgs{
		RequestWithJwt:  basic.RequestWithJwt{Token: token},
		CommitmentLevel: "confirmed",
		Tx:              []string{"VzAV/b3dTLGAt3lTCZNbeShqvrlzBuSIemtaTnyaspBc45p2rxiW6toICMT3vUMY1ytz3GRo/gtwAC9WtKxSJC7kHDbSR1qtQLroZbVTC3x4IHhab+SGTcOe+O8hZhnmS/DWqqrBdB5JyRfEfetuxAOeIMojtqWxNE342snhG/8vmklzZq2BmtPqJ/amOJX7C94tT1vMTCtJJeHsf4UVbazRZA4dQgQe1WTnKpuOj0/KSDQ6xmcsOfudDXptuYK5VY7gJhm8aogiTRaTT24sy3YEaUtzeE1keBoUYNRCMtcmuNwXa7Dm+BIOKsoXii/YS/9V9nh6hPgmB6i3/jHXHq4RtXfZbPyAN/ePZ9e2ZjTSk5UfYELQafVarrz5Gaappy+NrmTAmC6fprsVyJwHY8C+LlQBFl074f+ux/vaQE88mADKq0J0omlp7L1PGLdO2O+IOj/EqHqdHw2LmmRH7mHLTizzjCM5apCXqNm1Xiw/JKbfqwMFDzh+a1T6vkiIw+xCpCPAZQRkyjcUKybsbvtFjwa/bAAcjhQ0+wyMm1mijQid23Cy94oUMRArpdVnDHx2rMy3rJikXM2E5ToItkMusU1m6neElzPx/I4AqeCsZ3qlRNz/mZWHHk/NrKFmfzZD3O9xOq0idyK1Z7Xkk+iGKnGQxE9oqFshIRG8Gp+7LsdDuGDcUvsQtJDjbvSkRRsiUNXr8xuO7/5ZwbIQrqumnAfzR2tBCJfVJHZH0X1+m3nu/+L2ap+IpfaK+G5t+6ISurKSTTpxv2U7qm8t02azMJaXNSt5wCAcb79He2no436MFkar9LiakLSd4IfYIdT808xbXRHUSFViGxQvCTIlH+yPcE+FEg0ftuv06N0OFP7wq+UYlSPoPc2DGOGrC6ikEe3l50JZT3NrfJF7jyWh78pyGRJh6rL7nS09kpbDlBgsVwTKcznmRLajivYYOr+U+vMOozvYB+yE4MdpiMNDxjm/O7DZbR3ojvSPWp6LfeueXKuj1wIyV6gvwnVRZGGwNKxcD+UpF5zf0ijS0lRh6+gNXAR1EwGs+53R3fwTN8Pi+yFkZEDp6eyVFSk5BdKJYTPuMHBwZ7dsz/R3Bc22UczQY7AIdjR/NWdujY+uQsF1JOekd/iD+I56g7wAXNyZXWJoY3cz2KTU5mWjaDMjOk09xvvweV+mkvITN1T7bBZIUnV2Tz4flxYcFYCCG2J9jiFo/Su7y6J5GpnuQ7xKLUR47yR21ITS4YHzeKw7WBaCbEcfeC74ULOioiFfsszfLFN9UA81B7Pt8feLovvs/Wfvp9SGsqchKiN3XDkWzeY1h+FYZsaV3t6vLozmdxHOkUM+sn4neoT+Q/P4nw=="},
	}
	createResults := new(basic.SendBatchResponse)
	err = client.Call("Basic.SendTx", createArgs, createResults)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("results=%+v", createResults)
}

func testTpuCreate(t *testing.T, client *rpc.Client, token string, idC chan<- string) {
	var err error

	createArgs := &basic.ValidatorCreateArgs{
		RequestWithJwt: basic.RequestWithJwt{Token: token},
		Type:           "tpu",
	}
	createResults := new(basic.Job)
	err = client.Call("Basic.ValidatorCreate", createArgs, createResults)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("results=%+v", createResults)
	jobId := createResults.Id

	var instance *basic.ValidatorInstance
out:
	for i := 0; i < 10; i++ {
		time.Sleep(time.Minute)
		args := &basic.ValidatorStatusArgs{
			RequestWithJwt: basic.RequestWithJwt{Token: token},
			JobId:          jobId,
		}
		results := new(basic.ValidatorInstance)
		err = client.Call("Basic.ValidatorStatus", args, results)
		if err == nil {
			instance = results
			break out
		}
	}
	if instance == nil {
		t.Fatal("failed to create instance")
	}
	log.Printf("instance=%+v", instance)
	idC <- instance.Id

}

func testValidatorDestroy(t *testing.T, client *rpc.Client, token string, instanceId string) {
	var err error
	args := &basic.ValidatorDestroyArgs{
		RequestWithJwt: basic.RequestWithJwt{Token: token},
		Id:             instanceId,
	}
	results := new(basic.ValidatorDestroyResults)
	log.Printf("destroy=%+v", args)
	err = client.Call("Basic.ValidatorDestroy", args, results)
	if err != nil {
		t.Fatal(err)
	}
}

func testValidatorLS(t *testing.T, client *rpc.Client, token string) {
	var err error

	args := &basic.ValidatorLSArgs{
		RequestWithJwt: basic.RequestWithJwt{Token: token},
	}
	results := new(basic.ValidatorLSResults)
	err = client.Call("Basic.ValidatorLS", args, results)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("results=%+v", results)
}

func testMenu(t *testing.T, client *rpc.Client, token string) {
	var err error

	args := &basic.ServiceMenuArgs{
		RequestWithJwt: basic.RequestWithJwt{Token: token},
	}
	results := new(basic.ServiceMenuResults)
	err = client.Call("Basic.ServiceMenu", args, results)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("results=%+v", results)
}

func testPortal(t *testing.T, client *rpc.Client, token string) {
	var err error

	args := &basic.ServicePortalArgs{
		RequestWithJwt: basic.RequestWithJwt{Token: token},
	}
	results := new(basic.ServicePortalResults)
	err = client.Call("Basic.ServicePortal", args, results)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("results=%+v", results)
}

func testSession(t *testing.T, client *rpc.Client, tokenC chan<- string) {
	var err error
	apiKey, present := os.LookupEnv("API_KEY")
	if !present {
		t.Fatal("no api key")
	}

	args := &basic.RefreshArgs{ApiKey: apiKey}
	results := new(basic.RefreshResults)
	err = client.Call("Basic.Refresh", args, results)
	if err != nil {
		t.Fatal(err)
	}
	tokenC <- results.Token
}

func testTest(t *testing.T, client *rpc.Client) {
	var err error
	//make arguments object
	args := &test.Args{
		A: 2,
		B: 3,
	}
	//this will store returned result
	var result test.Result
	//call remote procedure with args
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		t.Fatal(err)
	}
	//we got our result in result
	if result != 6 {
		t.Fatal("bad multiplication")
	}

	err = client.Call("Arith.Substract", args, &result)
	if err != nil {
		t.Fatal(err)
	}

	//we got our result in result
	if result != -1 {
		t.Fatal("bad multiplication")
	}
}
