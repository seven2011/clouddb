package vo

//  user login

type UserLoginParams struct {
	Phone string 	`json:"phone"`
}
// User del
type UserDelParams struct {
	Id string 	`json:"id"`
	Token string `json:"token"`
}


//  user list
type UserListParams struct {
	Id string 	`json:"id"`
	Token string `json:"token"`
}
