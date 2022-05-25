package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides looks at size of each size, determines what type of triangle, if it is one.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	k = NaT

	if a > 0 && b > 0 && c > 0 {
		if (a+b >= c) && (a+c >= b) && (b+c >= a) {
			if a == b && b == c {
				k = Equ
			} else if (a == b && a != c) || (a == c && a != b) || (b == c && b != a) {
				k = Iso
			} else if a != b && a != c && b != c {
				k = Sca
			}
		}
	}
	return k
}
