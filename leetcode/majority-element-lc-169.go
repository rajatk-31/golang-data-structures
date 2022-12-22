package main

func main() {

}

// func majorityElement(nums []int) int {
//     counts:=0
//     var no int
//     sort.Slice(nums, func(i, j int) bool {
//     return nums[i] < nums[j]
// })
//     var lastVal int
//     lastCount:=0
//     for _,v:= range nums{

//	        if counts>len(nums)/2{
//	            return no
//	        }
//	        if lastVal==v{
//	            lastCount++;
//	        }else{
//	            if lastCount>counts {
//	                counts = lastCount
//	                no = lastVal
//	            }
//	            lastCount=1
//	            lastVal= v
//	        }
//	    }
//	    if lastCount>counts {
//	                counts = lastCount
//	                no = lastVal
//	            }
//	    return no
//	}
func majorityElement(nums []int) int {
	cnt := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			result = nums[i]
			cnt++
		} else if nums[i] == result {
			cnt++
		} else {
			cnt--
		}
	}
	return result
}
