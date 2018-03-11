package fixed

import(
    "fmt"
    "math"
)

type Fixed float64

var min Fixed
var max Fixed

const(
    Zero    = Fixed(0.0)
    One     = Fixed(1.0)
    Half    = Fixed(0.5)
    Third   = Fixed(1.0/3.0)
    Sixth   = Fixed(1.0/6.0)
)

func record(f Fixed) {
    if f > max {
        max = f
    } else if f < min {
        min = f
    }
}

func FromInt(i int) Fixed {
    result := Fixed(i)
    record(result)
    return result
}

func ToInt(x Fixed) int {
    record(x)
    return int(math.Round(float64(x)))
}

func FromFloat(f float64) Fixed {
    result := Fixed(f)
    record(result)
    return result
}

func ToFloat(x Fixed) float64 {
    record(x)
    return float64(x)
}

func Mul(x, y Fixed) Fixed {
    result := Fixed(x*y)
    record(result)
    return result
}

func Div(x, y Fixed) Fixed {
    result := Fixed(x/y)
    record(result)
    return result
}

func Abs(x Fixed) Fixed {
    result := x
    if x < 0 {
        result = -x
    }
    record(result)
    return result
}

func Floor(x Fixed) Fixed {
    result := Fixed(math.Floor(ToFloat(x)))
    record(result)
    return result
}

func Ceil(x Fixed) Fixed {
    result := Fixed(math.Ceil(ToFloat(x)))
    record(result)
    return result
}

func Round(x Fixed) Fixed {
    result := Fixed(math.Round(ToFloat(x)))
    record(result)
    return result
}

func Mod(x, y Fixed) Fixed {
    result := Fixed(math.Mod(ToFloat(x), ToFloat(y)))
    record(result)
    return result
}

func ReportLimits() {
    fmt.Printf("Fixed.min: %f\n", ToFloat(min))
    fmt.Printf("Fixed.max: %f\n", ToFloat(max))
}
