package main

import (
	"bufio"
	"fmt"
	"os"
)


type question struct {
	soal string
	option []string
	jawab int
}

func tampilkanSoal(q question) {
	fmt.Println(q.soal)
	for i, option := range q.option {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func cekJawaban(jawab int, Benar int) bool {
	return jawab == Benar
}


func main() {

	scn := bufio.NewScanner(os.Stdin)
	fmt.Print("Username :")

	scn.Scan()
	nama := scn.Text()

	questions := []question {
		{
			soal: "Siapakah manusia pertama ?",
			option: []string{"Widya & Abas", "Adam & Hawa", "Pandora", "Kobo"},
			jawab : 2,
		},
		{
			soal: "Berapakah 1+1 =...",
			option: []string{"2", "10", "Jendela", "error"},
			jawab : 2,
		},
		{
			soal: "Siapakah Kaisar yang telah menakhlukan kerajann di Eropa, Asia, India ?",
			option: []string{"Pedo", "Qin Shi Huang", "Alexander", "Kang Xi"},
			jawab : 3,
		},
		{
			soal: "Siapakah Presiden ke - 6",
			option: []string{"Susilo Bambang Yudhoyono", "Naruto Uzumaki", "Cecep Sutejo", "Ebil Korsa"},
			jawab : 1,
		},
		{
			soal: "Nama lain inti sel adalah...",
			option: []string{"Nukleus", "Ribosom", "Klorofil", "Mitokondria"},
			jawab : 1,
		},
		
	} 

	Benar := 0
	Salah := 0

	for _, question := range questions {
		tampilkanSoal(question)

		var jawab int
		fmt.Print("Jawab :")
		fmt.Scan(&jawab)
		fmt.Printf("\n")
		if cekJawaban(jawab, question.jawab){
			fmt.Print("\n")
			Benar++
		} else{
			fmt.Print("benar,\n")
			Salah++
		}
   
	}

	score := Benar * 2;
	var ket string


	fmt.Printf("Nama : %s \n", nama)
	fmt.Printf("Benar = %d \n", Benar)
	fmt.Printf("Salah = %d \n", Salah)
	fmt.Printf("Score = %d \n", score)
	fmt.Println(ket)


}