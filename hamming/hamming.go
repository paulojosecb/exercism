package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return distance, errors.New("different DNA sizes")
	}

	for index, _ := range a {
		if a[index] != b[index] {
			distance++
		}
	}

	return distance, nil
}
