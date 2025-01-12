package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnstakeRune{}

func NewMsgUnstakeRune(creator string, runeId string, owner string, amount uint64) *MsgUnstakeRune {
	return &MsgUnstakeRune{
		Creator: creator,
		RuneId:  runeId,
		Owner:   owner,
		Amount:  amount,
	}
}

func (msg *MsgUnstakeRune) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
