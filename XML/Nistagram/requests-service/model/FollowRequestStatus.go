package model

type FollowRequestStatus int

const(
	PENDING FollowRequestStatus= iota
	ACCEPTED
	REJECT
)
