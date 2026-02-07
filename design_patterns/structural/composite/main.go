package main

import "fmt"

// Composite Pattern
//
// 5-Year-Old Explanation:
// Imagine a big box of toys.
// Inside the box, you can have a loose toy (like a car).
// OR you can have ANOTHER box inside it, which has more toys!
// Whether you look at a single toy or a whole box of toys, you can treat them the same way: "Count the price" or "Weight it".
// You don't care if it's one item or a box of items.

type Component interface {
	Search(keyword string)
}

// File is a simple leaf node (Start, End)
type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}

// Folder is a composite that can hold Files or other Folders
type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.Components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}

func main() {
	fmt.Println("--- Composite Pattern: Boxes inside Boxes ---")

	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{Name: "Folder1"}
	folder1.Add(file1)

	folder2 := &Folder{Name: "Folder2"}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1) // Adding a folder inside a folder!

	fmt.Println("\nSearching in the top-level folder (Folder2):")
	// One call to search triggers search in ALL sub-folders and files!
	folder2.Search("rose")
}
