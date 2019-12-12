package TD2

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func zip(file_name, s string) int64 {
	fmt.Println("-- Début zippage " + file_name + " --")
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Name = (file_name + ".gzip")

	//Ecriture du fichier zippé
	_, err := zw.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(file_name+"_.gz", buf.Bytes(), 0666)

	//Recupération des stats sur le fichier
	fi, err := os.Stat(file_name + "_.gz")
	if err != nil {
		log.Fatal(err)
	}

	//Récupération de la taille du fichier zippé
	taille_zip := fi.Size()
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- Fin zippage " + file_name + " --")
	os.Remove(file_name + "_.gz")
	return taille_zip
}

func Affiche_matrice(matrice [][]float64) {
	for i := 0; i < 15; i++ {
		fmt.Println(matrice[i])
	}
}

// écrit la matrice dans un fichier txt

func ecrit_matrice(matrice [][]float64) {
	var matrice_string string
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			matrice_string = matrice_string + strconv.FormatFloat(matrice[i][j], 'f', 6, 64) + " "
		}
		matrice_string = matrice_string + "\n"
	}
	fmt.Println("Resultat :\n" + matrice_string)
	content := []byte(matrice_string)
	ioutil.WriteFile("matrice_distance.txt", content, 0640)
}

func distance(file_name1, file_name2, A, B string) float64 {
	fmt.Printf("-- Debut calcul distance(%s , %s) --\n", file_name1, file_name2)
	taille_Zip_A := zip(file_name1, A)
	taille_Zip_B := zip(file_name2, B)
	taille_Zip_AB := zip("A&B", A+B)
	taille_Zip_AA := zip("A&A", A+A)
	taille_Zip_BB := zip("B&B", B+B)
	fmt.Printf("-- Fin calcul distance(%s , %s) --\n", file_name1, file_name2)
	fmt.Printf("(A = %d, B = %d, A&B = %d, A&A = %d, B&B = %d)\n", taille_Zip_A, taille_Zip_B, taille_Zip_AB, taille_Zip_AA, taille_Zip_BB)
	var d float64
	op1 := float64(float64(taille_Zip_AB) / float64((taille_Zip_A + taille_Zip_B)))
	op2 := float64(float64(taille_Zip_AA) / float64((4 * taille_Zip_A)))
	op3 := float64(float64(taille_Zip_BB) / float64((4 * taille_Zip_B)))
	fmt.Printf("op1 = %f, op2 = %f, op3 = %f\n", op1, op2, op3)
	d = op1 - op2 - op3
	if d < 0 {
		return -d
	}
	return d
}

func ProcessFile(s string) string {
	var file_string string

	_, err := os.Stat(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "the file %s doesn't exist! \n", s)
		os.Exit(1)
	}

	file, _ := os.Open(s)

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	var lineNumber int
	var shortTreatment bool

	next_line_is_dna := false
	for scanner.Scan() {
		lineNumber++

		if shortTreatment && (lineNumber > 40000) {
			break
		}
		//var val_str string
		val_str := scanner.Text()
		if len(val_str) > 0 {
			if val_str[0] == '>' {
				next_line_is_dna = true
				continue
			} else if next_line_is_dna {
				file_string = file_string + val_str
				continue
			}
		}
	}

	return file_string
}

// cette fonction remplace dans la matrice

func remplimat(matrice *[][]float64, taille int, imin int, jmin int) {

	for i := 0; i < taille; i++ {
		for j := i; j < taille; j++ {
			if (*matrice)[i][imin] > (*matrice)[i][jmin] {

				(*matrice)[i][imin] = (*matrice)[i][jmin]

			}

		}
	}

}

func remplivec(vect *[]int, jinput int, it *int) {

	(*vect)[*it] = jinput

	*it++
}

func min_matrice(matrice *[][]float64, vect *[]int, it *int, taille int) (int, int, float64) {
	min := (*matrice)[0][1]
	ii_min := 0
	jj_min := 1
	for i := 0; i < taille; i++ {
		for j := i; j < taille; j++ {
			if (*matrice)[i][j] < min && i != j {

				for k := 0; k <= *it; k++ {
					if j == (*vect)[k] {
						break
					}
					//continue
					if j != (*vect)[k] && k == *it {
						min = (*matrice)[i][j]
						ii_min = i
						jj_min = j
					}
				}
			}
		}

	}

	//remplivec(vect, jj_min, it)

	return ii_min, jj_min, min
}

func Linkage(matrice *[][]float64, noms []string, taille int, vect *[]int) {
	chaine := ""
	a := 0
	b := 0
	c := 0
	var it *int = &a
	var i_min *int = &b
	var j_min *int = &c
	//*it = 0
	//var min float64
	//nb_colonne := taille

	for *it != taille-2 {

		*i_min, *j_min, _ = min_matrice(matrice, vect, it, taille)
		remplivec(vect, *j_min, it)

		remplimat(matrice, taille, *i_min, *j_min)

		chaine = chaine + "(" + noms[*i_min] + "," + noms[*j_min] + ")" + ";"
		//chaine = chaine + ";"

		//*it++
	}

	content := []byte(chaine)
	ioutil.WriteFile("chaine_netwick.txt", content, 0640)
}
