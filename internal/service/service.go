package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dmsrosa/jukebox/config"
)

// Song is a simple representation of a song.
type Song struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// Jukebox represents a circular queue of songs.
type Jukebox struct {
	head  int
	tail  int
	size  int
	songs []Song
}

// NewJukebox initializes a new jukebox with a fixed capacity.
func NewJukebox(options config.Options) *Jukebox {
	return &Jukebox{
		songs: make([]Song, options.Capacity),
	}
}

func (j *Jukebox) Current() (Song, error){
	if j.head == j.tail {
		return Song{},errors.New("empty queue")
	}
	return j.songs[j.head], nil
}

// Next makes it so the current song playing is skipped
func (j *Jukebox) Next() (Song, error){
	
	temp := (j.head + 1) % len(j.songs)

	if(temp == j.tail){
		return Song{},errors.New("no next")
	}

	j.head = temp

	return j.songs[temp], nil
}

// Shuffle shuffles the whole active part of the queue
func (j *Jukebox) Shuffle() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < j.size; i++ {
        r := rand.Intn(j.size)
        j.songs[(j.head+i)%len(j.songs)], j.songs[(j.head+r)%len(j.songs)] = j.songs[(j.head+r)%len(j.songs)], j.songs[(j.head+i)%len(j.songs)]
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
func (j *Jukebox) GetSongs() []Song {
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
