package mad

import (
	"testing"
	"fmt"
	"github.com/cvanderschuere/alsa-go"
	"time"
)

func Test(t *testing.T) {
	const filename = "test.mp3"

	decoder, err := New(filename)
	if (err != nil) {
		t.Errorf("Failed to inialize decoder. %s", err)
	}
	defer decoder.Close()

	fmt.Printf("File name: %s\n", filename)
	fmt.Printf("Length: %d\n", decoder.Length())
	fmt.Printf("Sample rate: %d\n", decoder.SampleRate())
	fmt.Printf("Number of channels: %d\n", decoder.Channels())

	//Open ALSA pipe
	controlChan := make(chan bool)
	streamChan := alsa.Init(controlChan)
	
	//Create stream
	dataChan := make(chan alsa.AudioData, 20)
	current_stream := alsa.AudioStream{Channels:2, Rate:int(16000),SampleFormat:alsa.INT16_TYPE, DataStream:dataChan}

	streamChan<-current_stream
	
	//Pause initially
	controlChan<-true
	time.Sleep(2*time.Second)
	//controlChan<-false
	
	buf := make([]byte, 4096)
	fmt.Printf("Decoded time:     0")
	for decoder.Read(buf) > 0 {
		current_stream.DataStream<-buf
		fmt.Printf("\b\b\b\b\b%5d", decoder.CurrentPosition())
	}
	fmt.Printf("\n")
	
	time.Sleep(10*time.Second)
}