package main

import "fmt"

// Media is the type for ABSTRACT factory
type Media interface {
	Plays()
	IsAvailable() bool
}

// Music is the CONCRETE factory for musics
type Music struct {
}

// implement a concrete factory for musics
func (m *Music) Plays() {
	fmt.Println("Playing music")
}

func (m *Music) IsAvailable() bool {
	return true
}

// Television is the CONCRETE factory for televisions
type Television struct {
}

// implement concrete factory for televisions
func (t *Television) Plays() {
	fmt.Println("Playing television")
}

func (t *Television) IsAvailable() bool {
	return false
}

type MediaFactory interface {
	New() Media
}

type MusicFactory struct{}

func (mf *MusicFactory) New() Media {
	return &Music{}
}

type TelevisionFactory struct{}

func (tf *TelevisionFactory) New() Media {
	return &Television{}
}

func main() {

	// create families of related objects

	// create one each of MusicFactory and TelevisionFactory
	musicFactory := MusicFactory{}
	televisionFactory := TelevisionFactory{}

	// create the New method to create music and television
	music := musicFactory.New()
	television := televisionFactory.New()

	music.Plays()
	television.Plays()

	fmt.Println("Music is available", music.IsAvailable())
	fmt.Println("Television is available", television.IsAvailable())
}
