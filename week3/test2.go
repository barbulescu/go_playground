package main
import(
  "fmt"
  "strconv"
  "sort"
)
func main() {
  s:=make([]int,3)
  var num string
  for i:=0;i<len(s);i++{
    fmt.Printf("enter a number ")
    fmt.Scanln(&num)
    if num=="X"{
      break
    }else{
      if num2,err:=strconv.Atoi(num);err==nil{
        s=append(s,num2)
        sort.Ints(s)
        fmt.Printf("%v\n",s[3:])

      }
    }
  }
}
