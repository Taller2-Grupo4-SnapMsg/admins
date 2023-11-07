package structs

type AdminStruct struct {
	Email     string
	Password  string
	TimeStamp int64
}

type Credentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
