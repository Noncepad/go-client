package basic

import (
	"encoding/base64"
	"errors"

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

func (e1 Basic) SendTx(args SendBatchArgs, results *SendBatchResponse) error {
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
	resp, err := e1.solanaClient.SendTx(ctx, payload)
	if err != nil {
		return err
	}
	*results = sendBatchRespToProto(resp)
	return nil
}
