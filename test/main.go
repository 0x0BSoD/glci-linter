package main

import (
	"errors"
	"fmt"
	"github.com/0x0BSoD/glci-linter/pkg/gitlab"
	"gopkg.in/yaml.v3"
	"os"
)

// Custom Resolvers ========

func resolveReference(node *yaml.Node) (*yaml.Node, error) {
	if node.Kind != yaml.SequenceNode {
		return nil, errors.New("include on a non-sequince node")
	}

	source := node.Content[0].Value
	part := node.Content[1].Value
	varName := ""
	if len(node.Content) == 3 {
		varName = node.Content[2].Value
	}

	fmt.Println(source)
	fmt.Println(part)
	fmt.Println(varName)

	// file, err := os.ReadFile(node.Value)
	// if err != nil {
	// 	return nil, err
	// }
	// var f Fragment
	// err = yaml.Unmarshal(file, &f)
	// return f.content, err
	return node, nil
}

func main() {
	gitlab.AddResolvers("!reference", resolveReference)

	f, err := os.ReadFile("root.yaml")
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(f, &gitlab.CustomTagProcessor{Target: &data})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
