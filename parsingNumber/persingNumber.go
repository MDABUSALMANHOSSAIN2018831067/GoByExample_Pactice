package main

import (
    "fmt"
    "strconv"
)

func main() {

    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)
	fmt.Printf("%T\n",f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)
	fmt.Printf("%T\n",i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)
	fmt.Printf("%T\n",d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)
	fmt.Printf("%T\n",u)

    k, _ := strconv.Atoi("135")
    fmt.Println(k)
	fmt.Printf("%T\n",k)

    v, e := strconv.Atoi("wat")
    fmt.Println(e)
	fmt.Printf("%T\n",v)


	ii := 10
	s1 := strconv.FormatInt(int64(ii), 10)
	s2 := strconv.Itoa(ii)
	fmt.Printf("%v, %v\n", s1, s2)
	fmt.Printf("%T   %T\n",s1,s2)
}