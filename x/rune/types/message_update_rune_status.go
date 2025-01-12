package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateRuneStatus{}

func NewMsgUpdateRuneStatus(authority string, runeId string, status string) *MsgUpdateRuneStatus {
	return &MsgUpdateRuneStatus{
		Authority: authority,
		RuneId:    runeId,
		Status:    status,
	}
}

func (msg *MsgUpdateRuneStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	return nil
}
