package main

import (
	"fmt"
	"mychain/utils"
)

func main () {
	fmt.Println("Let1's open a new world...")
	blockchain := utils.GenesisBlock("Hello Please call me No.1")
	utils.GenerateBlock(&blockchain, "send 999999999999$ to yuyang, ha ha ha.....")
	utils.Print(blockchain)	
}
