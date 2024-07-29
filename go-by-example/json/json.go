package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type response struct {
	Page   int `json:"my_page"`
	Fruits []string
}

func main() {

	r1 := response{
		Page:   1,
		Fruits: []string{"apple", "banana", "pear", "watermelon"},
	}

	rlt1, _ := json.Marshal(r1)
	fmt.Println(string(rlt1))
	r2 := response{}

	json.Unmarshal(rlt1, &r2)
	fmt.Println(r2)

	log.Println("-----------------")
	bolb, _ := json.Marshal(true)
	log.Println(string(bolb))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	log.Println("++++++++++++++++++++")
	byt := []byte(`{"go":1,"num":2.13,"strs":["abc","def"]}`)
	var dat map[string]interface{}

	json.Unmarshal(byt, &dat)
	log.Println(dat)

	num, ok := dat["num"].(float64)
	if ok {
		fmt.Printf("%T %f", num, num)
	} else {
		fmt.Printf("not ok ")
	}

	strs := dat["strs"].([]interface{})
	tmp := strs[0].(string)
	log.Println(tmp)

	log.Println(json.Valid(byt))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{
		"go":    1111,
		"apple": 100,
	}
	enc.Encode(d)

}
