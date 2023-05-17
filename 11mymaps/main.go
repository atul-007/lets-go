package main

import "fmt"

func main() {

	fmt.Println("maps")
	lang := make(map[string]string) //map[keytype]value type
	lang["go"] = "golang"
	lang["py"] = "python"
	lang["c"] = "cpp"
	lang["ry"] = "ruby"
	fmt.Println(lang)
	fmt.Println("py shots for", lang["py"])
	delete(lang, "ry") //can be used for slices also
	fmt.Println(lang)

}
