package superb

import "testing"

func TestIs(t *testing.T) {
	if !Is("superb") {
		t.Fatal("superb is a superb word")
	}
	if Is("failure") {
		t.Fatal("failure is not a superb word")
	}
}

func TestNext(t *testing.T) {
	if !Is(Next()) {
		t.Fatal("Next doesn't return a superb word")
	}
}

func TestMultipleNextDifferent(t *testing.T) {
	if Next() == Next() {
		t.Fatal("multiple Next calls return same word")
	}
}
