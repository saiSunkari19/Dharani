package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/tendermint/tendermint/libs/bech32"
)

type claimReq struct {
	Address string `json:"address"`
}

func faucetHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var claim claimReq
		decoder := json.NewDecoder(r.Body)
		decoderErr := decoder.Decode(&claim)
		if decoderErr != nil {
			log.Println(decoderErr)
		}
		readableAddress, decodedAddress, decodeErr := bech32.DecodeAndConvert(claim.Address)
		if decodeErr != nil {
			log.Println(decodeErr)
		}
		encodedAddress, encodeErr := bech32.ConvertAndEncode(readableAddress, decodedAddress)
		if encodeErr != nil {
			log.Println(encodeErr)
		}
		cmd := exec.Command("dharanicli", "tx", "send", "me", encodedAddress, "100cent", "-y")
		_, err := cmd.Output()
		if err != nil {
			log.Println(fmt.Sprintf("%s", err))
		}
		rest.PostProcessResponse(w, cliCtx, encodedAddress)
	}
}
