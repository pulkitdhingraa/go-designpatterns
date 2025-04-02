package iterator

import "fmt"

func Run() {
	playlist := &Playlist{
		Songs: []*Song{
			{"Song A", "Artist 1"},
			{"Song B", "Artist 2"},
			{"Song C", "Artist 3"},
		},
	}

	iterator := playlist.NewIterator()
	fmt.Println("Playing in order")
	for iterator.HasNext() {
		fmt.Println(iterator.Next().Artist)
	}
}