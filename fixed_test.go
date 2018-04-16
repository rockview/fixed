package fixed

import (
    "testing"
    "math"
)

func TestIntConversions(t *testing.T) {
    tests := []int{0, 1, -1, 32767, -32768}

    for _, test := range tests {
        have := ToInt(FromInt(test))
        if have != test {
            t.Errorf("have: %d, want: %d", have, test)
        }
    }
}

func TestFloatConversions(t *testing.T) {
    tests := []float64{0.0, 1.0, -1.0, 32767.0, -32768.0}

    for _, test := range tests {
        have := ToFloat(FromFloat(test))
        if have != test {
            t.Errorf("have: %f, want: %f", have, test)
        }
    }
}

func TestMul(t *testing.T) {
    var tests = []struct {
        a, b int
    }{
        {1, 0},
        {1, 2},
        {1, -2},
        {-1, 2},
        {-1, -2},
        {128, 127},
        {128, -128},
    }

    for _, test := range tests {
        have := ToInt(Mul(FromInt(test.a), FromInt(test.b)))
        want := test.a*test.b
        if have != want {
            t.Errorf("Mul: %d, %d: have: %d, want: %d", test.a, test.b, have, want)
        }
    }
}

func TestDiv(t *testing.T) {
    var tests = []struct {
        a, b int
    }{
        {0, 1},
        {1, 1},
        {4, 2},
        {4, -2},
        {-4, 2},
        {-4, -2},
        {32767, 217},
        {-32768, 256},
    }

    for _, test := range tests {
        have := ToInt(Div(FromInt(test.a), FromInt(test.b)))
        want := test.a/test.b
        if have != want {
            t.Errorf("Div: %d, %d: have: %d, want: %d", test.a, test.b, have, want)
        }
    }
}

func TestMod(t *testing.T) {
    var tests = []struct {
        a, b int
    }{
        {1, 3},
        {2, 3},
        {3, 3},
        {1, -3},
        {2, -3},
        {3, -3},
        {-1, 3},
        {-2, 3},
        {-3, 3},
        {-1, -3},
        {-2, -3},
        {-3, -3},
        {32767, 217},
        {-32768, 256},
    }

    for _, test := range tests {
        have := ToInt(Mod(FromInt(test.a), FromInt(test.b)))
        want := test.a%test.b
        if have != want {
            t.Errorf("Mod: %d, %d: have: %d, want: %d", test.a, test.b, have, want)
        }
    }
}

func TestFloor(t *testing.T) {
    tests := []float64{1.5, -1.5, 32767.5, -32767.5}

    for _, test := range tests {
        have := ToInt(Floor(FromFloat(test)))
        want := int(math.Floor(test))
        if have != want {
            t.Errorf("have: %d, want: %d", have, want)
        }
    }
}

func TestCeil(t *testing.T) {
    tests := []float64{1.5, -1.5, 32766.5, -32767.5}

    for _, test := range tests {
        have := ToInt(Ceil(FromFloat(test)))
        want := int(math.Ceil(test))
        if have != want {
            t.Errorf("have: %d, want: %d", have, want)
        }
    }
}
