package problems

import "fmt"

// Sample is just so we can keep the problems package in main.go
func SampleFunction() {
	fmt.Println("Hello World")
}

func modifyMap(m map[string]int) {
	m["hello"] = 2
	fmt.Println("Call Inside modifyMapFunction", m)
}

func TestMapOutput() {
	m := map[string]int{}

	m["hello"] = 1
	m["world"] = 2
	fmt.Println("Map before modifyMap Function: ", m)
	modifyMap(m)

	fmt.Println("Map after modifyMap Function: ", m)
}

func appendSlice(s *[]int) {
	*s = append(*s, 5)
	fmt.Println("Slice inside appendSlice Function: ", *s)
}

func TestSliceOutput() {
	s := &[]int{}
	*s = append(*s, 1)

	fmt.Println("Slice before appendSlice Function: ", *s)
	appendSlice(s)
	fmt.Println("Slice after appendSlice Function: ", *s)
}

func sliceSlice(s []int) {
	s = s[1:]

	fmt.Println("Slice inside sliceSlice Function: ", s)
}

func TestSlicedSliceOutput() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice before appendSlice Function: ", s)
	sliceSlice(s)
	fmt.Println("Slice after appendSlice Function: ", s)

}
