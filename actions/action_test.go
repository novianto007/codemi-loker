package actions

import (
	"codemi/loker/helpers"
	"strings"
	"testing"
)

type mockWriter struct {
	MsgTemp string
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.MsgTemp = string(p)
	return len(p), nil
}

func TestInit(t *testing.T) {
	writer := mockWriter{}
	helpers.PrintInit(&writer)
	Init([]string{})
	if !strings.Contains(writer.MsgTemp, "1 parameter") {
		t.Fatal("giving wrong message on parameter not valid ", writer.MsgTemp)
	}
	Init([]string{"notvalid"})
	if !strings.Contains(writer.MsgTemp, "angka") {
		t.Fatal("giving wrong message on parameter not number ", writer.MsgTemp)
	}
	Init([]string{"10"})
	if !strings.Contains(writer.MsgTemp, "10") {
		t.Fatal("giving wrong message on init success ", writer.MsgTemp)
	}
	Init([]string{"10"})
	if !strings.Contains(writer.MsgTemp, "sudah") {
		t.Fatal("giving wrong message on reinit ", writer.MsgTemp)
	}
}

func TestStatus(t *testing.T) {
	writer := mockWriter{}
	helpers.PrintInit(&writer)
	Status([]string{})
	if !strings.Contains(writer.MsgTemp, "kosong") {
		t.Fatal("give wrong message on status when not init")
	}
}
