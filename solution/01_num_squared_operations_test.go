package solution

import (
	"testing"
)

func TestNumSquaredOperations(t *testing.T) {

	if val := NumSquaredOperations(2, 10); val != 1 {
		t.Errorf("Maximum squared root operations number is 9. Sould be 1. Get %v.", val)
	}

	if val := NumSquaredOperations(6000, 7000); val != 3 {
		t.Errorf("Maximum squared root operations number is 6561. Sould be 3. Get %v.", val)
	}

	if val := NumSquaredOperations(2, 1000000); val != 4 {
		t.Errorf("Maximum squared root operations number is 65536. Sould be 4. Get %v.", val)
	}
}
