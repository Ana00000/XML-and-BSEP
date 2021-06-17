package model

type RequestStatus int

const (
	PENDING RequestStatus = iota
	ACCEPTED
	REJECT
)
