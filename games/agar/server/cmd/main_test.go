package main

// TODO (hub33k): setup test workflow

// https://go.dev/doc/tutorial/add-a-test

import (
	"reflect"
	"testing"

	"server/pkg/packets"

	"google.golang.org/protobuf/proto"
)

func TestPacketMarshal(t *testing.T) {
	original := &packets.Packet{
		SenderId: 69,
		Msg:      packets.NewChat("Hello, world!"),
	}

	data, err := proto.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal packet: %v", err)
	}

	expectedBytes := []byte{8, 69, 18, 15, 10, 13, 72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}

	if !reflect.DeepEqual(data, expectedBytes) {
		t.Errorf("Marshalled bytes do not match expected.\nGot:      %v\nExpected: %v", data, expectedBytes)
	}
}

func TestPacketUnmarshal(t *testing.T) {
	data := []byte{8, 69, 18, 15, 10, 13, 72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	expected := &packets.Packet{
		SenderId: 69,
		Msg:      packets.NewChat("Hello, world!"),
	}

	var result packets.Packet
	err := proto.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal packet: %v", err)
	}

	if !proto.Equal(&result, expected) {
		t.Errorf("Unmarshalled packet does not match expected.\nGot:      %v\nExpected: %v", &result, expected)
	}
}

func TestHelloEmpty(t *testing.T) {
}
