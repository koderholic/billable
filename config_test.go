package main

import (
	"billable/config"
	"testing"
)

func TestConfigInitReadsAndLoadsConfigFileCorrectly(t *testing.T) {

	Config := config.Data{}
	Config.Init("test")
	logPath := Config.LogPath

	if logPath != "testing" {
		t.Errorf("Expected logPath to be %s. Got %s\n", "testing", logPath)
	}

}
