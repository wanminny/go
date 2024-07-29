package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("name = %v, id =%v, origin = %v", p.Name, p.Id, p.Origin)
}

func main() {

	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	//fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	//json.MarshalIndent()
	var p Plant
	xml.Unmarshal(out, &p)
	log.Println(p)
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"go>child>plant"`
	}
	nesting := &Nesting{}

	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))

}
