package helpers

import "strings"

func GenerateDataStructures(dataArray [][]string, fileName string) (map[string][]string, map[string]string) {
	var dataStructure = make(map[string][]string)
	var nameStructure = make(map[string]string)

	for _, record := range dataArray {
		record0, record1 := strings.TrimSpace(record[0]), strings.TrimSpace(record[1])

		if record0 == "Terminate" {
			TerminateEntities(fileName, record1)
		} else {
			ministryName := record0
			childName := record1

			//if ministry name change is detected, create a change name request
			ministryNameArray := strings.Split(ministryName, "->")
			if len(ministryNameArray) == 2 {
				ministryName = strings.TrimSpace(ministryNameArray[1])
				nameStructure[ministryName] = strings.TrimSpace(ministryNameArray[0])
			}

			dataStructure[ministryName] = append(dataStructure[ministryName], childName)
		}
	}

	return dataStructure, nameStructure
}
