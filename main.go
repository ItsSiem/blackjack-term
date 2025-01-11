package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	bj "github.com/kraanter/blackjack/pkg/blackjack"
	"blackjack-term/table"
)

var (
	width  = 80
	height = 24

	// Colors
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	// Styles
	header_style = lipgloss.NewStyle().
			Width(width).
			AlignHorizontal(lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderBottom(true)
)

type model struct {
	game          *bj.BlackjackGame
	player        *bj.Player
	selected_hand int
	kaasje        int
}

func initalModel() model {

	go func() {
		time.Sleep(10 * time.Millisecond)
		game.SetPlayerBet(player.PlayerNum, 25)
	}()
	game.Start()
	return model{
		game:   game,
		player: player,
	}
}

type update struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("kaas", msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch strings.ToLower(msg.String()) {
		case "q", "esc", "ctrl+c":
			// Exit dialog
			return m, tea.Quit
		case "left":
			if m.selected_hand-1 < 0 {
				//	m.selected_hand = len(m.player.Hands) - 1
			} else {
				m.selected_hand -= 1
			}
		case "right":
			//if (m.selected_hand + 1 >= len(m.player.hands)) {
			//	m.selected_hand = 0
			//}else {
			//m.selected_hand += 1
			//}

		// Game inputs
		case "h":
			m.game.PlayerHit(m.player.PlayerNum)
		case "s":
			m.game.PlayerStand(m.player.PlayerNum)
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
	return m, nil
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
	s = lipgloss.JoinVertical(lipgloss.Center, s, table.RenderHand(m.game.Dealer, true))
	return s
}

func print_player(m model) string {
	p_width := 60
	p := m.player
	//TODO multi hand cards := lipgloss.NewStyle().Margin(0, 2).Render(print_hand(p.hands[m.selected_hand].cards))
	cards := lipgloss.NewStyle().Margin(0, 2).Render(table.RenderHand(p.Hand, false))
	page := lipgloss.NewStyle().
		Foreground(subtle).
		Bold(true).
		// TODO multihand Render(fmt.Sprintf("%d/%d", m.selected_hand + 1, len(p.hands)))
		Render(fmt.Sprintf("%d/%d", 1, 1))
	page_spacer := lipgloss.NewStyle().
		Width((p_width-lipgloss.Width(cards))/2 - lipgloss.Width(page)).
		Render()
	betAmount := uint(0)
	if p.Hand != nil {
		betAmount = p.Hand.Bet
	}
	bet := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFBF00")).
		Bold(true).
		Render(fmt.Sprintf("%dg", betAmount))
	bet_spacer := lipgloss.NewStyle().
		Width(p_width/2 - lipgloss.Width(cards)/2 - lipgloss.Width(bet)).
		Render()
	balance := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFBF00")).
		Bold(true).
		Render(fmt.Sprintf("%dg", p.Balance))

	hr := lipgloss.NewStyle().
		Width(p_width).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(subtle).
		Render("")

	s := lipgloss.JoinHorizontal(lipgloss.Bottom, bet, bet_spacer, cards, page_spacer, page)
	s = lipgloss.JoinVertical(lipgloss.Center, s, hr)
	s = lipgloss.JoinVertical(lipgloss.Left, s, balance)
	return s
}

var game = bj.CreateGame(&bj.Settings{TimeBetweenRounds: 5 * time.Second, DealCardTime: 500 * time.Millisecond})
var player = game.AddPlayerWithBalance(100)

func main() {
	// set up logger
	f, _ := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetOutput(f)

	p := tea.NewProgram(initalModel(), tea.WithAltScreen())

	// set up game
	game.OnGameUpdate = func(bg *bj.BlackjackGame) {
		go func() {
			log.Printf("\n---\ngame_update: %v\n\nplayers:", game.GameState)

			log.Println(player.String())

			log.Println("Dealer: ", game.Dealer.String())

			if game.GameState == bj.NoState {
				go func() {
					time.Sleep(10 * time.Millisecond)
					game.SetPlayerBet(player.PlayerNum, 25)
				}()
				go game.Start()
			}
			p.Send(update{})
		}()
	}

	// start ui
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
