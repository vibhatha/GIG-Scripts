package main

import (
	"GIG/commons"
	"fmt"
)

func main()  {
fmt.Println(commons.StringMatchPercentage("minister for rural affairs","minister for urban affairs"))
fmt.Println(commons.StringMatchPercentage("securities exchang commission","securities exchange commission"))
fmt.Println(commons.StringMatchPercentage("department police","department excise"))
fmt.Println(commons.StringMatchPercentage("department hindu religious cultural affairs","department muslim religious cultural affairs"))
fmt.Println(commons.StringMatchPercentage("sri lanka national freedom from hunger campaign board","sri lanka national freedom hunger campaign board"))
}
