// MIT License
//
// Copyright 2018 Jeremy Hall
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Quick and dirty 16.16 fixed point implementation to test suitability for use
// with the Simplex Noise Algorithm. No attempt is made to handle overflow or
// underflow.

package fixed

type Fixed int32

const(
    Zero    = Fixed(0x00000000)
    One     = Fixed(0x00010000)
    Half    = Fixed(0x00008000)
    Third   = Fixed(0x00005555)
    Sixth   = Fixed(0x00002aaa)
)

func FromInt(i int) Fixed {
    return Fixed(i<<16)
}

func ToInt(x Fixed) int {
    return int((x + Half)>>16)
}

func FromFloat(f float64) Fixed {
    return Fixed(f*65536.0)
}

func ToFloat(x Fixed) float64 {
    return float64(x)/65536.0
}

func Mul(x, y Fixed) Fixed {
    return Fixed((int64(x)*int64(y))>>16)
}

func Div(x, y Fixed) Fixed {
    return Fixed((int64(x)<<16)/int64(y))
}

func Mod(x, y Fixed) Fixed {
    return Fixed(int64(x)%int64(y))
}

func Abs(x Fixed) Fixed {
    if x < 0 {
        return -x
    }
    return x
}

func Floor(x Fixed) Fixed {
    return x&^Fixed(0xffff)
}

func Ceil(x Fixed) Fixed {
    if uint32(x)&0xffff != 0 {
        return Floor(x + One)
    }
    return x
}

func Round(x Fixed) Fixed {
    return Floor(x + Half)
}
