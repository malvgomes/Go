package mathematics

import "testing"

const defaultErr = "Expected value: %v. Received value: %v"

func TestAvg(t *testing.T) {
	expValue := 7.28
	value := Avg(7.2, 9.9, 6.1, 5.9)

	if value != expValue {
		t.Errorf(defaultErr, expValue, value)
	}
}
