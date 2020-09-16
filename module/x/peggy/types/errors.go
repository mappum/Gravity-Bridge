package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInternal redacts the error message
	ErrInternal  = sdkerrors.Register(ModuleName, 1, "internal")
	ErrDuplicate = sdkerrors.Register(ModuleName, 2, "duplicate")
	ErrInvalid   = sdkerrors.Register(ModuleName, 3, "invalid")
	ErrTimeout   = sdkerrors.Register(ModuleName, 4, "timeout")
	ErrUnknown   = sdkerrors.Register(ModuleName, 5, "unknown")
	ErrEmpty     = sdkerrors.Register(ModuleName, 6, "empty")
)