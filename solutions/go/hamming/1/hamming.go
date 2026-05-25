package hamming

import "errors"

func Distance(a, b string) (int, error) {
    aSize := len(a)
    bSize := len(b)

    if aSize != bSize {
        return -1, errors.New("a and b are not the same size")
    }

    countDiff := 0
    
    for i := 0; i < aSize; i++ {
    	if a[i] != b[i] {
            countDiff += 1
        }
    }

    return countDiff, nil
}
