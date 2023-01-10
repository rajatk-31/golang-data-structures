package main

func isIsomorphic(s string, t string) bool {

	mapS := make(map[byte]byte)
	mapT := make(map[byte]byte)

	for i := range s {

		vS, vT := s[i], t[i]

		if m, ok := mapS[vS]; ok {
			if m != vT {
				return false
			}
		} else {
			mapS[vS] = vT
		}

		if m, ok := mapT[vT]; ok {
			if m != vS {
				return false
			}
		} else {
			mapT[vT] = vS
		}

	}

	return true
}
