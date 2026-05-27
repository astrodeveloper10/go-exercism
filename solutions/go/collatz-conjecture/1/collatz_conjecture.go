package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("Input cannot be less than or equal to 0")
    }
    
    count := 0

    for {
        if n == 1 {
            return count, nil
        }

        if n%2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
        
        count++
    }
}
