package trivial

import "fmt"

func DataStructures() {
	Slice()
	Map()
}

func Slice() {
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users := initialUsers[2:7]
	fmt.Println(users)
	users = append(users[:1], users[3:]...)
	fmt.Println(users)
}

func Map() {
	var people = map[string]int{
		"Bob":   1,
		"Alice": 2,
		"Kate":  3,
		"Sam":   4,
	}
	fmt.Println(people["Alice"])
	fmt.Println(people["Bob"])
	people["Bob"] = 5
	fmt.Println(people["Bob"])

	if val, ok := people["Kate"]; ok {
		fmt.Println(val)
	}

	for key, value := range people {
		fmt.Println(key, value)
	}
	delete(people, "Bob")
	fmt.Println(people)
}
