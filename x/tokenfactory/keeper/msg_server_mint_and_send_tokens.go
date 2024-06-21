package keeper

import (
	"context"

	"tokenfactory/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintAndSendTokens(goCtx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the denom
	denom, found := k.GetDenom(ctx, msg.Denom)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
	}

	// Verifying the existence and ownership of the denom.
	if denom.Owner != msg.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "unauthorized")
	}

	// Ensuring minting does not exceed the maximum supply.
	if denom.MaxSupply < denom.Supply+msg.Amount {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minting exceeds max supply")
	}
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}

	var mintCoins sdk.Coins

	// 增加新币
	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.Amount))))

	// mint the coins
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return nil, err
	}

	// send the coins
	if err := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins); err != nil {
		return nil, err
	}

	var denomObject = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        denom.Description,
		MaxSupply:          denom.MaxSupply,
		Supply:             msg.Amount,
		Precision:          denom.Precision,
		Ticker:             denom.Ticker,
		Url:                denom.Url,
		CanChangeMaxSupply: denom.CanChangeMaxSupply,
	}
	k.SetDenom(ctx, denomObject)
	return &types.MsgMintAndSendTokensResponse{}, nil
}
