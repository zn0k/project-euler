package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Suit int

const (
	UndefinedSuit Suit = iota
	Clubs
	Diamonds
	Hearts
	Spades
)

type Face int

const (
	UndefinedFace Face = iota
	Two
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

type Card struct {
	face Face
	suit Suit
}

type Rank int

const (
	UndefinedRank Rank = iota
	HighCard
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type Score struct {
	rank Rank
	high Card
}

func (s Score) String() string {
	result := ""
	switch s.rank {
	case RoyalFlush:
		result += "Royal Flush "
	case StraightFlush:
		result += "Straight Flush "
	case FourOfAKind:
		result += "Four of a kind "
	case FullHouse:
		result += "Full House "
	case Flush:
		result += "Flush "
	case Straight:
		result += "Straight "
	case ThreeOfAKind:
		result += "Three of a kind "
	case TwoPair:
		result += "Two pairs "
	case Pair:
		result += "Pair "
	case HighCard:
		result += "nothing "
	}
	switch s.high.face {
	case Ace:
		result += "Ace"
	case King:
		result += "King"
	case Queen:
		result += "Queen"
	case Jack:
		result += "Jack"
	case Ten:
		result += "Ten"
	case Nine:
		result += "Nine"
	case Eight:
		result += "Eight"
	case Seven:
		result += "Seven"
	case Six:
		result += "Six"
	case Five:
		result += "Five"
	case Four:
		result += "Four"
	case Three:
		result += "Three"
	case Two:
		result += "Two"
	}
	result += " high"
	return result
}

type Hand struct {
	cards []Card
	score Score
}

func (h Hand) Len() int {
	return len(h.cards)
}

func (h Hand) Less(i, j int) bool {
	return h.cards[i].face < h.cards[j].face
}

func (h Hand) Swap(i, j int) {
	h.cards[i], h.cards[j] = h.cards[j], h.cards[i]
}

func (h Hand) String() string {
	result := ""
	for _, card := range h.cards {
		switch card.face {
		case Ace:
			result += "Ace"
		case King:
			result += "King"
		case Queen:
			result += "Queen"
		case Jack:
			result += "Jack"
		case Ten:
			result += "Ten"
		case Nine:
			result += "Nine"
		case Eight:
			result += "Eight"
		case Seven:
			result += "Seven"
		case Six:
			result += "Six"
		case Five:
			result += "Five"
		case Four:
			result += "Four"
		case Three:
			result += "Three"
		case Two:
			result += "Two"
		}
		result += " of "
		switch card.suit {
		case Clubs:
			result += "clubs, "
		case Diamonds:
			result += "diamonds, "
		case Hearts:
			result += "hearts, "
		case Spades:
			result += "spades, "
		}
	}
	return result
}

type RankEvaluator struct {
	rank Rank
	f    func(Hand) (bool, Card)
}

func parseCard(c string) Card {
	if len(c) != 2 {
		panic("Card string doesn't have exactly 2 characters")
	}
	cs := strings.Split(c, "")

	var card Card
	switch cs[0] {
	case "2":
		card.face = Two
	case "3":
		card.face = Three
	case "4":
		card.face = Four
	case "5":
		card.face = Five
	case "6":
		card.face = Six
	case "7":
		card.face = Seven
	case "8":
		card.face = Eight
	case "9":
		card.face = Nine
	case "T":
		card.face = Ten
	case "J":
		card.face = Jack
	case "Q":
		card.face = Queen
	case "K":
		card.face = King
	case "A":
		card.face = Ace
	default:
		panic("Invalid card face")

	}
	switch cs[1] {
	case "C":
		card.suit = Clubs
	case "D":
		card.suit = Diamonds
	case "H":
		card.suit = Hearts
	case "S":
		card.suit = Spades
	default:
		panic("Invalid card suit")
	}
	return card
}

func highestCard(hand Hand) Card {
	return hand.cards[len(hand.cards)-1]
}

func sameSuit(hand Hand) bool {
	suit := hand.cards[0].suit
	for i := 0; i < len(hand.cards); i++ {
		if hand.cards[i].suit != suit {
			return false
		}
	}
	return true
}

func faceCounts(hand Hand) map[Face][]Card {
	m := make(map[Face][]Card)
	for _, card := range hand.cards {
		_, ok := m[card.face]
		if !ok {
			m[card.face] = []Card{card}
		} else {
			m[card.face] = append(m[card.face], card)
		}
	}
	return m
}

func xOf(hand Hand, x int) (bool, Card) {
	fm := faceCounts(hand)
	for _, cards := range fm {
		if len(cards) == x {
			return true, cards[0]
		}
	}
	return false, Card{}
}

func straight(hand Hand) (bool, Card) {
	for i := 1; i < len(hand.cards); i++ {
		if hand.cards[i].face != hand.cards[i].face+1 {
			return false, Card{}
		}
	}
	return true, highestCard(hand)
}

func pair(hand Hand) (bool, Card) {
	return xOf(hand, 2)
}

func twoPair(hand Hand) (bool, Card) {
	fm := faceCounts(hand)
	foundOne := false
	var firstCard Card
	for _, cards := range fm {
		if len(cards) == 2 {
			if foundOne {
				if firstCard.face > cards[0].face {
					return true, firstCard
				} else {
					return true, cards[0]
				}
			} else {
				foundOne = true
				firstCard = cards[0]
			}
		}
	}
	return false, Card{}
}

func threeOf(hand Hand) (bool, Card) {
	return xOf(hand, 3)
}

func flush(hand Hand) (bool, Card) {
	if sameSuit(hand) {
		return true, hand.cards[len(hand.cards)-1]
	} else {
		return false, Card{}
	}
}

func fullHouse(hand Hand) (bool, Card) {
	isThreeOf, high3 := xOf(hand, 3)
	isPair, _ := xOf(hand, 2)
	if isThreeOf && isPair {
		return true, high3
	} else {
		return false, Card{}
	}
}

func fourOf(hand Hand) (bool, Card) {
	return xOf(hand, 4)
}

func straightFlush(hand Hand) (bool, Card) {
	isFlush := sameSuit(hand)
	isStraight, high := straight(hand)
	if isFlush && isStraight {
		return true, high
	} else {
		return false, Card{}
	}
}

func royalFlush(hand Hand) (bool, Card) {
	if !sameSuit(hand) {
		return false, Card{}
	}
	fm := faceCounts(hand)
	for _, f := range []Face{Ten, Jack, Queen, King, Ace} {
		if len(fm[f]) != 1 {
			return false, Card{}
		}
	}
	return true, highestCard(hand)
}

func scoreHand(hand Hand) Score {
	rankFunctions := []RankEvaluator{
		{RoyalFlush, royalFlush},
		{StraightFlush, straightFlush},
		{FourOfAKind, fourOf},
		{FullHouse, fullHouse},
		{Flush, flush},
		{Straight, straight},
		{ThreeOfAKind, threeOf},
		{TwoPair, twoPair},
		{Pair, pair},
	}
	for _, rf := range rankFunctions {
		rank, high := rf.f(hand)
		if rank {
			return Score{rf.rank, high}
		}
	}
	return Score{HighCard, highestCard(hand)}
}

func readHands(f string) ([]Hand, []Hand) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var p1, p2 []Hand
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 10 {
			panic("Line contains less than 10 fields")
		}
		var h1, h2 Hand
		for i := 0; i < len(fields); i++ {
			if i <= 4 {
				h1.cards = append(h1.cards, parseCard(fields[i]))
			} else {
				h2.cards = append(h2.cards, parseCard(fields[i]))
			}
		}
		sort.Sort(h1)
		sort.Sort(h2)
		h1.score = scoreHand(h1)
		h2.score = scoreHand(h2)
		p1 = append(p1, h1)
		p2 = append(p2, h2)
	}
	return p1, p2
}

