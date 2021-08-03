package main

import (
	"container/heap"
	"fmt"
)

func main() {
	pool := []card{
		{1, 1},
		{1, 2},
		{2, 1},
		{3, 1},
		{2, 4},
		{3, 2},
	}

	result := make(HandHeap, 0) // allocate storage space for hand
	start := Hand{0, 0, pool}
	visited := make(map[int]bool)
	maxCardsInHand := 4
	numResults := 10
	dfs(pool, start, &result, visited, maxCardsInHand, numResults)
	printHeapInDescendingOrder(result)
}

func dfs(pool []card, hand Hand, result *HandHeap, visited map[int]bool, maxCardsInHand, numResults int) {
	if visited[hand.cards] {
		return
	}

	visited[hand.cards] = true
	if hand.len() == maxCardsInHand {
		insertResult(hand.copy(), result, numResults)
		return
	}

	for cardIndex := range pool {
		if hand.hasCard(cardIndex) {
			continue
		}
		hand.addCard(cardIndex)
		dfs(pool, hand, result, visited, maxCardsInHand, numResults)
		hand.removeCard(cardIndex)
	}
}

// inserts a value into the heap maintaining maxSize
// if the heap is too large
func insertResult(new Hand, result *HandHeap, maxSize int) {
	heap.Push(result, new)
	if result.Len() > maxSize {
		heap.Pop(result)
	}
}

func printHeapInDescendingOrder(h HandHeap) {
	resultArr := make([]string, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		resultItem := heap.Pop(&h).(Hand)
		resultArr[i] = fmt.Sprintf("hand=%v, value=%d\n", resultItem.String(), resultItem.getValue())
	}

	for _, val := range resultArr {
		print(val)
	}
}
