package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck to be 52 but got %v", len(d))
	}

	if d[0] != "2 ♥" {
		t.Errorf("expected first card to be [2 ♥] but receiving %v", d[0])
	}

	if d[len(d)-1] != "A ♠" {
		t.Errorf("expected first card to be [A ♠] but receiving %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("__desktesting")

	d := newDeck()
	d.saveToFile("__desktesting")

	loadedDeck := newDeckFromFile("__desktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("__desktesting")
}
