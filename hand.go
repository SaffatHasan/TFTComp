package main

import (
	"fmt"
)

type Hand struct {
	cards int
	size  int
	pool  []card
}

func (h *Hand) addCard(cardIndex int) {
	// O(1)
	h.cards = setBit(h.cards, uint(cardIndex))
	h.size += 1
}

func (h *Hand) removeCard(cardIndex int) {
	// O(1)
	h.cards = unsetBit(h.cards, uint(cardIndex))
	h.size -= 1
}

func (h *Hand) hasCard(cardIndex int) bool {
	// O(1)
	return hasBit(h.cards, uint(cardIndex))
}

func (h *Hand) len() int {
	return h.size
}

func (h *Hand) copy() Hand {
	return Hand{h.cards, h.size, h.pool}
}

func (h *Hand) getValue() int {
	classMap := make(map[int]int)
	eleMap := make(map[int]int)

	for cardIndex := range h.pool {
		if !h.hasCard(cardIndex) {
			continue
		}
		classMap[h.pool[cardIndex].class] += 1
		eleMap[h.pool[cardIndex].element] += 1
	}

	missing := countMissing(classMap) + countMissing(eleMap)

	return h.len()*2 - missing
}

func countMissing(m map[int]int) (result int) {
	for _, val := range m {
		if val == 1 {
			result += 1
		}
	}
	return
}

func (h *Hand) String() string {
	result := "[ "
	for i := 0; i < len(h.pool); i++ {
		if !h.hasCard(i) {
			result = fmt.Sprintf("%s_ ", result)
		} else {
			result = fmt.Sprintf("%s%d ", result, i)
		}
	}
	return fmt.Sprintf("%s]", result)
}
