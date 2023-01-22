package main

import "bytes"

// func decodeMessage(key string, message string) string {
//     x:='a'
//     all := map[string]rune{}
//     for _,v:= range key{
//         y:= string(v);
//         if v!=' '{
//             if _,ok:= all[y]; !ok{
//             all[y] =x
//             x++;
//         }}
//         // fmt.Println(i,v)

//	        // fmt.Println(x,">>>", )
//	    }
//	    final := ""
//	    for _,v := range message{
//	        if v==' '{
//	            final = final + " ";
//	        }else{
//	             final = final + string(all[string(v)]);
//	        }
//	    }
//	    // fmt.Println(final)
//	    return final
//	}
func decodeMessage(key string, message string) string {
	table := map[rune]byte{' ': ' '}
	for _, c := range key {
		if _, ok := table[c]; !ok {
			table[c] = byte(len(table) - 1 + 'a')
		}
	}

	var buffer bytes.Buffer
	for _, c := range message {
		buffer.WriteByte(table[c])
	}

	return buffer.String()
}
