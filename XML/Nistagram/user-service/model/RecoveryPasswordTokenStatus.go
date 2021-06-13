package model

type RecoveryPasswordTokenStatus int

const(
	VALID RecoveryPasswordTokenStatus= iota
	VERIFIED
	INVALID
)