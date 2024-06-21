package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDenom{}

func NewMsgCreateDenom(
	owner string,
	denom string,
	description string,
	ticker string,
	precision int32,
	url string,
	maxSupply int32,
	canChangeMaxSupply bool,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Owner:              owner,
		Denom:              denom,
		Description:        description,
		Ticker:             ticker,
		Precision:          precision,
		Url:                url,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	tickerLen := len(msg.Ticker)

	if tickerLen < 3 || tickerLen > 10 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ticker length must be between 3 and 10")
	}
	if msg.MaxSupply <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max supply must be positive")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDenom{}

func NewMsgUpdateDenom(
	owner string,
	denom string,
	description string,
	ticker string,
	precision int32,
	url string,
	maxSupply int32,

	canChangeMaxSupply bool,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Owner:              owner,
		Denom:              denom,
		Description:        description,
		Url:                url,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgUpdateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if msg.MaxSupply <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max supply must be positive")
	}
	return nil
}

