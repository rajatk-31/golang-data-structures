package main

func main() {

}

// func strStr(haystack string, needle string) int {
//     found:=false
//     index:=-1
//     i:=0;

//     var other []int
//     for i<len(haystack) {
//         if haystack[i]== needle[0] && index ==-1 &&!found {
//             index = i
//             found = true
//         }
//         if found && haystack[i]== needle[0] && index !=i {
//             other = append(other, i)
//         }
//         if found {
//             ind:=i-index
//             if ind >= len(needle){
//                 break;
//             }
//             if haystack[i]!= needle[ind]{
//                 index = -1
//                 found = false
//             }
//         }
//         if i== len(haystack)-1 && i-index<len(needle)-1{
//             index = -1
//              found = false
//         }

//	        i++;
//	        if !found && len(other)>0{
//	            i = other[0]
//	            other = other[1:]
//	        }
//	    }
//	    return index
//	}
func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}

	return -1
}
