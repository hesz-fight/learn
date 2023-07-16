package test

type PlaceHolder struct{}

func main() {
	map1 := make(map[int]PlaceHolder)
	map1[0] = PlaceHolder{}
}
