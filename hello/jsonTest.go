package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"
)

func main() {
	sample3()
}
func sample4() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{"Lou", "127.0.0.1"})
	s.Servers = append(s.Servers, Server{"", "127.0.0.2"})
	fmt.Println(s.Servers)
	marshal, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshal))
}
func sample3() {
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents":["Gomz", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		errors.New("errorororororo")
	}
	for x, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case float64:
			fmt.Println("float64===", vv)
		case string:
			fmt.Println("string===", vv)
		case []interface{}:
			fmt.Println("interface===", vv)
			for _, vvv := range vv {
				fmt.Printf("========> %v\n", vvv)
			}
		}

		fmt.Println(x, v)
	}
	fmt.Println(f)
}

func sample1() {
	var s Serverslice
	str := `{"servers":[
		{"ServerName":"ShangHai_VPN", "ServerIp":"127.0.0.1"},
		{"ServerName":"BeiJing_VPN", "ServerIp":"127.0.0.2"}
		]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func sample2() {
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents":["Gomz", "Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int64:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case interface{}:
			fmt.Println(k, "is array:")
			fmt.Println(vv)
			//for i, u := range vv {
			//	fmt.Println(i, u)
			//}
		default:
			fmt.Println(k, "is of ")
		}

	}
	fmt.Println(m)
	var c = []int{1, 2, 3, 4}
	for _, n := range c {
		fmt.Println(n)
	}
}

type Server struct {
	ServerName string `json:"serverName,omitempty"`
	ServerIp   string `json:"serverIp"`
}

type Serverslice struct {
	Servers [] Server `json:"servers"`
}
