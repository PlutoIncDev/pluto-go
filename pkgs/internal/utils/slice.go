package utils

func StringInSlice(a string, l []string) bool {
	for _, b := range l {
		if b == a {
			return true
		}
	}
	return false
}

func SlicesContainSameElement(l1 []string, l2 []string) bool {
	for _, l := range l1 {
		if StringInSlice(l, l2) {
			return true
		}
	}

	for _, l := range l2 {
		if StringInSlice(l, l1) {
			return true
		}
	}

	return false
}