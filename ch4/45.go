package main

import "fmt"

func main() {
	s := []string {"123","233","233","test","asfdasdf","abc","abc","def"}
	s = deduplicate(s)
	fmt.Println(s)
	
}

func deduplicate(s []string) []string{
	l := len(s)
	for i :=0;i<l-1; {
		if s[i] == s[i+1] {
			copy(s[i:],s[i+1:])
			l--
			continue
		}
		i ++
	}
	return s[:l]
}
