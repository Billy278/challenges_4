package main

import (
	"challenges_4/member"
	"fmt"
	"os"
	"strconv"
)

func main() {

	search, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("input yang ada masukkan salah")
		return
	}

	memberRes := member.MemberImpl{}

	result, err := memberRes.GetMember(search)
	if err != nil {
		fmt.Printf("Data Absen yang anda Input %v", err)
		return
	}
	fmt.Printf("data Absen Member ke :%v \n %+v\n", search, result)

}
