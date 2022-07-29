ackage main

import (
	"fmt"
	"image/jpeg"
	//"io/ioutil"
	"log"
	"os"
	"reflect"
	"github.com/corona10/goimagehash"
)

func main() {

	directory := "./output_images_madrid"

	
	outputDirRead, err := os.Open(directory)
	if err != nil {
		log.Println(err)
	}
	
	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		log.Println(err)
	}
	filesToAnalyse := []string{}
	
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		
		outputNameHere := outputFileHere.Name()
		filesToAnalyse = append(filesToAnalyse, outputNameHere)
		
	}
	fmt.Print(len(filesToAnalyse))





//	files, err := ioutil.ReadDir(".")
//	if err != nil {
//		log.Fatal(err)
//	}
//	filesToAnalyse := []string{}
//	for _, file := range files {
//		filesToAnalyse = append(filesToAnalyse, file.Name())
		//fmt.Println(file.Name())
//	}


	//	filesToAnalyse := []string{
	//		"TRYP_Madrid_Atocha_Hotel-002106a_hb_a_937.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_a_938.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_a_939.jpg",
	//		"TRYP_Madrid_Atocha_Hotel-002106a_hb_a_940.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_002.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_003.jpg",
	//		"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_005.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_007.jpg", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_008.jpg",
	//		"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_009.jpg",
	//"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_012", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_013",
	//"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_017", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_018", "TRYP_Madrid_Atocha_Hotel-002106a_hb_r_019",
	//"TRYP_Madrid_Atocha_Hotel-002106a_hb_r_020",
	//	}

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
			file2, err := os.Open(ListDebut[IndexCompare])
			//file2, err := base64.StdEncoding.de(file2)
			if err != nil {
				log.Println(err)
			}
			img2, err := jpeg.Decode(file2)
			if err != nil {
				log.Panic(err)
			}
			hash2, err := goimagehash.AverageHash(img2)
			fmt.Print("voici hash2", hash2)
			fmt.Print("voici le type de hash2", reflect.TypeOf(hash2))
			if hash2 == nil {

			}
			if err != nil {
				log.Panic(err)
			}
			HashImageEnQuestion := hash2
			distance, err := ElementToCompare.Distance(HashImageEnQuestion)
			if err != nil {
				log.Println(err)
			}
			if distance >= seuil {
				ListFin = append(ListFin, ListDebut[IndexCompare])
			}
		}
	}
	fmt.Println("///////////// RETURN VALUE =======================================")
	fmt.Println("INDEX CHECK + LISTFIN :: ", IndexCheck, ListFin)
	return ListFin, nil
}
