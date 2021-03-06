package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	types2 "github.com/dharani/types"
	"github.com/dharani/x/dharani/types"
)

func CommandSellProperty(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell",
		Short: "to sell property value",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI(nil).WithTxEncoder(client.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			_propertyID := viper.GetString(flagPropertyID)
			area := viper.GetUint64(flagArea)
			coin := viper.GetString(flagPrice)
			propertyID, err := types2.NewPropertyIDFromString(_propertyID)
			if err != nil {
				return err
			}
			cost, err := sdk.ParseCoin(coin)
			if err != nil {
				return err
			}
			
			msg := types.NewMsgSellProperty(ctx.FromAddress, propertyID, area, cost)
			
			return client.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	
	cmd.Flags().String(flagPropertyID, "", "property id")
	cmd.Flags().Int64(flagArea, 0, "area in square meters")
	cmd.Flags().String(flagPrice, "", "price to sell property")
	
	_ = cmd.MarkFlagRequired(flagPropertyID)
	_ = cmd.MarkFlagRequired(flagArea)
	_ = cmd.MarkFlagRequired(flagPrice)
	
	return cmd
}
