package main
package main

import "fmt"

/*
 * 这是从控制台读取一行的帮助函数
 */
func readLine() (string ,error) {
	var (
		cc []byte
		err error
	)
	for {
		var c byte
		n, err := fmt.Scanf("%c", &c)
		if err == nil && n == 1 && c != 10 {
			cc = append(cc, c)
		} else {
			break
		}
	}
	return string(cc), err
}


/* Write Code Here */
func ThreeDigitNumbers (selectedDigits_size int, selectedDigits []int32) []int32 {


}
func main() {
	var res []int32

	selectedDigits_size := 0
	fmt.Scanf("%d", &selectedDigits_size)
	var selectedDigits []int32
	for selectedDigits_i := 0; selectedDigits_i < selectedDigits_size; selectedDigits_i++ {
		var selectedDigits_item int32
		fmt.Scanf("%d", &selectedDigits_item)

		selectedDigits = append(selectedDigits, selectedDigits_item)
	}


	res = ThreeDigitNumbers (selectedDigits_size, selectedDigits)
	res_size := len(res)
	for res_i := 0; res_i < res_size; res_i++ {

		fmt.Printf("%d\n", res[res_i])

	}

}
