package main

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Message struct {
	//UserFromName string `json:"userfromname"`
	UserNameTo string `json:"usernameto"`
	Message    string `json:"message"`
}
