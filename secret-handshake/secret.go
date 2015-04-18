package secret

var signals = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code int) (hs []string) {
  switch {
    case code < 1 || code > 31:
    case code&16 == 0:
      for _,s := range signals {
        if code&1 != 0 {
          hs = append(hs, s)
        }
        code >>= 1
      }
    default:
      for i:=3; i>=0; i-- {
        if code&8 != 0 {
          hs = append(hs, signals[i])
        }
        code <<= 1
      }
  }
  return
}
