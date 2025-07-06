package main

import "fmt"

type Player interface {
	Play()
	Stop()
}

type Recorder interface {
	Record()
	StopRecording()
}

type PlayerRecorder interface {
	Player
	Recorder
}

type Microphone struct {
	tone string
}

func (m Microphone) Record() {
	fmt.Println("Recording from microphone", m.tone)
}

func (m Microphone) StopRecording() {
	fmt.Println("Stopping recording from microphone.")
}

type Radio struct {
	station string
}

func (r Radio) Play() {
	fmt.Println("Playing from radio: ", r.station)
}

func (r Radio) Stop() {
	fmt.Println("Stoppint radio: ", r.station)
}

type Smartphone struct {
	model string
}

func (s Smartphone) Play() {
	fmt.Println("Playing music on", s.model)
}

func (s Smartphone) Stop() {
	fmt.Println("Music stopped")
}

func (s Smartphone) Record() {
	fmt.Println("Recording video on", s.model)
}

func (s Smartphone) StopRecording() {
	fmt.Println("Recording saved")
}

func PlayMusic(p Player) {
	p.Play()
}

func RecordAudio(r Recorder) {
	r.Record()
}

func Execute(pr PlayerRecorder) {
	pr.Play()
	pr.Record()
	pr.Stop()
	pr.StopRecording()
}

func SmartPlay(p Player) {
	p.Play()

	if recorder, ok := p.(Recorder); ok {
		fmt.Println("I'm can record")
		recorder.Record()
	}
}

func main() {
	smartphone := Smartphone{"Iphone"}
	radio := Radio{"98 FM"}
	microphone := Microphone{"C"}
	Execute(smartphone)
	PlayMusic(radio)
	PlayMusic(smartphone)
	RecordAudio(microphone)
	RecordAudio(smartphone)
	SmartPlay(smartphone)
}
