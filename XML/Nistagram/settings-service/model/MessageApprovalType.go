package model

type MessageApprovalType int

const (
	PUBLIC MessageApprovalType = iota
	FRIENDS_ONLY
)
