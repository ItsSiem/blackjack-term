package table

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	bj "github.com/kraanter/blackjack/pkg/blackjack"
)

var (
	// Styles
	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())
)

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
