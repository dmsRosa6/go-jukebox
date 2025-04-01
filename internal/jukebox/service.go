package jukebox

import (
	"errors"
	"fmt"
)

type Song struct {
	ID       int
	Title    string
	Artist   string
	FilePath string // location of the audio file
}

// Jukebox represents a circular queue of songs.
type Jukebox struct {
	head  int
	tail  int
	size  int
	songs []Song // using a slice to store songs
}

// NewJukebox initializes a new jukebox with a fixed capacity.
func NewJukebox(capacity int) *Jukebox {
	return &Jukebox{
		songs: make([]Song, capacity),
	}
}

// Enqueue adds a new song to the jukebox queue.
func (j *Jukebox) Enqueue(song Song) error {
	if j.size == len(j.songs) {
		return errors.New("jukebox is full")
	}
	j.songs[j.tail] = song
	j.tail = (j.tail + 1) % len(j.songs)
	j.size++
	return nil
}

// Dequeue removes and returns the song at the front of the queue.
func (j *Jukebox) Dequeue() (Song, error) {
	if j.size == 0 {
		return Song{}, errors.New("jukebox is empty")
	}
	song := j.songs[j.head]
	j.head = (j.head + 1) % len(j.songs)
	j.size--
	return song, nil
}

// Return all the songs on the jukebox playlist
func (j *Jukebox) GetJukeboxSongs() []Song {
    if j.size == 0 {
        return []Song{}
    }
    if j.head <= j.tail {
        return j.songs[j.head:j.tail]
    }
    return append(j.songs[j.head:], j.songs[:j.tail]...)
}


// PrintQueue is a helper to print the current state of the jukebox.
func (j *Jukebox) PrintQueue() {
	fmt.Println("Current Queue:")
	for i, count := j.head, 0; count < j.size; count++ {
		fmt.Printf("  %v\n", j.songs[i])
		i = (i + 1) % len(j.songs)
	}
}
