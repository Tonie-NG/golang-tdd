package iteration


func Repeat(a string, x int) string {
  var rstring string;

  for i := 0; i < x; i++ {
    rstring += a
  }
  return rstring
}
