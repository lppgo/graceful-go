package user

import (
	"errors"
)
 
var getUserNameById = GetUserNameById
var getUserBanlanceById = GetUserBalanceById

func GetUserNameById(id int) (string, error)  {
	//todo： query database
	return "", errors.New("not found")
}

func GetUserBalanceById(id int) (float64, error)  {
	//todo： query database
	return 0, errors.New("not found")
}

func GetUserInfo(id int) (string, float64, error)  {
	name , err := getUserNameById(id)
	if err != nil {
		return name, 0, err
	}
	balance , err := getUserBanlanceById(id)
	if err != nil {
		return name, 0, err
	}
	return name, balance, nil
}