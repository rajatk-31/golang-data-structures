package main

func main() {

}

func intToRoman(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hrns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ths := []string{"", "M", "MM", "MMM"}

	return ths[num/1000] + hrns[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}
