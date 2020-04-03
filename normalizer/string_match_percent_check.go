package main

import (
	"GIG/commons"
	"fmt"
)

func main()  {
fmt.Println(libraries.StringMatchPercentage("minister for rural affairs","minister for urban affairs"))
fmt.Println(libraries.StringMatchPercentage("securities exchang commission","securities exchange commission"))
fmt.Println(libraries.StringMatchPercentage("department police","department excise"))
fmt.Println(libraries.StringMatchPercentage("department hindu religious cultural affairs","department muslim religious cultural affairs"))
fmt.Println(libraries.StringMatchPercentage("sri lanka national freedom from hunger campaign board","sri lanka national freedom hunger campaign board"))
}
