package client

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/evdatsion/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterEvidences(cdc)
}
