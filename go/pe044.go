package main
import "fmt"

func genPentagon(n int) int {
  return n * (3 * n - 1) / 2
}

func findInSlice(slice []int, element int) bool {
  for _, el := range slice {
    if el == element {
      return true
    }
  }
  return false
}

func main() {
  // Calculate and cache pentagon numbers in advance
  N := 3000
  pentagons := make([]int, 0)
  nPentMap := make(map[int]int)
  for i := 1; i<N; i++ {
    buff := genPentagon(i)
    pentagons = append(pentagons, buff)
    nPentMap[i] = buff
  }

  D := 10000000000
  for j := 1; j < N; j++ {
    for k := j+1; k < N; k++ {
      pj := nPentMap[j]
      pk := nPentMap[k]
      sum := pj + pk
      diff := pk - pj
      if findInSlice(pentagons, sum) && findInSlice(pentagons, diff) {
        if diff < D {
          D = diff
        }
      }
    }
  }
  fmt.Println(D);
}
