package triangle

import "sort"
import "math"

type Kind string

const (
  Equ = "equilateral"
  Iso = "isosceles"
  Sca = "scalene"
  NaT = "not a triangle"
  )

func KindFromSides(a,b,c float64) Kind{
  if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
    return NaT
  }

  if math.IsInf(a,1) || math.IsInf(b,1) || math.IsInf(c,1) {
    return NaT
  }

  if math.IsInf(a,-1) || math.IsInf(b,-1) || math.IsInf(c,-1) {
    return NaT
  }

  s := []float64{a,b,c}
  sort.Float64s(s)

  switch {
    case s[0] + s[1] <= s[2]:
      return NaT
    case s[0] == s[1] :
      if s[1] == s[2] {
        return Equ
      } else {
        return Iso
      }
    case s[0] == s[2] || s[1] == s[2] :
      return Iso
    default:
      return Sca
  }
}
