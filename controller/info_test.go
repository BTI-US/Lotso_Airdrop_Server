package controller

import (
	"testing"
)

func TestApplyAirdrop(t *testing.T) {
	type test struct {
		name    string
		address string
	}

	var tests []test

	tests = append(tests, test{name: "1", address: "0x475Bc2321C3Dd6bf44D6276BF6F2ba578872aD24"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
