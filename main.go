package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type hand struct {
	cards []card
	bet int
}

func hit(h *hand) {
	h.cards = append(h.cards, card{Spades, Ace})
}

type player struct {
	hands []hand
	balance int
}

type model struct {
	dealer []card
	player *player
	selected_hand int
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

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch strings.ToUpper(msg.String()) {
		case "Q":
			// Exit dialog
			fmt.Println("Bye!")
			os.Exit(0)

		// Game inputs
		case "H":
			hit(&m.player.hands[m.selected_hand]);
		case "S":
			// Stand
		case "D":
			// Double
		case "P":
			// sPlit
		}
	}
	return m, nil;
}

func (m model) View() string {
	s := "Blackjack Terminal Edition\n\n"

	for _, card := range m.dealer {
		s += fmt.Sprintf("%s ", card)
	}
	s += "\n\n"
	for _, card := range m.player.hands[m.selected_hand].cards {
		s += fmt.Sprintf("%s ", card)
	}
	return s
}

func main() {
	p := tea.NewProgram(initalModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
