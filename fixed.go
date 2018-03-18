package fixed

import(
    "fmt"
)

type Fixed int32

var min Fixed
var max Fixed

const(
    Zero    = Fixed(0x00000000)
    One     = Fixed(0x00010000)
    Half    = Fixed(0x00008000)
    Third   = Fixed(0x00005555)
    Sixth   = Fixed(0x00002aaa)
)

func record(f Fixed) {
    if f > max {
        max = f
    } else if f < min {
        min = f
    }
}

func FromInt(i int) Fixed {
    result := Fixed(i<<16)
    record(result)
    return result
}

func ToInt(x Fixed) int {
    record(x)
    return int((x + Half)>>16)
}

func FromFloat(f float64) Fixed {
    result := Fixed(f*65536.0)
    record(result)
    return result
}

func ToFloat(x Fixed) float64 {
    record(x)
    return float64(x)/65536.0
}

func Mul(x, y Fixed) Fixed {
    result := Fixed((int64(x)*int64(y))>>16)
    record(result)
    return result
}

func Div(x, y Fixed) Fixed {
    result := Fixed((int64(x)<<16)/int64(y))
    record(result)
    return result
}

func Mod(x, y Fixed) Fixed {
    result := Fixed(int64(x)%int64(y))
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
    result := x&^Fixed(0xffff)
    record(result)
    return result
}

func Ceil(x Fixed) Fixed {
    result := x
    if uint32(x)&0xffff != 0 {
        result = Floor(x + One)
    }
    record(result)
    return result
}

func Round(x Fixed) Fixed {
    result := Floor(x + Half)
    record(result)
    return result
}

func ReportLimits() {
    fmt.Printf("Fixed.min: %f\n", ToFloat(min))
    fmt.Printf("Fixed.max: %f\n", ToFloat(max))
}
