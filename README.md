# Optimal TFT Composition

This is a playground repository meant to solve a niche problem.

## Problem Statement

### Define
- **Card** has two attributes: a class and an element. They do not overlap.
- **Hand** is a collection of cards
    The value of a hand is the total number of combos for the classes and elements.
    e.g. if we consider the two cards {warrior, fire} and {assassin, fire} the two elements overlap and so this hand has a value of 2.
    e.g. {warrior, fire}, {assassin, fire}, {warrior, water} -> value = 4
    e.g. {assassin, water}, {warrior, fire}, {warrior, water} -> value = 5

### Given
- A number of cards per hand
- A number of best results
- A pool of cards to choose from (without replacement)

## Output
- A list in descending order of the best combinations

