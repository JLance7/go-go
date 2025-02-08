package maps

import (
	"fmt"
)

func main(){
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Google"] = "google.com"
	fmt.Println(websites)
	websites["NewSite"] = "nice.com"

	delete(websites, "Google")
	fmt.Println(websites)

	/*
	Difference between maps and structs:
		- maps can use any type as a key
		- sturcts have multiple values for a record (id, title, content) (clearly defined data) whereas a map is used as a collection of values with labels
	*/


}