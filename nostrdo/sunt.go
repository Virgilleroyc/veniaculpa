s := []int{2, 3, 5, 7, 11, 13}
firstValue := s[0] // Get the first element of the slice
s = s[1:] // Remove the first element from the slice

fmt.Println(firstValue) // Print the removed element
fmt.Println(s) // Print the updated slice
