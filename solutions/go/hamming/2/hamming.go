package hamming

import "errors"

func Distance(a, b string) (int, error) {
    aSize := len(a)
    bSize := len(b)

    if aSize != bSize {
        return 0, errors.New("a and b are not the same size")
    }

    hammingDistance := 0
    
    for i := range aSize {
    	if a[i] != b[i] {
            hammingDistance++
        }
    }

    return hammingDistance, nil
}
