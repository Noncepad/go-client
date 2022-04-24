package basic

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"time"

	pbsol "github.com/noncepad/client/proto/solana"
)

type SendBatchArgs struct {
	Tx              []string `json:"tx"`
	CommitmentLevel string   `json:"commitment"`
}

type SendBatchResponse struct {
	IsDone    bool     `json:"done"`
	Signature []string `json:"signature"`
}

func sendBatchRespToProto(resp *pbsol.SendBatchResponse) SendBatchResponse {
	if resp == nil {
		return SendBatchResponse{IsDone: false, Signature: []string{}}
	}
	if resp.Signature == nil {
		resp.Signature = make([][]byte, 0)
	}
	list := make([]string, len(resp.Signature))
	for i := 0; i < len(resp.Signature); i++ {
		list[i] = base64.StdEncoding.EncodeToString(resp.Signature[i])
	}
	return SendBatchResponse{Signature: list, IsDone: true}
}

type SendTxJob struct {
	Id int `json:"id"`
}

func (e1 Basic) SendTx(args SendBatchArgs, results *SendTxJob) error {
	ctx, err := e1.Ctx()
	if err != nil {
		return err
	}
	payload := &pbsol.SendBatchRequest{}

	switch args.CommitmentLevel {
	case "finalized":
		payload.ConfirmationLevel = pbsol.ConfirmationLevel_FINALIZED
	case "processed":
		payload.ConfirmationLevel = pbsol.ConfirmationLevel_PROCESSED
	case "confirmed":
		payload.ConfirmationLevel = pbsol.ConfirmationLevel_CONFIRMED
	default:
		return errors.New("unknown confirmation level")
	}
	if len(args.Tx) < 1 {
		return errors.New("no transactions")
	}

	payload.Tx = make([][]byte, len(args.Tx))
	for i := 0; i < len(args.Tx); i++ {
		payload.Tx[i], err = base64.StdEncoding.DecodeString(args.Tx[i])
		if err != nil {
			return err
		}
	}

	id := e1.get_job_id()

	*results = SendTxJob{Id: id}

	go e1.loopFinishSendTxJob(ctx, id, payload)

	return nil
}

func (e1 Basic) loopFinishSendTxJob(ctx context.Context, id int, payload *pbsol.SendBatchRequest) {
	resp, err := e1.solanaClient.SendTx(ctx, payload)
	log.Printf("recevied response for job id=%d with resp=%+v and err=%s", id, resp, err)
	e1.job_set(id, resp, err)
}

type SendTxJobStatusArgs struct {
	Id int `json:"id"`
}

func (e1 Basic) GetSendTxResult(args SendTxJobStatusArgs, results *SendBatchResponse) error {
	ans := e1.job_get(args.Id)
	if ans == nil {
		return errors.New("job does not exist")
	} else if ans.IsDone {
		*results = sendBatchRespToProto(ans.Ans)
	} else {
		*results = SendBatchResponse{IsDone: false, Signature: []string{}}
	}

	return nil
}

type TxJob struct {
	IsDone bool
	Ans    *pbsol.SendBatchResponse
	Err    error
}

func (e1 Basic) get_job_id() int {
	idC := make(chan int, 1)
	e1.internalC <- func(in *internal) {
		in.jobId++
		id2 := in.jobId
		idC <- id2
		in.jobMap[id2] = &TxJob{IsDone: false}
	}
	return <-idC
}

func (e1 Basic) job_set(id int, d *pbsol.SendBatchResponse, err error) {
	e1.internalC <- func(in *internal) {
		x, present := in.jobMap[id]
		if present {
			x.IsDone = true
			x.Err = err
			x.Ans = d
			go e1.job_delete(id, 5*time.Minute)
		} else {
			log.Print("response from server with no job id")
		}
	}
}

func (e1 Basic) job_delete(id int, expire time.Duration) {
	select {
	case <-e1.ctx.Done():
	case <-time.After(expire):
		e1.internalC <- func(in *internal) {
			delete(in.jobMap, id)
		}
	}
}

func (e1 Basic) job_get(id int) *TxJob {
	dC := make(chan *TxJob, 1)
	e1.internalC <- func(in *internal) {
		x, present := in.jobMap[id]
		if present {
			dC <- x
		} else {
			dC <- nil
		}
	}
	return <-dC
}
