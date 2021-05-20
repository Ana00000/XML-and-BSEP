package model

type LikedStatus int

const(
	LIKED LikedStatus= iota
	DISLIKED
	NEUTRAL
)