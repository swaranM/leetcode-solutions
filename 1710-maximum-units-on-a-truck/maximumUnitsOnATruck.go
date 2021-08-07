package maximumunitsonatruck

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(a, b int) bool { return boxTypes[a][1] > boxTypes[b][1] })
	ans := 0
	for _, box := range boxTypes {
		if box[0] <= truckSize {
			ans += box[0] * box[1]
			truckSize -= box[0]
		} else {
			ans += box[1] * truckSize
			break
		}
	}
	return ans
}
