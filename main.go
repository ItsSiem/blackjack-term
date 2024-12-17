package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	width = 80
	height = 24

	// Colors
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	table     = lipgloss.Color("#0c3008")

	// Styles
	header_style = lipgloss.NewStyle().
		Width(width).
		AlignHorizontal(lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderBottom(true)
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
				{
					bet: 250,
					cards: []card{
						{Diamonds, Ace},
						{Diamonds, Six},
					},
				},
			},
			balance: 100,
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
		switch strings.ToLower(msg.String()) {
		case "q", "esc", "ctrl+c":
			// Exit dialog
			return m, tea.Quit
		case "left":
			if (m.selected_hand - 1 < 0) {
				m.selected_hand = len(m.player.hands) - 1
			}else {
				m.selected_hand -= 1
			}
		case "right":
			if (m.selected_hand + 1 >= len(m.player.hands)) {
				m.selected_hand = 0
			}else {
				m.selected_hand += 1
			}

		// Game inputs
		case "h":
			hit(&m.player.hands[m.selected_hand]);
		case "s":
			// Stand
		case "d":
			// Double
		case "p":
			// sPlit
		}
	case tea.WindowSizeMsg:
		// TODO: responsive?
		//width = msg.Width
		//height = msg.Height
	}
	return m, nil;
}

func (m model) View() string {
	s := header_style.Render("Blackjack Terminal Edition") + "\n"

	s = lipgloss.JoinVertical(lipgloss.Center, s, print_dealer(m))

	hr := lipgloss.NewStyle().
		Width(50).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(subtle).
		Render("")

	s = lipgloss.JoinVertical(lipgloss.Center, s, hr)

	s += "\n\n\n\n\n"

	s = lipgloss.JoinVertical(lipgloss.Center, s, print_player(m))

	s = lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(subtle).
		Render(s)
	return s
}

func print_dealer(m model) string {
	s := lipgloss.NewStyle().Faint(true).Render(fmt.Sprintf("Dealer:"))
	s = lipgloss.JoinVertical(lipgloss.Center, s, print_cards(m.dealer))
	return s
}

func print_cards(cards []card) string {
	s := ""
	for _, c := range cards {
		s = lipgloss.JoinHorizontal(lipgloss.Top, s, print_card(c))
	}
	return s
}

func print_card(c card) string {
	s := lipgloss.NewStyle().
		Bold(true).
		Align(lipgloss.Top, lipgloss.Left).
		Render(c.suit.String())
	s += "\n"
	s += lipgloss.NewStyle().
		Bold(true).
		MarginLeft(2).
		Render(c.rank.String())
	s = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Render(s)
	return s
}

func print_player(m model) string {
	p_width := 60
	p := m.player
	cards := lipgloss.NewStyle().Margin(0, 2).Render(print_cards(p.hands[m.selected_hand].cards))
	page := lipgloss.NewStyle().
		Foreground(subtle).
		Bold(true).
		Render(fmt.Sprintf("%d/%d", m.selected_hand + 1, len(p.hands)))
	page_spacer := lipgloss.NewStyle().
		Width((p_width - lipgloss.Width(cards)) / 2 - lipgloss.Width(page)).
		Render();
	bet := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFBF00")).
		Bold(true).
		Render(fmt.Sprintf("%dg", p.hands[m.selected_hand].bet))
	bet_spacer := lipgloss.NewStyle().
		Width(p_width / 2 - lipgloss.Width(cards) / 2 - lipgloss.Width(bet)).
		Render()
	balance := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFBF00")).
		Bold(true).
		Render(fmt.Sprintf("%dg", p.balance))

	hr := lipgloss.NewStyle().
		Width(p_width).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(subtle).
		Render("")

	s := lipgloss.JoinHorizontal(lipgloss.Bottom, bet, bet_spacer , cards, page_spacer, page)
	s = lipgloss.JoinVertical(lipgloss.Center, s, hr)
	s = lipgloss.JoinVertical(lipgloss.Left, s, balance)
	return s
}

func main() {
	p := tea.NewProgram(initalModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
