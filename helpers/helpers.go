package helpers

import "github.com/tugsBoy/webApiWarriors/types"

func InitMockDB() *types.MockDB {
	mockDB := &types.MockDB{
		PlayerSaves: []types.PlayerSave{
			{
				Warriors: []types.Warrior{
					{Name: "Aragorn", Sex: "Male", Age: "87", Strength: 20, Dex: "18", Intelligence: 15, Faith: 10, Arcane: 5, Luck: 12},
					{Name: "Legolas", Sex: "Male", Age: "2931", Strength: 18, Dex: "20", Intelligence: 16, Faith: 8, Arcane: 6, Luck: 14},
				},
			},
			{
				Warriors: []types.Warrior{
					{Name: "Gimli", Sex: "Male", Age: "139", Strength: 19, Dex: "14", Intelligence: 12, Faith: 7, Arcane: 5, Luck: 10},
					{Name: "Boromir", Sex: "Male", Age: "41", Strength: 18, Dex: "15", Intelligence: 14, Faith: 9, Arcane: 5, Luck: 12},
				},
			},
			{
				Warriors: []types.Warrior{
					{Name: "Frodo", Sex: "Male", Age: "50", Strength: 10, Dex: "12", Intelligence: 14, Faith: 18, Arcane: 5, Luck: 20},
					{Name: "Sam", Sex: "Male", Age: "38", Strength: 12, Dex: "11", Intelligence: 13, Faith: 19, Arcane: 5, Luck: 18},
				},
			},
		},
	}
	return mockDB
}
