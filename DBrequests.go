package main

func DBloginUser(userName string, password string) *User {
	user := new(User)
	err := database.Model(user).Where("name = ? AND password = ?", userName, password).Select()
	if err != nil {
		return nil
	}
	return user
}

func DBfindUser(userName string) *User {
	user := new(User)
	err := database.Model(user).Where("name = ?", userName).Select()
	if err != nil {
		return nil
	}
	return user
}

func DBregisterUser(userName string, password string) bool {
	_, err := database.Model(&User{
		Id:       0,
		Name:     userName,
		Password: password,
	}).Insert()
	if err != nil {
		panic(err)
		return false
	}
	return true
}
