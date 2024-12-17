package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type card struct {
	suit suit
	rank rank 
}
func (c card) String() string {
	var style lipgloss.Style
	switch c.suit {
	case Hearts:
		style = hearts_style
	case Diamonds:
		style = diamonds_style
	case Spades:
		style = spades_style
	case Clubs:
		style = clubs_style
	}
	return style.Render(fmt.Sprintf("%s%s", c.suit, c.rank))
}

type suit int
const (
	Hearts = iota
	Diamonds
	Spades
	Clubs
)
func (s suit) String() string {
	switch s {
	case Hearts:
		return hearts_style.Render("♥")
	case Diamonds:
		return diamonds_style.Render("♦")
	case Spades:
		return spades_style.Render("♠")
	case Clubs:
		return clubs_style.Render("♣")
	default:
		panic(fmt.Errorf("Invalid suit"))
	}
}

type rank int
const (
	Two = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)
func (r rank) String() string {
	switch r {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		panic(fmt.Errorf("Invalid rank"))
	}
}

var hearts_style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#ff007c"))
var diamonds_style = hearts_style
var spades_style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#09031c"))
var clubs_style = spades_style
