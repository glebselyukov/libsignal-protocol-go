package tests

import (
	"fmt"
	"testing"

	"github.com/prospik/libsignal-protocol-go/util/keyhelper"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
}
