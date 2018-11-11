package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p(s.Contains("test", "es"))
	p(s.Count("test", "t"))
	p(s.HasPrefix("test", "te"))
	p(s.HasSuffix("test", "st"))
	p(s.Index("test", "e"))
	p(s.Join([]string{"a", "b"}, "-"))
	p(s.Repeat("a", 5))
	p(s.Replace("foo", "o", "0", -1))
	p(s.Replace("foo", "o", "0", 1))
	p(s.Split("a-b-c-d-e", "-"))
	p(s.ToLower("TEST"))
	p(s.ToUpper("test"))
	p()
}