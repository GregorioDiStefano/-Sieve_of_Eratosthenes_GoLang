package main
import "fmt"
import "math"

func get_primes(n int) []int {
    var max int = n + 1
    var primes []int
    array := make([]bool, max)

    for index, _ := range array {
        array[index] = true
    }

    max_index := int((math.Sqrt(float64(max))))

    for index, _ := range array {
        if (index <= 1) {
            continue
        } else if (index > max_index) {
            break
        } else {
            for j, _ := range array[index:] {
                v := int(math.Pow(float64(index), float64(2))) + (j * index)
                if (v > max - 1) {
                    break
                } else {
                    array[v] = false
                }
            }
        }
    }

    for index, element := range array {
        if (element) {
            primes = append(primes, index)
        }
    }

    return primes[2:]
}


func main() {
    fmt.Println(get_primes(30))
}
