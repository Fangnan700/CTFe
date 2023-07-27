package main

import (
	"CTFe/internal/server"
	"testing"
)

func TestMain(m *testing.M) {
	server.SetCTFeTokenStatus("111111", 1)
}
