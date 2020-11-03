package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Datas []Data
type Data struct {
	Id           string `json:"id"`
	ParentId     string `json:"parent_id"`
	BusinessName string `json:"business_name"`
}

type r struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Children []r    `json:"children"`
}

func main() {
	data, err := ioutil.ReadFile("./tree.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var business Datas
	datamap := make(map[string]Datas, 0)
	err = json.Unmarshal(data, &business)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(len(business))
	for _, v := range business {
		datamap[v.ParentId] = append(datamap[v.ParentId], v)
	}
	result := buildTree(datamap, Data{Id: "0", BusinessName: "root"})
	js, _ := json.Marshal(result)
	ioutil.WriteFile("./tmp_tree.json", js, 0777)
	fmt.Println("end")
}

func buildTree(d map[string]Datas, dd Data) r {
	tmp, ok := d[dd.Id]
	var tree r
	tree.ID = dd.Id
	tree.Name = dd.BusinessName
	if !ok {
		return tree
	}
	for _, v := range tmp {
		tree.Children = append(tree.Children, buildTree(d, v))
	}

	return tree
}
