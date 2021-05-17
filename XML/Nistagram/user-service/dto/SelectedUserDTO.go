package dto

type SelectedUserDTO struct {
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Website string `json:"website"`
	Biography string `json:"biography"`
	ProfileVisibility string `json:"profileVisibility"` //public or private profile
	FollowingCheck bool `json:"followingCheck"` //returns true if logged in user already follows selected user
	//posts
}
