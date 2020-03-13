package struct_deep_compare

import (
	"reflect"
)

type Profile struct {
	Age  int
	Name string
}

type Post struct {
	Title string
}

type Store struct {
	ID   int
	Name string
}

type User struct {
	ID      int
	Account *string
	Profile Profile
	Store   *Store
	Posts   []Post
}

func StructDeepCompare(obj1, obj2 interface{}) bool {
	return reflect.DeepEqual(obj1, obj2)
}
