package table

import (
	"github.com/charmbracelet/lipgloss"
	bj "github.com/kraanter/blackjack/pkg/blackjack"
)

const (
	blackjackColor = lipgloss.Color("5")
	standColor     = lipgloss.Color("0")
	bustColor      = lipgloss.Color("1")
	defaultColor   = lipgloss.Color("7")
)

func RenderHand(hand *bj.Hand, hidefirst bool) string {
	s := ""
	if hand == nil {
		return s
	}

	for i, c := range hand.Cards {
		if hidefirst && i == 0 {
			s = lipgloss.JoinHorizontal(lipgloss.Left, s, RenderCard(c, true))
			continue
		}
		s = lipgloss.JoinHorizontal(lipgloss.Left, s, RenderCard(c, false))
	}
	color := defaultColor
	if hand.IsLocked() {
		color = standColor
	}
	if hand.Total() > 21 {
		color = bustColor
	}
	if hand.Total() == 21 {
		color = blackjackColor
	}
	return lipgloss.NewStyle().Foreground(color).Render(s)
}
