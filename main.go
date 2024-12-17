package main

import (
	//tea "github.com/charmbracelet/bubbletea"
	"fmt"
	//"os"
)

type hand struct {
	cards []card
	bet int
}

type player struct {
	hands []hand
	balance int
}

type model struct {
	dealer []card
	player *player
}
func initalModel() model {
	return model {
		dealer: []card{
			{Hearts, Two},
			{Spades, King},
		},
		player: &player{
			hands: []hand{
				{
					bet: 10,
					cards: []card{
						{Diamonds, Seven},
						{Spades, Nine},
						{Diamonds, Ace},
					},
				},
			},
		},
	}
}

func main() {
	model := initalModel()
	fmt.Printf("%s\n", model.dealer[0])	
}
