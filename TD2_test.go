// TD2.go
package TD2

import (
	"fmt"
	_ "math/rand"
	"testing"
)

var s string

func TestZip(t *testing.T) {
	file_bac_sub := "NC_000964.fna"

	zip(file_bac_sub, s)

	//t.Logf("count %+v\n", on)

	zip(" ", s)
}

func TestEcrit_matrice(t *testing.T) {

	var matrice [][]float64
	matrice = make([][]float64, 15, 15)
	for i := 0; i < 15; i++ {
		matrice[i] = make([]float64, 15)
	}
	ecrit_matrice(matrice)
}
func TestAffiche_matrice(t *testing.T) {
	var matrice [][]float64
	matrice = make([][]float64, 15, 15)
	for i := 0; i < 4; i++ {
		matrice[i] = make([]float64, 15)
	}

	Affiche_matrice(matrice)
}

func TestMin_matrice(t *testing.T) {
	var matrice [][]float64
	var it *int
	var taille int
	matrice = make([][]float64, 15, 15)
	for i := 0; i < 15; i++ {
		matrice[i] = make([]float64, 15)
	}
	var vect = make([]int, 15)

	min_matrice(&matrice, &vect, it, taille)

	//min_matrice(matrice, 10)

}

func TestProcessFile(t *testing.T) {
	ProcessFile("NC_000964.fna")

}

func TestDistance(t *testing.T) {

	distance("NC_000964.fna", "NC_000964.fna", " ", " ")
	var matrice [][]float64
	matrice = make([][]float64, 15, 15)
	for i := 0; i < 15; i++ {
		matrice[i] = make([]float64, 15)
	}

}

func TestLinkage(t *testing.T) {
	nb_fichiers := 15
	var vect = make([]int, nb_fichiers)

	matrice := [][]float64{
		{0.000000, 0.000015, 0.006327, 0.004817, 0.000054, 0.000145, 0.000355, 0.000000, 0.000000, 0.006122, 0.003577, 0.008052, 0.002962, 0.009782, 0.001446},
		{0.000015, 0.000000, 0.006717, 0.004704, 0.000160, 0.000253, 0.000224, 0.000019, 0.000020, 0.005970, 0.003313, 0.008171, 0.002750, 0.006491, 0.001158},
		{0.006327, 0.006717, 0.000000, 0.000093, 0.004997, 0.005002, 0.008688, 0.006222, 0.006214, 0.001114, 0.000046, 0.000001, 0.000007, 0.000093, 0.011470},
		{0.004817, 0.004704, 0.000093, 0.000000, 0.004979, 0.004896, 0.008670, 0.006222, 0.004740, 0.000943, 0.008877, 0.000097, 0.000001, 0.000026, 0.002324},
		{0.000054, 0.000160, 0.004997, 0.004979, 0.000000, 0.000000, 0.001049, 0.000135, 0.000133, 0.005826, 0.003887, 0.007500, 0.003451, 0.008666, 0.002478},
		{0.000145, 0.000253, 0.005002, 0.004896, 0.000000, 0.000000, 0.001046, 0.000134, 0.000132, 0.005825, 0.003886, 0.005006, 0.003449, 0.004855, 0.002473},
		{0.000145, 0.000253, 0.005002, 0.004896, 0.000000, 0.000000, 0.001046, 0.000134, 0.000132, 0.005825, 0.003886, 0.005006, 0.003449, 0.004855, 0.002473},
		{0.000000, 0.000019, 0.006222, 0.006222, 0.000135, 0.000134, 0.014783, 0.000000, 0.000000, 0.005805, 0.003600, 0.006223, 0.002989, 0.006017, 0.001477},
		{0.000000, 0.000020, 0.006214, 0.004740, 0.000133, 0.000132, 0.000255, 0.000000, 0.000000, 0.006120, 0.003603, 0.008034, 0.002993, 0.009615, 0.001616},
		{0.006122, 0.005970, 0.001114, 0.000943, 0.005826, 0.005825, 0.006366, 0.005805, 0.006120, 0.000000, 0.001841, 0.002170, 0.002494, 0.003706, 0.007535},
		{0.003577, 0.003313, 0.000046, 0.008877, 0.003887, 0.003886, 0.002439, 0.003600, 0.003603, 0.001841, 0.000000, 0.000046, 0.007633, 0.000055, 0.000470},
		{0.008052, 0.008171, 0.000001, 0.000097, 0.007500, 0.005006, 0.008418, 0.006223, 0.008034, 0.002170, 0.000046, 0.000000, 0.010078, 0.000028, 0.007789},
		{0.002962, 0.002750, 0.000007, 0.000001, 0.003451, 0.003449, 0.001679, 0.002989, 0.002993, 0.002494, 0.007633, 0.010078, 0.000000, 0.000026, 0.000419},
		{0.009782, 0.006491, 0.000093, 0.000026, 0.008666, 0.004855, 0.008263, 0.006017, 0.009615, 0.003706, 0.000055, 0.000028, 0.000026, 0.000000, 0.010729},
		{0.001446, 0.001158, 0.011470, 0.002324, 0.002478, 0.002473, 0.000331, 0.001477, 0.001616, 0.007535, 0.000470, 0.007789, 0.000419, 0.010729, 0.000000},
	}
	/*var matrice [][]float64
	matrice = make([][]float64, 15, 15)
	for i := 0; i < 15; i++ {
		matrice[i] = make([]float64, 15)
	}
	for k := 0; k < 15; k++ {
		for j := 0; j < 15; j++ {

			//if k < j {
			matrice[k][j] = rand.Float64() * 0.8
			matrice[j][k] = matrice[k][j]

			if k == j {
				matrice[k][j] = 0.000000000000000
			}
			//}
		}
	}*/

	noms := []string{"Bacillus_subtilis", "Bacillus_amyloliquefaciens_FZB42", "Bacillus_pumilus_SAFR_032",
		"Bacillus_thuringiensis_BMB171", "Bacillus_cereus_03BB102", "Bacillus_anthracis_Ames",
		"Bacillus_coagulans_2_6", "Bacillus_atrophaeus_1942", "Bacillus_licheniformis_ATCC_14580",
		"Escherichia_coli_K_12_substr__MG1655", "Pseudomonas_aeruginosa_LESB58", "Rhodobacter_sphaeroides_ATCC_17025",
		"Streptomyces_flavogriseus_ATCC_33331", "Micrococcus_luteus_NCTC_2665_uid59033", "Lactococcus_lactis_Il1403"}

	Affiche_matrice(matrice)

	Linkage(&matrice, noms, nb_fichiers, &vect)

	for r := 0; r < nb_fichiers; r++ {
		fmt.Println(vect[r])
	}

}
