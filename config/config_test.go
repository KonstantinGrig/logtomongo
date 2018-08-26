package config

import (
	"os"
	"testing"
)

func TestGetListFileInfo2(t *testing.T) {
	os.Setenv("LOGTOMONGO_ENV", "test")
	var conf = Config{}
	conf.Init()
	logFiles := conf.GetListFileInfo()
	var v = logFiles[0].Type
	if v != "first_format" {
		t.Error("Expected first_format, got ", v)
	}
	v = logFiles[1].Type
	if v != "second_format" {
		t.Error("Expected: second_format, got ", v)
	}
	v = logFiles[0].Path
	if v != "../test_data/first.log" {
		t.Error("Expected: ../test_data/first.log, got ", v)
	}
	v = logFiles[1].Path
	if v != "../test_data/second.log" {
		t.Error("Expected: ../test_data/second.log, got ", v)
	}
}
