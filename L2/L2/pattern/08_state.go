package main

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Состояние позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.

	+
	1. Вместо использования большого количества условных операторов (if-else) для управления поведением объекта, 
	   можно использовать состояния, что делает код более понятным и поддерживаемым.

	-
	1. Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/


import (
	"fmt"
)

type PlaybackState interface {
	Play()
	Pause()
	Stop()
}

type PlayingState struct {
	audioPlayer *AudioPlayer
}

func (ps *PlayingState) Play() {
	fmt.Println("Audio is already playing.")
}

func (ps *PlayingState) Pause() {
	fmt.Println("Pausing the audio.")
	ps.audioPlayer.ChangeState(ps.audioPlayer.pausedState)
}

func (ps *PlayingState) Stop() {
	fmt.Println("Stopping the audio.")
	ps.audioPlayer.ChangeState(ps.audioPlayer.stoppedState)
}

type PausedState struct {
	audioPlayer *AudioPlayer
}

func (ps *PausedState) Play() {
	fmt.Println("Resuming playback.")
	ps.audioPlayer.ChangeState(ps.audioPlayer.playingState)
}

func (ps *PausedState) Pause() {
	fmt.Println("Audio is already paused.")
}

func (ps *PausedState) Stop() {
	fmt.Println("Stopping the audio.")
	ps.audioPlayer.ChangeState(ps.audioPlayer.stoppedState)
}

type StoppedState struct {
	audioPlayer *AudioPlayer
}

func (ss *StoppedState) Play() {
	fmt.Println("Starting playback.")
	ss.audioPlayer.ChangeState(ss.audioPlayer.playingState)
}

func (ss *StoppedState) Pause() {
	fmt.Println("Audio is stopped. Cannot pause.")
}

func (ss *StoppedState) Stop() {
	fmt.Println("Audio is already stopped.")
}

type AudioPlayer struct {
	playingState PlaybackState
	pausedState  PlaybackState
	stoppedState PlaybackState
	currentState PlaybackState
}

func NewAudioPlayer() *AudioPlayer {
	playingState := &PlayingState{}
	pausedState := &PausedState{}
	stoppedState := &StoppedState{}

	audioPlayer := &AudioPlayer{
		playingState: playingState,
		pausedState:  pausedState,
		stoppedState: stoppedState,
		currentState: stoppedState,
	}

	playingState.audioPlayer = audioPlayer
	pausedState.audioPlayer = audioPlayer
	stoppedState.audioPlayer = audioPlayer

	return audioPlayer
}

func (ap *AudioPlayer) ChangeState(state PlaybackState) {
	ap.currentState = state
}

func (ap *AudioPlayer) Play() {
	ap.currentState.Play()
}

func (ap *AudioPlayer) Pause() {
	ap.currentState.Pause()
}

func (ap *AudioPlayer) Stop() {
	ap.currentState.Stop()
}

func main() {
	audioPlayer := NewAudioPlayer()

	audioPlayer.Play()
	audioPlayer.Pause()
	audioPlayer.Play()
	audioPlayer.Stop()
}
