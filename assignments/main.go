package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	arraySign([]int{2, 1})                    
	arraySign([]int{-2, 1})                   
	arraySign([]int{-1, -2, -3, -4, 3, 2, 1})

	isAnagram("anak", "kana") 
	isAnagram("anak", "mana")
	isAnagram("anagram", "managra") 

	findTheDifference("abcd", "abcde") 
	findTheDifference("abcd", "abced")
	findTheDifference("", "y")

	canMakeArithmeticProgression([]int{1, 5, 3})   
	canMakeArithmeticProgression([]int{5, 1, 9})
	canMakeArithmeticProgression([]int{1, 2, 4, 8})

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	sign := 1

	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			sign = -sign
		}
	}
	fmt.Println(sign)

	return sign
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[rune]int)
	countT := make(map[rune]int)

	for _, char := range s {
		countS[char]++
	}

	for _, char := range t {
		countT[char]++
	}

	for key, val := range countS {
		if countT[key] != val {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	var result byte = 0

	for i := range s {
		result ^= s[i]
	}
	for i := range t {
		result ^= t[i]
	}

	return result
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	diff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}

type Deck struct {
	cards []Card
}

// Card represents a card in a "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New inserts 52 cards into deck d, sorted by symbol & then number.
func (d *Deck) New() {
	d.cards = make([]Card, 0, 52)
	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards = append(d.cards, Card{symbol: symbol, number: number})
		}
	}
}

// PeekTop returns n cards from the top
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekBottom returns n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex returns a card at the specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffles the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut performs a single "Cut" technique. Moves n top cards to the bottom
func (d *Deck) Cut(n int) {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

// ToString converts a card to its string representation
func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s of %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	fmt.Println("Peek top 3 cards:")
	top3Cards := deck.PeekTop(3)
	for _, c := range top3Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println("Peek card at index:")
	fmt.Println(deck.PeekCardAtIndex(12).ToString()) 
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) 
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) 
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) 
	fmt.Println("---\n")

	fmt.Println("Shuffling and peeking top 10 cards:")
	deck.Shuffle()
	top10Cards := deck.PeekTop(10)
	for _, c := range top10Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println("Cutting and peeking bottom 10 cards:")
	deck.New() 
	deck.Cut(5)
	bottom10Cards := deck.PeekBottom(10)
	for _, c := range bottom10Cards {
		fmt.Println(c.ToString())
	}
}
