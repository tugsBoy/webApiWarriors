package services

import (
	"errors"
	"fmt"

	"github.com/tugsBoy/webApiWarriors/types"
)

// type ProcessLoginer interface {
// 	ProcessLogin(user types.User) (types.PlayerSave, error)
// }

//type processLoginImpl struct{}

func createDBStub() []types.PlayerSave {
	stubPlayerTable := []types.PlayerSave{
		{
			LoginId: "Justin",
			Warriors: []types.Warrior{
				{Name: "Aragorn", Sex: "Male", Age: 87, Strength: 20, Dex: "No!", Intelligence: 15, Faith: 10, Arcane: 5, Luck: 12},
				{Name: "Legolas", Sex: "Male", Age: 2931, Strength: 18, Dex: "No!", Intelligence: 16, Faith: 8, Arcane: 6, Luck: 14},
			},
		},
		{
			LoginId: "Gabe",
			Warriors: []types.Warrior{
				{Name: "Gimli", Sex: "Male", Age: 139, Strength: 19, Dex: "No!", Intelligence: 12, Faith: 7, Arcane: 5, Luck: 10},
				{Name: "Boromir", Sex: "Male", Age: 41, Strength: 18, Dex: "No!", Intelligence: 14, Faith: 9, Arcane: 5, Luck: 12},
			},
		},
		{
			LoginId: "Chris",
			Warriors: []types.Warrior{
				{Name: "Frodo", Sex: "Male", Age: 50, Strength: 10, Dex: "No!", Intelligence: 14, Faith: 18, Arcane: 5, Luck: 20},
				{Name: "Sam", Sex: "Male", Age: 38, Strength: 12, Dex: "No!", Intelligence: 13, Faith: 19, Arcane: 5, Luck: 18},
			},
		},
	}
	return stubPlayerTable
}

func ProcessLogin(user types.User) (types.PlayerSave, error) {
	stubPlayerTable := createDBStub()
	if user.LoginID == "" {
		return types.PlayerSave{}, types.ErrBlankUserID
	}
	for _, e := range stubPlayerTable {
		if e.LoginId == user.LoginID {
			return e, nil
		}
	}
	return types.PlayerSave{}, types.ErrUserNotFound
}

func CreateNewUser(user types.User) error {
	if user.LoginID == "" {
		return errors.New("JSON parse error:\n Cannot parse incoming JSON to backend data structure")
	}
	newPlaerSave := types.PlayerSave{
		LoginId:  user.LoginID,
		Warriors: []types.Warrior{},
	}
	stubPlayerTable := createDBStub()
	stubPlayerTable = append(stubPlayerTable, newPlaerSave)
	fmt.Println(stubPlayerTable)
	for _, player := range stubPlayerTable {
		if player.LoginId == user.LoginID {
			return nil
		}
	}
	return errors.New("Faild to create user " + user.LoginID + ".")
}
