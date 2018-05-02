package main

import "fmt"

func main(){
	var nome,nomes string
	var numerosDeTitans, tamanhoMuralha ,tamanhoTitan int
	fmt.Scan(&numerosDeTitans,&tamanhoMuralha)
	for i := 0; i < numerosDeTitans; i ++ {
		fmt.Scan(&nome, &tamanhoTitan)

		if tamanhoTitan > tamanhoMuralha {
			fmt.Printf("%v %v\n", nome, nomes)
		}
	}



}
