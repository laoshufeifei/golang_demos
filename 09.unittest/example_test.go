package testdemo

import (
	"math/rand"
	"testing"
)

func TestMyAdd(t *testing.T) {
	if MyAdd(1, 2) != 3 {
		t.Errorf("1 + 2 != 3")
	}
}

func TestMyAdds(t *testing.T) {
	var tests = []struct {
		param1 int
		param2 int
		expect int
	}{
		{4, 5, 9},
		{2, 4, 6},
	}

	for _, test := range tests {
		if ret := MyAdd(test.param1, test.param2); ret != test.expect {
			t.Errorf("%d + %d = %d != %d", test.param1, test.param2, ret, test.expect)
		}
	}
}

// benchmark test: function name must start with Becnhmark. Param must be type of *testing.B
func BenchmarkTestMyAdds(t *testing.B) {
	rand.Intn(100)

	for i := 0; i < 10000; i++ {
		p1 := rand.Intn(1000)
		p2 := rand.Intn(2000)
		expect := p1 + p2
		ret := MyAdd(p1, p2)
		if ret != expect {
			t.Errorf("MyAdd(%d, %d) = %d != expect(%d)", p1, p2, ret, expect)
		}
	}
}
