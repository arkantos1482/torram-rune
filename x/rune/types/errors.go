package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/rune module sentinel errors
var (
	ErrInvalidSigner     = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample            = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrRuneNotStaked     = sdkerrors.Register(ModuleName, 1, "rune is not staked")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 2, "unauthorized")
	ErrInvalidRuneStatus = sdkerrors.Register(ModuleName, 3, "invalid rune status for the operation")
)
