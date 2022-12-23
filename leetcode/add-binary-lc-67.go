package main

func main() {

}

func addBinary(a string, b string) (out string) {
	a, b, out = reverse(a), reverse(b), ""
	carry := false
	for i := 0; i < len(a) || i < len(b); i++ {
		sum := 0
		if i < len(a) && a[i] == '1' {
			sum += 1
		}
		if i < len(b) && b[i] == '1' {
			sum += 1
		}
		if carry {
			sum += 1
		}
		if sum%2 == 0 {
			out += "0"
		} else {
			out += "1"
		}
		carry = sum > 1
	}
	if carry {
		out += "1"
	}
	return reverse(out)
}

func reverse(s string) (t string) {
	for _, r := range s {
		t = string(r) + t
	}
	return
}
