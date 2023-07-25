package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRelayer(t *testing.T) {
	incomingChan := make(chan []byte)
	outgoingChan := make(chan []byte)
	doneChan := make(chan struct{})

	relayer := MessageRelayer{incoming: incomingChan, outgoing: outgoingChan, done: doneChan}

	go relayer.relay()
	incomingChan <- sampleContourData()
	outgoingData := <-outgoingChan
	close(incomingChan)

	got := outgoingData
	want := sampleContourData()
	if !cmp.Equal(got, want) {
		t.Errorf("Incoming and Outgoing Data, wanted\n %s\n got\n %s\n", string(want), string(got))
	}

}

func sampleContourData() []byte {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	data, err := os.ReadFile(strings.Join([]string{basepath, "../sample_contour.xml"}, "/"))
	if err != nil {
		log.Println("sample data:", err)
	}
	return data
}
