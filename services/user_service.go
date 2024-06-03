package services

import (
	"github.com/tugsBoy/webApiWarriors/types"
)

// type ProcessLoginer interface {
// 	ProcessLogin(user types.User) (types.PlayerSave, error)
// }

//type processLoginImpl struct{}

func ProcessLogin(user types.User) (types.PlayerSave, error) {
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
