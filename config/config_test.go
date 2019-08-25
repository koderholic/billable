package config

import (
	"io/ioutil"
	"testing"
)

type Test struct {
	In         string
	Out        string
	ConfigFile string
}

var test = Test{
	ConfigFile: "../test/config.yaml",
	In:         "../test",
	Out:        "../log.json",
}

func TestConfigLoadsCorrectlyWithConfigPath(t *testing.T) {

	Config := Data{}
	Config.Init(test.In)
	logPath := Config.LogPath

	if logPath != test.Out {
		t.Errorf("Expected logPath to be %s. Got %s\n", test.Out, logPath)
	}

}

func TestConfigLoadsCorrectlyWithoutConfigPath(t *testing.T) {

	Config := Data{}
	Config.Init("")
	logPath := Config.LogPath

	if logPath != test.Out {
		t.Errorf("Expected logPath to be %s. Got %s\n", test.Out, logPath)
	}

}

func TestConfigReturnsErrorForInvalidConfigFile(t *testing.T) {

	rewriteFileContent := "testConfigReload"
	err := ioutil.WriteFile(test.ConfigFile, []byte("logPath :"+rewriteFileContent), 0644)
	if err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

	Config := Data{}
	err = Config.Init(test.In)
	if err == nil {
		t.Errorf("Expected returned error. Got %s\n", err)
	}

	rewriteFileContent = "../log.json"
	if err := ioutil.WriteFile(test.ConfigFile, []byte("logPath: "+rewriteFileContent), 0644); err != nil {
		t.Errorf("Test terminated. Reason >> %s\n", err)
	}

}
