package main

import "fmt"

func main() {

}

func hammingWeight(num uint32) int {

	count := uint32(0)
	fmt.Println(count, num)
	for ; num > uint32(0); num >>= 1 {
		count += num & 1
	}

	return int(count)
}

// func hammingWeight(num uint32) int {
//     t := strconv.FormatUint(uint64(num), 2)
//     count:=0;
//     for _,v:=range t{
//         if(v=='1'){
//             count++;
//         }
//     }
//     return count
// }
