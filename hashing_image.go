package main

import (
	"fmt"
	"github.com/corona10/goimagehash"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	//	files, err := ioutil.ReadDir(".")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	filesToAnalyse := []string{}
	//	for _, file := range files {
	//		filesToAnalyse = append(filesToAnalyse, file.Name())
	//		//fmt.Println(file.Name())
	//	}

	filesToAnalyse := []string{
		"TRYP_Madrid_Atocha_Hotel-002106a_hb_a_937.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_a_938.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_a_939.jpg",
		"TRYP_Madrid_Atocha_Hotel-002106a_hb_a_940.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_002.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_003.jpg",
		"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_005.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_007.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_008.jpg",
		"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_009.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_010", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_011",
		//"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_012", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_013",
		//"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_017", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_018", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_019",
		//"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_020",
	}

	res, err := CompareImages(filesToAnalyse)
	if res == nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}
	if res != nil {
		fmt.Println(res)
	}
	//fmt.Println(res)

	os.Exit(0)

}

///func isElementExist(s []string, str string) bool {
///	for _, v := range s {
///		if v == str {
///			return true
///		}
///	}
///	return false
///}

func CompareImages(listImageFiles []string) ([]string, error) {
	var ListDebut []string
	var ListFin []string
	ListFin = listImageFiles
	IndexCheck := 0
	seuil := 20
	for IndexCheck < len(ListFin) {
		fmt.Println("=======================================")
		ListDebut = ListFin
		fmt.Println("INDEX CHECK + LISTFIN :: ", IndexCheck, ListFin)
		ListFin = []string{}
		for IndexSkip := 0; IndexSkip <= IndexCheck; IndexSkip++ {
			ListFin = append(ListFin, ListDebut[IndexSkip])
			fmt.Println("IMAGES TO SKIP :: ", ListFin)
		}
		file1, _ := os.Open(ListDebut[IndexCheck])
		IndexCheck += 1
		img1, _ := jpeg.Decode(file1) //TODO: Faire une function ici qui utilise le bon decode pour l'extension
		hash1, _ := goimagehash.AverageHash(img1)
		ElementToCompare := hash1
		for IndexCompare := IndexCheck; IndexCompare < len(ListDebut); IndexCompare++ {
			file2, _ := os.Open(ListDebut[IndexCompare])
			img2, _ := jpeg.Decode(file2)
			hash2, _ := goimagehash.AverageHash(img2)
			HashImageEnQuestion := hash2
			distance, _ := ElementToCompare.Distance(HashImageEnQuestion)
			if distance >= seuil {
				ListFin = append(ListFin, ListDebut[IndexCompare])
			}
		}
	}
	fmt.Println("///////////// RETURN VALUE =======================================")
	fmt.Println("INDEX CHECK + LISTFIN :: ", IndexCheck, ListFin)
	return ListFin, nil
}

/*func CompareImages(listImageFiles []string) ([]string, error){
	var ListDebut []string
	var ListFin []string
	ListFin = listImageFiles
	IndexCheck := 0
	seuil := 20
	for IndexCheck < len(ListFin) {
		ListDebut = ListFin
		ListFin = []
		for IndexSkip in range(IndexCheck){
			ListFin = append(ListFin, ListDebut[IndexSkip])
			UniqueImages = append(UniqueImages, other_image)
		}
		IndexCheck += 1
		file1, _ := os.Open(ListDebut[IndexSkip])
		img1, _ := jpeg.Decode(file1)
		hash1, _ := goimagehash.AverageHash(img1)
		defer file1.Close()
		ElementToCompare =  hash1
		for IndexCompare in range(IndexSkip + 1, len(ListDebut)) {
			file2, _ := os.Open(ListDebut[IndexCompare])
			img2, _ := jpeg.Decode(file2)
			hash2, _ := goimagehash.AverageHash(img2)
			defer file1.Close()
			HashImageEnQuestion = hash2
			distance, _ := ElementToCompare.Distance(HashImageEnQuestion)
			if(distance <= seuil){
				ListFin = append()
			}
		}
	}
	return ListFin, nil
}*/

//	file1, _ := os.Open("cat_V1.jpg")
//	file2, _ := os.Open("cat_V1.jpg")
//	file3, _ := os.Open("cat_V2.jpg")
//	file4, _ := os.Open("dog.jpg")
//	defer file1.Close()
//	defer file2.Close()
//	defer file3.Close()
//	defer file4.Close()

//	img1, _ := jpeg.Decode(file1)
//	img2, _ := jpeg.Decode(file2)
//	img3, _ := jpeg.Decode(file3)
//	img4, _ := jpeg.Decode(file4)
//	hash1, _ := goimagehash.AverageHash(img1)
//	hash2, _ := goimagehash.AverageHash(img2)
//	hash3, _ := goimagehash.AverageHash(img3)
//	hash4, _ := goimagehash.AverageHash(img4)
//	distance, _ := hash1.Distance(hash2)
//	distance_2, _ := hash1.Distance(hash3)
//	distance_3, _ := hash1.Distance(hash4)
//	fmt.Printf("Distance between identique images with AverageHash: %v\n", distance)
//	fmt.Printf("Distance between similar images with AverageHash : %v\n", distance_2)
//	fmt.Printf("Distance between different images with AverageHash: %v\n", distance_3)

//	hash1, _ = goimagehash.DifferenceHash(img1)
//	hash2, _ = goimagehash.DifferenceHash(img2)
//	hash3, _ = goimagehash.DifferenceHash(img3)
//	hash4, _ = goimagehash.DifferenceHash(img4)
//	distance, _ = hash1.Distance(hash2)
//	distance_2, _ = hash1.Distance(hash3)
//	distance_3, _ = hash1.Distance(hash4)
//	fmt.Printf("Distance between identique images with DifferenceHash: %v\n", distance)
//	fmt.Printf("Distance between similar images with DifferenceHash: %v\n", distance_2)
//	fmt.Printf("Distance between different images with DifferenceHash: %v\n", distance_3)
