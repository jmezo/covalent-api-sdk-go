package tests

import (
	"testing"

	"github.com/jmezo/covalent-api-sdk-go/chains"
	"github.com/jmezo/covalent-api-sdk-go/testutil"
)

func TestApprovals(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.SecurityService.GetApprovals(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}

func TestNftApprovals(t *testing.T) {
	// Test case 1
	resp, err := testutil.Client.SecurityService.GetNftApprovals(chains.EthMainnet, "demo.eth")
	if err != nil {
		t.Errorf("error: %s", err)
	} else {
		if len(resp.Data.Items) == 0 {
			t.Errorf("Error - length of items is 0")
		}
	}
}
