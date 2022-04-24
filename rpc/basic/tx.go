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
	Signature []string `json:"signature"`
}

func sendBatchRespToProto(resp *pbsol.SendBatchResponse) SendBatchResponse {
	list := make([]string, len(resp.Signature))
	for i := 0; i < len(resp.Signature); i++ {
		list[i] = base64.StdEncoding.EncodeToString(resp.Signature[i])
	}
	return SendBatchResponse{Signature: list}
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
	log.Printf("recevied response for job id=%d", id)
	e1.job_set(id, resp, err)
}

type SendTxJobStatusArgs struct {
	Id int `json:"id"`
}

func (e1 Basic) GetSendTxResult(args SendTxJobStatusArgs, results *SendBatchResponse) error {
	ans := e1.job_get(args.Id)
	if ans == nil {
		return errors.New("job does not exist")
	}
	*results = sendBatchRespToProto(ans.Ans)
	return nil
}

type TxJob struct {
	Ans *pbsol.SendBatchResponse
	Err error
}

func (e1 Basic) get_job_id() int {
	idC := make(chan int, 1)
	e1.internalC <- func(in *internal) {
		in.jobId++
		idC <- in.jobId
	}
	return <-idC
}

func (e1 Basic) job_set(id int, d *pbsol.SendBatchResponse, err error) int {
	idC := make(chan int, 1)
	e1.internalC <- func(in *internal) {
		in.jobId++
		id2 := in.jobId
		idC <- id2
		in.jobMap[id2] = &TxJob{Err: err, Ans: d}
	}
	go e1.job_delete(id, 5*time.Minute)
	return <-idC
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
