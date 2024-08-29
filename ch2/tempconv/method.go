package main

func main() {

}
func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)
	for i, x := range numbers {
		if j, ok := m[target-x]; ok {
			return []int{j+ 1, i + 1}
		}
		m[x] = i
	}
	return nil
}