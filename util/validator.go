package util

import (
	"errors"

	pbsol "github.com/noncepad/client/proto/solana"
)

type ValidatorType = string

const (
	VALIDATOR_MAINNET ValidatorType = "mainnet"
	VALIDATOR_DEVNET  ValidatorType = "devnet"
	TPU_MAINNET       ValidatorType = "tpu"
)

func ValidatorTypeStringToProto(x ValidatorType) (pbsol.ValidatorType, error) {

	switch x {
	case VALIDATOR_MAINNET:
		return pbsol.ValidatorType_MAINNET, nil
	case VALIDATOR_DEVNET:
		return pbsol.ValidatorType_DEVNET, nil
	case TPU_MAINNET:
		return pbsol.ValidatorType_MAINNET_TPU, nil
	default:
		return pbsol.ValidatorType_MAINNET, errors.New("no matching type")
	}
}
