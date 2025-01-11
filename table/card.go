package table

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	bj "github.com/kraanter/blackjack/pkg/blackjack"
)

var (
	// Colors
	red   = lipgloss.Color("1") // TODO: figure out what ansi colors to use
	black = lipgloss.Color("2")

	// Styles
	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())

	heartsStyle = lipgloss.NewStyle().
			Foreground(red)
	diamondsStyle = heartsStyle
	spadesStyle   = lipgloss.NewStyle().
			Foreground(black)
	clubsStyle = spadesStyle
)

func GetFaceStyle(card *bj.Card) lipgloss.Style {
	switch card.Suit {
	case bj.Clubs:
		return clubsStyle
	case bj.Diamonds:
		return diamondsStyle
	case bj.Hearts:
		return heartsStyle
	case bj.Spades:
		return spadesStyle
	default:
		panic(fmt.Sprintf("unexpected blackjack.Suit: %#v", card.Suit))
	}
}

func RenderCard(card *bj.Card, facedown bool) string {
	if facedown {
		return cardStyle.Render("╭─╮\n╰─╯")
	}

	face := fmt.Sprintf("%v  \n  %v", card.Suit.String(), card.Face.String())
	if card.Face == bj.Ten {
		face = fmt.Sprintf("%v  \n %v", card.Suit.String(), card.Face.String())
	}

	return cardStyle.Render(face)
}
