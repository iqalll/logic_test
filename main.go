package main

// importing fmt

// Golang program to reverse a string
// func reverse(str string) (result string) {
// 	for _, v := range str {
// 		result = string(v) + result
// 	}
// 	return result
// }

// func main() {
// 	str := "I am A Great human"
// 	strRev := reverse(str)
// 	fmt.Println(str)
// 	fmt.Println(strRev)
// }

// Golang program to reverse a int
// func reverseInt(input int) (result int) {
// 	strInput := strconv.Itoa(input)
// 	resultStr := strconv.Itoa(result)
// 	for _, v := range strInput {
// 		resultStr = string(v) + resultStr
// 		result, _ = strconv.Atoi(resultStr)
// 	}
// 	return
// }
// func main() {
// 	output := reverseInt(567890)
// 	fmt.Println(output)
// }

// Golang program to Find Middle Alphabet
// func findMiddleAlphabet(a string, b string, alphabet string) (huruf1 string, huruf2 string) {
// 	var (
// 		awal, akhir int
// 	)
// 	arr := strings.Split(alphabet, "")
// 	for i := 0; i < len(arr); i++ {
// 		if a == arr[i] {
// 			awal = i
// 		}
// 		if b == arr[i] {
// 			akhir = i
// 		}
// 	}
// 	jarak := akhir - awal + 1
// 	if jarak%2 == 0 {
// 		middle := jarak / 2
// 		awal = awal + middle - 1
// 		akhir = akhir - middle + 1
// 		for i := 0; i < len(arr); i++ {
// 			if awal == i {
// 				huruf1 = arr[i]
// 			}
// 			if akhir == i {
// 				huruf2 = arr[i]
// 			}
// 		}
// 	}
// 	if jarak%2 != 0 {
// 		middle := jarak/2 + awal
// 		for i := 0; i < len(arr); i++ {
// 			if middle == i {
// 				huruf1 = arr[i]
// 			}
// 		}
// 	}
// 	return huruf1, huruf2
// }

// func main() {
// 	output1, output2 := findMiddleAlphabet("V", "Z", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
// 	fmt.Println(output1, output2)
// }

// Golang program to Fibonacci
// func main() {
// 	nums := []int{15, 1, 3}

// 	fmt.Println(nearestFibonacci(nums))
// }

// func nearestFibonacci(nums []int) int {
// 	var num int
// 	for _, val := range nums {
// 		num += val
// 	}

// 	if num == 0 {
// 		return 0
// 	}

// 	first := 0
// 	second := 1
// 	third := first + second
// 	for third <= num {
// 		first, second = second, third
// 		third = first + second
// 	}

// 	subtraction := func(num1, num2 int) int {
// 		if num1 > num2 {
// 			return num1 - num2
// 		}
// 		return num2 - num1
// 	}

// 	if math.Abs(float64(third-num)) >= math.Abs(float64(second-num)) {
// 		return subtraction(num, second)
// 	}

// 	return subtraction(num, third)
// }

// Golang program to  Multiple 3 and 5
// func main() {
// 	number := 25

// 	fmt.Println(solution(number))
// }

// func solution(number int) string {
// 	var a int
// 	var arr []int
// 	for i := 1; i < number; i++ {
// 		if i%3 == 0 || i%5 == 0 {
// 			a += i
// 			arr = append(arr, i)
// 		}
// 	}

// 	ets := strings.Trim(strings.Replace(fmt.Sprint(arr), " ", " + ", -1), "[]")

// 	res := fmt.Sprintf("%d = %s", a, ets)

// 	return res
// }

// Golang program to  Alternate Case
// func main() {
// 	fmt.Println(alternateCase("Hello World"))
// }

// func alternateCase(s string) string {
// 	arrS := strings.Split(s, "")

// 	var abj []string
// 	for _, v := range arrS {
// 		if v == strings.ToLower(v) {
// 			abj = append(abj, strings.ToUpper(v))
// 			continue
// 		}

// 		abj = append(abj, strings.ToLower(v))
// 	}

// 	result := strings.Join(abj, "")

// 	return result
// }

// Golang program to Map Array Data
// func main() {
// 	nums := []int{16, 17, 4, 3, 5, 2}

// 	fmt.Println(productArray(nums))
// }

// func productArray(nums []int) []int {
// 	var result []int

// 	prod := 1
// 	for _, val := range nums {
// 		prod *= val
// 	}

// 	for _, val := range nums {
// 		result = append(result, int(float64(prod)*math.Pow(float64(val), -1)))
// 	}

// 	return result
// }

// Golang program to Biggest Number
// func main() {
// 	res := maxRedigit(1234)
// 	fmt.Println(*res)
// }

// func maxRedigit(data int) *int {
// 	if data <= 0 {
// 		return nil
// 	}

// 	strData := strconv.Itoa(data)
// 	arrData := strings.Split(strData, "")

// 	a := 0
// 	sort.Slice(arrData, func(i, j int) bool {
// 		if arrData[i] > arrData[j] {
// 			a++
// 		}
// 		return arrData[i] > arrData[j]
// 	})

// 	if a < 1 {
// 		return nil
// 	}

// 	dataStr := strings.Join(arrData, "")
// 	res, _ := strconv.Atoi(dataStr)

// 	return &res
// }
