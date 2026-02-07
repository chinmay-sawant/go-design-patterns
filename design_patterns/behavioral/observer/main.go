package main

import "fmt"

// Observer Pattern
//
// 5-Year-Old Explanation:
// Imagine a Newspaper Delivery.
// The Newspaper Boy (Subject) has a list of houses (Observers).
// When a new paper is ready, he goes to EVERY house on his list and throws the paper.
// If you want the paper, you put your name on the list (Subscribe).
// If you don't want it anymore, you take your name off (Unsubscribe).

// Observer interface
type Observer interface {
	Update(news string)
}

// Subject interface
type Subject interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	NotifyAll()
}

// -- Concrete Observer --
type Reader struct {
	Name string
}

func (r *Reader) Update(news string) {
	fmt.Printf("%s recieved news: %s\n", r.Name, news)
}

// -- Concrete Subject --
type NewspaperAgency struct {
	subscribers []Observer
	latestNews  string
}

func (n *NewspaperAgency) Subscribe(o Observer) {
	n.subscribers = append(n.subscribers, o)
	fmt.Println("New subscriber added.")
}

func (n *NewspaperAgency) Unsubscribe(o Observer) {
	// (Simplification: Removing from slice logic logic omitted for brevity in 5yo example,
	// but logically this removes them from the list)
	fmt.Println("Subscriber removed (pretend logic).")
}

func (n *NewspaperAgency) NotifyAll() {
	for _, sub := range n.subscribers {
		sub.Update(n.latestNews)
	}
}

func (n *NewspaperAgency) PublishNews(text string) {
	fmt.Printf("\n--- Breaking News: %s ---\n", text)
	n.latestNews = text
	n.NotifyAll()
}

func main() {
	fmt.Println("--- Observer Pattern: Newspaper Delivery ---")

	agency := &NewspaperAgency{}

	// People sign up
	reader1 := &Reader{Name: "Alice"}
	reader2 := &Reader{Name: "Bob"}
	reader3 := &Reader{Name: "Charlie"}

	agency.Subscribe(reader1)
	agency.Subscribe(reader2)
	agency.Subscribe(reader3)

	// News happens!
	agency.PublishNews("Ice Cream is now FREE!")

	// More news!
	agency.PublishNews("School is closed tomorrow!")
}
