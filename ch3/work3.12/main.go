package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s,%s is chaos : %v\n", "test11", "test11", isChaos("test11", "test11"))
	fmt.Printf("%s,%s is chaos : %v\n", "test12", "test11", isChaos("test12", "test11"))
	fmt.Printf("%s,%s is chaos : %v\n", "ts1e1t", "test11", isChaos("ts1e1t", "test11"))
}

func isChaos(s1, s2 string) bool {
	return bubble(s1) == bubble(s2) && s1 != s2
}

func bubble(s string) string {
	tmp := []byte(s)
	for i := 0; i < len(tmp)-1; i++ {
		for j := i + 1; j < len(tmp); j++ {
			if tmp[i] > tmp[j] {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
	}
	return string(tmp)
}
