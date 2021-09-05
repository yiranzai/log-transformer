package cmd

import "testing"

func TestRun(t *testing.T) {
	cfgPath = "../conf/config.yaml"
	Run()
}
