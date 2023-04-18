package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	map1 := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	map1 = removeUnrelated(map1, "address")
	fmt.Println(map1)

	map2 := map[string]interface{}{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	map2 = removeUnrelated(map2, "jump")
	fmt.Println(map2)

	map3 := map[string]interface{}{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	map3 = removeUnrelated(map3, "tiga")
	fmt.Println(map3)
}
