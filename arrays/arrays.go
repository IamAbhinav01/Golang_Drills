package main

import "fmt"


func arrays_demo(){
	var arr1 [] int //empty array
	arr1 = append(arr1, 1,2,3,4,5)

	fmt.Println("Array 1 : ",arr1)
}

func main(){
	arrays_demo()
}

// | Function   | Purpose                              | Example             |
// | ---------- | ------------------------------------ | ------------------- |
// | `len()`    | Returns length                       | `len(arr)`          |
// | `cap()`    | Returns capacity (mainly for slices) | `cap(slice)`        |
// | `copy()`   | Copies elements                      | `copy(dst, src)`    |
// | `append()` | Adds elements to a slice             | `append(slice, 10)` |
// | Function            | Purpose                 |
// | ------------------- | ----------------------- |
// | `slices.Contains()` | Check if element exists |
// | `slices.Index()`    | Find index of element   |
// | `slices.Sort()`     | Sort ascending          |
// | `slices.SortFunc()` | Custom sort             |
// | `slices.Reverse()`  | Reverse slice           |
// | `slices.Equal()`    | Compare slices          |
// | `slices.Clone()`    | Copy slice              |
// | `slices.Delete()`   | Delete elements         |
// | `slices.Insert()`   | Insert elements         |
// | `slices.Replace()`  | Replace range           |
// | Function            | Purpose           |
// | ------------------- | ----------------- |
// | `sort.Ints()`       | Sort int slice    |
// | `sort.Strings()`    | Sort string slice |
// | `sort.Float64s()`   | Sort float slice  |
// | `sort.SearchInts()` | Binary search     |
// | Operation           | Code                                   |
// | ------------------- | -------------------------------------- |
// | Length              | `len(arr)`                             |
// | Access              | `arr[i]`                               |
// | Update              | `arr[i] = value`                       |
// | Traverse            | `for i, v := range arr`                |
// | Copy                | `copy(dst, src)`                       |
// | Sort                | `slices.Sort(arr)` or `sort.Ints(arr)` |
// | Search              | `slices.Index(arr, x)`                 |
// | Reverse             | `slices.Reverse(arr)`                  |
// | Append (slice only) | `append(slice, x)`                     |
