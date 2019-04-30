package user

import (
	"errors"
	"sync"
	"testing"
)

type UserData struct {
	name string
	balance float64
}

type MockUser struct {
	userMap map[int] *UserData
} 

var once sync.Once 
var mock *MockUser

func InitTest()  {
	once.Do(func() {
		getUserNameById = GetUserNameByIdFromMock
		getUserBanlanceById = GetUserBabanceByIdFromMock
		mock = &MockUser{
			userMap: make(map[int]*UserData),
		}
		mock.userMap[1] = &UserData{"xiao hong", 100}
		mock.userMap[2] = &UserData{"xiao ming", 50}
	})
} 

func GetUserNameByIdFromMock(id int)  (string, error)  {
	if user, exist := mock.userMap[id]; exist {
		return user.name, nil
	}
	return "", errors.New("not found")
}

func GetUserBabanceByIdFromMock(id int)  (float64, error)  {
	if user, exist := mock.userMap[id]; exist {
		return user.balance, nil
	}
	return 0, errors.New("not found")
}
 
func Test_GetUserInfo(t *testing.T)  {
	InitTest()
	name, balance, err := GetUserInfo(1)
	if err != nil{
		t.Error(err)
		return
	} 
	t.Logf("id: %v, name:%v, balance:%v", 1, name, balance)

	name, balance, err = GetUserInfo(2)
	if err != nil{
		t.Error(err)
		return
	}
	t.Logf("id: %v, name:%v, balance:%v", 2, name, balance)
}