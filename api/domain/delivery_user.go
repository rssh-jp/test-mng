package domain

type RecvLogin struct {
	ID       string `json:"id" form:"id" query:"id"`
	Password string `json:"password" form:"password" query:"password"`
}

type SendLogin struct {
	Message string `json:"message"`
	Token   Token  `json:"token"`
}

type RecvUsersFetch struct {
	Token string `json:"token" form:"token" query:"token"`
}

type SendUsersFetch struct {
	Message string `json:"message"`
	Users   []User `json:"users"`
}

type RecvUsersUpdate struct {
	Token string `json:"token" form:"token" query:"token"`
	User  User   `json:"user" form:"user" query:"user"`
}

type SendUsersUpdate struct {
	Message string `json:"message"`
}

type RecvUsersGetOwn struct {
	Token string `json:"token" form:"token" query:"token"`
}

type SendUsersGetOwn struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}
