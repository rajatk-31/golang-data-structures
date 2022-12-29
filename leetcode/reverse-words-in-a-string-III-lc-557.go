package main

import "bytes"

func main() {

}

// func reverseWords(s string) string {
//     splitted:=strings.Split(s," ")
//     for i := range splitted{
//         splitted[i] = reverseString(strings.Split(splitted[i],""))
//     }
//     return strings.Join(splitted, " ")
// }

// func reverseString(s []string) string {
//     for i:=0;i<len(s)/2;i++{
//         s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
//     }
//     return strings.Join(s, "")
// }

// without library and faster
func reverseWords(s string) string {
	var res bytes.Buffer
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			res.WriteByte(s[i])
			i++
			continue
		}
		j := i
		for j < len(s) && s[j] != ' ' {
			j++
		}
		for k := j - 1; k >= i; k-- {
			res.WriteByte(s[k])
		}
		i = j
	}
	return res.String()
}
