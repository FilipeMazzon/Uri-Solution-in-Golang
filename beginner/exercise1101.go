package main

import "fmt"

func main() {
	var m ,n ,sum int

	fmt.Scan(&m,&n)

	for m>0 && n >0 {
		if m < n{
			m,n = n,m
		}
		for i:=n;i<=m;i++{
			fmt.Printf("%d ",i)
			sum+=i
		}
		fmt.Printf("Sum=%d\n",sum)
		sum=0
		fmt.Scan(&m,&n)
	}
}