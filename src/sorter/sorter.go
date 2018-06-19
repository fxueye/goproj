package main
import (
    "flag"
    "fmt"
    "bufio"
    "io"
    "os"
    "strconv"
    "time"
    "algorithms/qsort"
    "algorithms/bubblesort"
)
var infile *string = flag.String("i", "infile", "File")
var outfile *string = flag.String("o", "outfile", "out file")
var algorithm *string = flag.String("a", "qsort", "sort")
func main() {
    flag.Parse()
    if infile != nil{
        fmt.Println("infile=",*infile,"outfile=",*outfile,"algorithm=",*algorithm)
    }
    values , err := readVlaues(*infile)
    if err == nil{
        fmt.Println("read values:",values)
        t1 := time.Now()
        switch *algorithm{
        case "qsort":
          qsort.QuickSort(values)
        case "bubblesort":
          bubblesort.Bubblesort(values)
        default:
          fmt.Println("sorting algorithm",*algorithm)
        }
        t2 := time.Now()
        fmt.Println("sorting process costs",t2.Sub(t1),"to complete")
        writeValues(values, *outfile)
    }else{
        fmt.Println(err)
    }
        

}
func readVlaues(infile string) (values []int, err error){
    file,err := os.Open(infile)
    if err != nil{
        fmt.Println("Failed to open:",infile)
        return
    }
    defer file.Close()
    br := bufio.NewReader(file)
    values = make([]int ,0)
    for{
       line , isPrefix,err1 := br.ReadLine()
       if err1 != nil{
        if err1 != io.EOF{
            err = err1
        }
        break
       }
       if isPrefix{
        fmt.Println("A too long line")
        return
       } 
       str := string(line)
       value,err1 := strconv.Atoi(str)
       if err1 != nil{
        err = err1
        return
       }
       values = append(values,value)
    }
    return
}
func writeValues(values []int , outfile string) error {
  file, err := os.Create(outfile)
  if err != nil{
    fmt.Println("failed to create!",outfile)
    return err
  }
  defer file.Close()
  for _,value := range values{
    str := strconv.Itoa(value)
    file.WriteString(str + "\n")
  }
  return nil
}
