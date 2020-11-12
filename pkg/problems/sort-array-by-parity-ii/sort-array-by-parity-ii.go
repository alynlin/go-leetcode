package main

import "fmt"

/**
2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000
*/
func sortArrayByParityII(A []int) []int {

	jarray := []int{}
	oarray := []int{}

	for _, val := range A {
		if val%2 == 0 {
			oarray = append(oarray, val)
		} else {
			jarray = append(jarray, val)
		}
	}

	res := []int{}
	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			if i == 0 {
				res = append(res, oarray[0])
			} else {
				if i/2 > len(oarray)-1 {
					res = append(res, oarray[0])
				} else {
					res = append(res, oarray[i/2])
				}
			}
		} else {
			if i/2 > len(jarray)-1 {
				res = append(res, jarray[0])
			} else {
				res = append(res, jarray[i/2])
			}
		}
	}

	return res
}

func main() {
	a := []int{4, 1, 5, 7}
	r := sortArrayByParityII(a)

	fmt.Println(r)

}
