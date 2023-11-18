package handler

import "testing"

func TestShowTable(t *testing.T) {
	var info []Info
	var err error
	if info, err = ShowTable(); err != nil {
		t.Error(err)
	}

	t.Log(info)
}
