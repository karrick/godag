package godag_test

import (
	"testing"

	"github.com/karrick/godag"
)

func helperStringSlices(tb testing.TB, a, b []string) {
	tb.Helper()

	if got, want := len(a), len(b); got != want {
		tb.Fatalf("GOT: %v; WANT: %v", got, want)
	}

	for i := 0; i < len(a) && i < len(b); i++ {
		if i >= len(a) {
			tb.Errorf("INDEX: %d; WANT: %v", i, b[i])
		} else if i >= len(b) {
			tb.Errorf("INDEX: %d; GOT: %v", i, a[i])
		} else if got, want := a[i], b[i]; got != want {
			tb.Errorf("INDEX: %d; GOT: %v; WANT: %v", 0, got, want)
		}
	}
}

func TestOrderEmpty(t *testing.T) {
	d := godag.New()

	got, err := d.Order()
	if err != nil {
		t.Fatal(err)
	}

	helperStringSlices(t, got, nil)
}

func TestOrderSingleItem(t *testing.T) {
	d := godag.New()

	d.Insert("foo", nil)

	got, err := d.Order()
	if err != nil {
		t.Fatal(err)
	}

	helperStringSlices(t, got, []string{"foo"})
}

func TestOrderAfterOrderedInsertion(t *testing.T) {
	d := godag.New()

	d.Insert("foo", nil)
	d.Insert("bar", []string{"foo"})

	got, err := d.Order()
	if err != nil {
		t.Fatal(err)
	}

	helperStringSlices(t, got, []string{"foo", "bar"})
}

func TestOrderAfterInvertedInsertion(t *testing.T) {
	d := godag.New()

	d.Insert("bar", []string{"foo"})
	d.Insert("foo", nil)

	got, err := d.Order()
	if err != nil {
		t.Fatal(err)
	}

	helperStringSlices(t, got, []string{"foo", "bar"})
}
