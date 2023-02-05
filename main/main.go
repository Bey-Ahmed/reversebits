package main

import (
	"fmt"

	"piscine"
)

func main() {
	var oct byte = 38
	/*
	* Affichage de notre entier de type byte en binaire géré par printf
	* Le "b" est pour binary
	* Le 8 est le nombre d'espace que je lui alloue pour son affichage
	* Le 0 est la valeur par laquelle je comble les vides si les 8 espaces ne sont pas remplies
	*/
	fmt.Printf("%08b\n", oct)
	
	/* 
	* Conversion en binaire de notre entier de type byte
	* Changement de l'ordre
	* Conversion en entier du nouveau nombre binaire résultant
	* Récupération de ce nombre puis affichage en binaire
	*/
	reverse := piscine.ReverseBits(oct)
	fmt.Printf("%08b\n", reverse)
}
