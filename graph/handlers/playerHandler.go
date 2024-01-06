package handlers

import "github.com/Rayato159/gqlgen-practice/graph/model"

var players = []*model.Player{
	{
		ID:    "PLAYER-001",
		Name:  "Rayato",
		Level: 1,
		Class: &model.AllPlayerClass[1],
		Items: []*model.Item{
			{
				ID:   "ITEM-001",
				Name: "Sword",
			},
			{
				ID:   "ITEM-002",
				Name: "Potion",
			},
		},
	},
	{
		ID:    "PLAYER-002",
		Name:  "DancingWithMyGraph",
		Level: 99,
		Class: &model.AllPlayerClass[0],
		Items: []*model.Item{
			{
				ID:   "ITEM-002",
				Name: "Potion",
			},
		},
	},
}

func GetPlayers() []*model.Player {
	return players
}

func CreatePlayer(req *model.NewPlayer) *model.Player {
	newPlayer := &model.Player{
		ID:    "PLAYER-003",
		Name:  req.Name,
		Level: 1,
		Class: &req.Class,
		Items: []*model.Item{
			{
				ID:   "ITEM-002",
				Name: "Potion",
			},
		},
	}

	players = append(players, newPlayer)
	return newPlayer
}
