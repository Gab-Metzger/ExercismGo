package diffsquares

func SquareOfSums(n int) int {
  var res int
  for i:=0; i <= n; i++ {
    res = res + i
  }
  return res*res
}

func SumOfSquares(n int) int {
  var res int
  for i:=0; i <= n; i++ {
    res = res + i*i
  }
  return res
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}
