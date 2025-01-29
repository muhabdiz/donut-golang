package main

import (
	"fmt"
	"math"
	"time"
)

const (
	width  = 80
	height = 22
)

func main() {
	a, b := 0.0, 0.0
	screen := make([]byte, width*height)
	zBuffer := make([]float64, width*height)

	for {
		for i := range screen {
			screen[i] = ' '
			zBuffer[i] = 0
		}

		for theta := 0.0; theta < 2*math.Pi; theta += 0.07 {
			for phi := 0.0; phi < 2*math.Pi; phi += 0.02 {
				c := math.Sin(phi)
				d := math.Cos(theta)
				e := math.Sin(a)
				f := math.Sin(theta)
				g := math.Cos(a)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(phi)
				m := math.Cos(b)
				n := math.Sin(b)
				t := c*h*g - f*e

				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				o := x + width*y
				N := int(8 * ((f*e - c*d*g) + l*d*m - t*n))

				if 0 <= y && y < height && 0 <= x && x < width && D > zBuffer[o] {
					zBuffer[o] = D
					charSet := ".,-~:;=!*#$@"
					if N > 0 && N < len(charSet) {
						screen[o] = charSet[N]
					} else {
						screen[o] = charSet[0]
					}
				}
			}
		}
		fmt.Print("\033[H")
		for i := 0; i < width*height; i++ {
			if i%width == 0 {
				fmt.Println()
			}
			fmt.Printf("%c", screen[i])
		}

		a += 0.04
		b += 0.08
		time.Sleep(15 * time.Millisecond)
	}
}

// percobaan
