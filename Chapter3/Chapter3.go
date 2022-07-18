//These are Chapter3 examples of the book Go Programming Language

package main

import (
    "fmt"
    "math"
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "bytes"
    "strconv"
)

const  (
    width, height = 600, 320 //canvas size in pixels
    cells = 100 //number of grid cells
    xyrange  = 30.0 //axis rangers (-xyrange..+xyrange)
    xyscale = width / 2 / xyrange //pixels per x or y unit
    zscale = height * 0.4 //pixels per z unit
    angle = math.Pi // 6
    e  = 2.71828182845904523536028747135266249775724709369995957496696763
    pi = 3.14159265358979323846264338327950288419716939937510582097494459
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°

func main() {
	
	//The code below shows how bitwise operations can be used to interpret a uint8 value as a compact and efficient set of 8 independent bits.
	var x uint8 = 1<<1 | 1<<5
    var y uint8 = 1<<1 | 1<<2
    fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
    fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
    fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
    fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
    fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
    fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
    for i := uint(0); i < 8; i++ {
	   if x&(1<<i) != 0 { // membership test
           fmt.Println(i) // "1", "5"
        }
    }
    fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
	
	medals := []string{"gold", "silver", "bronze"}
    for i := len(medals) - 1; i >= 0; i-- {
         fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}


    // Surface computes an SVG rendering of a 3-D surface function.
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j)
            bx, by := corner(i, j)
            cx, cy := corner(i, j+1)
            dx, dy := corner(i+1, j+1)
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        } 
    }
    fmt.Println("</svg>")

    //complex numbers
    var b complex128 = complex(1, 2) // 1+2i
    var m complex128 = complex(3, 4) // 3+4i
    fmt.Println(b*m)      // "(-5+10i)"
    fmt.Println(real(b*m))   // "-5"
    fmt.Println(imag(b*m))  // "10"
    fmt.Println(1i * 1i) // "(-1+0i)"


    // Mandelbrot emits a PNG image of the Mandelbrot fractal.
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height          = 1024, 1024
    )
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value z.
            img.Set(px, py, mandelbrot(z))
        }
    }    
    png.Encode(os.Stdout, img) // NOTE: ignoring errors


    //boolean example
    t := 0
    r := 2
    fmt.Println(btoi(t,r))

    //string length  examples
    s := "hello, world"
    fmt.Println(len(s))     // "12"
    fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')
    fmt.Println(s[:5]) // "hello"
    fmt.Println(s[7:]) // "world"
    fmt.Println(s[:])  // "hello, world"

    //This does not modify the string that v originally held but causes v to hold the new string formed by the += statement;
    v := "left foot"
    l := v
    v += ", right foot"
    fmt.Println(v) // "left foot, right foot"
    fmt.Println(l) // "left foot"

    // comma inserts commas in a non-negative decimal integer string.
    number := "12345678910"
    fmt.Println(comma(number))

    // intsToString is like fmt.Sprintf(values) but adds commas.
    fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
    fmt.Println(intsToString([]int{10, 12, 13})) // "[10, 12, 13]"
    fmt.Println(intsToString([]int{14, 32, 23})) // "[14, 32, 23]"

    //To convert an integer to a string, one option is to use fmt.Sprintf; another is to use the function strconv.Itoa (‘‘integer to ASCII’’):
    w := 123
    q := fmt.Sprintf("%d", w)
    fmt.Println(q, strconv.Itoa(w)) // "123 123"
    fmt.Println(strconv.FormatInt(int64(w), 2)) // "1111011"
} 


//The function corner returns two values, the coordinates of the corner of the cell.
func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)
    // Compute surface height z.
    z := f(x, y)
    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}
func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    return math.Sin(r) / r
}


// Mandelbrot emits a PNG image of the Mandelbrot fractal.
func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15
    
    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}


// btoi returns 1 if b is true and 0 if false.
func btoi(b,a int) bool {
    if b<=a { 
        return true
    }
    return false
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    n := len(s)
    if n <= 3 {
      return s 
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

