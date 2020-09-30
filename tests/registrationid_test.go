package tests

import (
	"testing"

	"github.com/prospik/libsignal-protocol-go/util/keyhelper"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	t.Log(regID)
}
