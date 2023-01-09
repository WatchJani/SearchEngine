package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"root/model/squad"
)

func main() {
	squad := squad.NewEmpty()
	//otvorimo file
	file, err := os.Open("./data.json")

	if err != nil {
		log.Fatal(err)
	}

	//procitamo file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	//ubacimo sve iz fila u strukturu
	err = json.Unmarshal(data, squad)
	if err != nil {
		log.Fatal(err)
	}

	//igramo se sa strukturom
	//fmt.Println(squad.Members[0].Age)

	fmt.Println(squad.FindByName("M").Age)

}
