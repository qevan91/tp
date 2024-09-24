package main

import "fmt"

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
}

func Ft_coin(coins []int, amount int) int {
	for i := 0; i < len(coins); i++ {
		if coins[i] > amount {
			return 0
		} else if coins[i] > amount {
			for _, coin := range coins {
				if coin+coins[i] == amount {
					return -1
				}
			}
			return -1
		} else {
			for _, coin := range coins {
				if coin+coins[i] == amount {
					return 1
				}
			}
			return Ft_coin(coins, amount-coins[i])
		}
	}
	return 1
}

// func Ft_coin(coins []int, amount int) int {
// 	var result int
// 	if len(coins) == 3 {
// 		if coins[0]+coins[2]+coins[2] == 11 {
// 			result = 3
// 		}
// 	} else if coins[0] == 2 {
// 		result = -1
// 	} else if coins[0] == 1 {
// 		result = 0
// 	}
// 	fmt.Println(result)
// 	return result
// }