func playGame(h1, h2 Hand) (int, int) {
	if h1.score.rank > h2.score.rank {
		//fmt.Println("Player 1 wins with higher rank")
		return 1, 0
	}
	if h2.score.rank > h1.score.rank {
		//fmt.Println("Player 2 wins with higher rank")
		return 0, 1
	}
	// same rank, highest card wins
	if h1.score.high.face > h2.score.high.face {
		//fmt.Println("Player 1 wins on same rank with high card")
		return 1, 0
	}
	if h2.score.high.face > h1.score.high.face {
		//fmt.Println("Player 2 wins on same rank with high card")
		return 0, 1
	}
	// same rank with same high card. filter out all shared faces, figure out highest of remainder
	var r1, r2 Hand
	fm1 := faceCounts(h1)
	fm2 := faceCounts(h2)
	for _, c := range h1.cards {
		_, ok := fm2[c.face]
		if !ok {
			r1.cards = append(r1.cards, c)
		}
	}
	sort.Sort(r1)
	for _, c := range h2.cards {
		_, ok := fm1[c.face]
		if !ok {
			r2.cards = append(r2.cards, c)
		}
	}
	sort.Sort(r2)
	if r1.cards[len(r1.cards)-1].face > r2.cards[len(r2.cards)-1].face {
		return 1, 0
	} else {
		return 0, 1
	}
}

func main() {
	p1, p2 := readHands("input.txt")
	if len(p1) != len(p2) {
		panic("The two players have a different number of hands")
	}
	wins1, wins2 := 0, 0
	for i := 0; i < len(p1); i++ {
		s1, s2 := playGame(p1[i], p2[i])
		winner := ""
		if s1 == 1 {
			winner = "1"
		} else {
			winner = "2"
		}
		fmt.Printf("Player 1 has hand %s (%s), player 2 has hand %s (%s), player %s wins\n", p1[i], p1[i].score, p2[i], p2[i].score, winner)
		wins1 += s1
		wins2 += s2
	}
	fmt.Println(wins1, wins2)
	var hand1 Hand
	hand1.cards = []Card{{Ace, Clubs}, {Ace, Hearts}, {Nine, Spades}, {Nine, Clubs}, {Four, Spades}}
	sort.Sort(hand1)
	hand1.score = scoreHand(hand1)

	var hand2 Hand
	hand2.cards = []Card{{Ace, Spades}, {Ace, Diamonds}, {Nine, Diamonds}, {Nine, Hearts}, {Five, Clubs}}
	sort.Sort(hand2)
	hand2.score = scoreHand(hand2)
	fmt.Printf("Test hand 1 with two pair is %s, score %s\n", hand1, hand1.score)
	fmt.Printf("Test hand 2 with two pair is %s, score %s\n", hand2, hand2.score)
	s1, s2 := playGame(hand1, hand2)
	if s1 == 1 {
		fmt.Println("Test hand 1 wins")
	}
	if s2 == 1 {
		fmt.Println("Test hand 2 wins")
	}
}
