package main
import "fmt"

func isIsrat(x int, y int, z int) bool {
  // is Integer-Sided Right-Angle Triangle
  if x*x + y*y == z*z {
    return true
  } else {
    return false
  }
}

func main() {
  // const LMAX = 1500
  answer := 0
  var numTri [150000]int
  var numTriX [50000]int
  const LMAX = len(numTri)
  x := 0
  y := 0
  maxz := 0
  z := 0
  L := 0
  L3 := LMAX / 3
  buff := 0
  multiple := 0
  for x = 1; x < L3; x++ {
    if numTriX[x-1] > 1 {
      fmt.Println("skipped x:", x)
      continue
    }
    for y = L3; y > x; y-- {
      maxz = LMAX - x - y
      for z = y+1; z <= maxz; z++ {
        L = x + y + z
        if numTri[L-1] > 1 {
          continue
        }
        if isIsrat(x, y, z) {
          numTri[L-1]++
          numTriX[x-1]++
          multiple = 2
          buff = L * multiple
          for buff <= LMAX {
            numTri[buff-1]++
            numTriX[x*multiple-1]++
            multiple++
            buff = L * multiple
          }
        }
        if numTri[L-1] > 1 {
          break
        }
      }
      if numTri[L-1] > 1 {
        break
      }
    }
  }

  for _, val := range numTri {
    if val == 1 {
      answer++
    }
  }

  fmt.Println(answer)
}
