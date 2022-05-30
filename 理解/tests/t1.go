package main

import "fmt"

func main() {
	s1 := "hello你好"
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(len([]byte(s1)))

	for _, char := range s1 {
		fmt.Print(char)
		fmt.Println()
	}
	fmt.Print("--------byte存储---------")
	fmt.Println()
	for _, bchar := range []byte(s1) {
		fmt.Print(bchar)
		fmt.Println()
	}
	fmt.Print("--------rune存储---------")
	fmt.Println()
	for _, rchar := range []rune(s1){
		fmt.Print(rchar)
		fmt.Println()
	}
}

// select count(*) from MME where  NE_NAME ="YACMME01BHW"  between  "20200801000000" and "20201001000000" ;

// select count(*) from MME where NE_NAME ="YACMME01BHW"  and DATETIME_KEY in ("20200801000000", "20201001000000");

// select count(*) from MME where NE_NAME ="YACMME01BHW"  and (DATETIME_KEY > "20200801000000" and  DATETIME_KEY< "20201001000000"));