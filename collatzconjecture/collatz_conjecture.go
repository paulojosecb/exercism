package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0

	if n <= 0 {
		return 0, errors.New("n is lower than 1")
	}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		count++
	}

	return count, nil
}
