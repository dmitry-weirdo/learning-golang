package main

import "fmt"

func main() {

	i := 10

	if i < 15 {
		goto myLabel // works -> jump out of the block into the containing block
	} else if i < 100 {
		goto labelAfterVariableDeclaration // does NOT work -> after a variable declaration
	} else {
		goto labelInOtherBlock // does NOT work -> label in another block is not found
	}

myLabel: // out of block, within the containing block
	j := 42
labelAfterVariableDeclaration:

	for ; i < 15; i++ {
	labelInOtherBlock:
		fmt.Println(i)
	}
}
