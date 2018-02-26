package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseYamlSuccessfully(t *testing.T) {
	tempfile, _ := ioutil.TempFile("", "")
	tempfileName := tempfile.Name()

	defer os.Remove(tempfileName)
	defer tempfile.Close()

	ioutil.WriteFile(
		tempfileName,
		[]byte("address: hoge@example.com\npassword: xxx"),
		0644,
	)

	conf, err := readYamlConfig(tempfileName)
	if err != nil || conf.Address != "hoge@example.com" || conf.Password != "xxx" {
		t.Error("Failed to read the yaml file")
	}
}
