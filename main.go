package main
import "fmt"
import "time"
import "log"
import "math/rand"
var foundWaldo = make(chan int)
var gridSlices = make(chan []rune)
func main(){
  start := time.Now()
  fmt.Println(fatorial(20));
  log.Printf("demorou %s", time.Since(start))
  start = time.Now()
  findWaldo(hideWaldo(1000,1000));
  fmt.Printf("result: %d\n", <-foundWaldo)
  log.Printf("demorou %s", time.Since(start))
}
func fatorial(num uint64) uint64{
  if num==1 {
    return 1
  } else {
    return num * (fatorial(num-1))
  }
}
func hideWaldo(w int, h int)string{
    fmt.Printf("map size: %d\r\n",w*h)
    res := ""
    rand.Seed(time.Now().UnixNano())
    pos := rand.Intn((w*h)-1)
    fmt.Printf("waldo hide himself at %d\r\n",pos)
    for i:=0;i<(w*h);i++{
      pixel:="@"
      if i==pos{
        pixel="a"
      }
      res+=pixel
      // if i!=0 && (i+1)%w==0{
      //  res +="\r\n"
      // }
    }
    return res
}
func findWaldo(puzzle string){
  go chunkSlice([]rune(puzzle),1000)
  j:=0
  for i:=range gridSlices {
    go checkYourGrid(i,"a",j)
    j++
  }
}
func checkYourGrid(runes []rune, charToFind string, key int){
  charToFindRune := []rune(charToFind)[0]
  for i := 0; i < len(runes) ; i++ {
    if runes[i]==charToFindRune{
      fmt.Printf("found waldo at %d\r\n",(key*1000)+i)
      foundWaldo<-(key*1000)+i
      close(foundWaldo)
      close(gridSlices)
      //return i
    }
  }
  //return -1
}
func chunkSlice(slice []rune, chunkSize int) {
	for {
		if len(slice) == 0 {
			break
		}
		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
    gridSlices<-slice[0:chunkSize]
		slice = slice[chunkSize:]
	}
  close(gridSlices)
}
