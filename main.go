package main

import ("fmt"
		"strconv"
)

func main(){
	var first2name string = "Ri"
	var last2name string
	last2name = "co"
	var lastname = "Basyar"
	a,b,c := "Rico", "Basyar", "Barwaqa"
	x,y,z,q := 1,3,5,"Empat"


	fmt.Printf("halo %s %s!\n", first2name+last2name, lastname)
	fmt.Printf("%s %s %s", a, b, c + "h\n")
	fmt.Println(strconv.Itoa((x+y)*z/5) + q)

}


