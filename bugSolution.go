func processData(data []int) {  for i := 0; i < len(data); i++ {    if data[i] == 0 && i+1 < len(data){      fmt.Println(data[i+1])    }  }} 