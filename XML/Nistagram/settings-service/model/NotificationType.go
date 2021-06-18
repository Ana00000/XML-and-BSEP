package model

type NotificationType int

const (
	ALL_NOTIFICATIONS NotificationType = iota
	FRIENDS_NOTIFICATIONS
	NONE
)