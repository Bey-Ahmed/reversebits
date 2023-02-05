package piscine

func ReverseBits(oct byte) byte {
	/*
	* Création d'un tableau (digits) de taille 8 car 1 octet = 8 bits
	* N.B. : octet en français = byte en anglais
	 */
	digits := make([]byte, 8)

	/*
	* La variable counter permettra de suivre le nombre d'éléments du tableau.
	* Si ce nombre est inférieur à 8 on ajoute des zéros afin de remplir toutes
	* les 8 places du tableau
	 */
	counter := 0

	/*
	* A chaque itération on divise le nombre donné en argument (oct) par 2 et on
	* garde le reste de cette division dans notre table digits (à la dernière position courante).
	* De ce fait, le tableau digits aura la conversion de oct en binaire mais dans l'ordre
	* inverse. Et vu qu'on a justement besoin de cet ordre inverse, on aura pas besoin de changer
	* l'ordre de digits.
	 */
	for oct != 0 {
		digits[counter] = byte(oct % 2)
		oct /= 2
		counter++
	}

	// Ajout des zéros pour les places restantes du tableau digits
	size := len(digits)
	for i := counter; i < size; i++ {
		digits[counter] = 0
	}

	// Conversion de en valeur décimale du nombre binaire obtenu dans notre tableau digits
	var reverse byte = 0
	for _, digit := range digits {
		reverse = reverse*2 + digit
	}

	// Retour de cette valeur qui est un entier de type byte i.e compris entre 0 et 2^8-1 (= 255)
	return reverse
}
