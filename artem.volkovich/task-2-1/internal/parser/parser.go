package parser

import "fmt"

func Parse() ([]int, error) {
	var N, K int
	_, err := fmt.Scan(&N)
	if (err != nil) {
		return nil, fmt.Errorf("Offices input error: %w", err)
	}
	var output []int
	for i:=0;i<N;i++ {
		min := 15
		max := 30
		_, err = fmt.Scan(&K)
		if (err != nil) {
			return nil, fmt.Errorf("Employee input error: %w", err)
		}
		for j:=0;j<K;j++ {
			var operator string
			var operand int
			_, err = fmt.Scan(&operator, &operand)
			if (err != nil) {
				return nil, fmt.Errorf("Temperature input error: %w", err)
			}
			if operator != ">=" && operator != "<=" {
				return nil, fmt.Errorf("Operator input error")
			}
			switch operator {
			case ">=":
				if operand > min {
					min = operand
				}
			case "<=":
				if operand < max {
					max = operand
				}
		}
			if min > max {
				output = append(output, -1)
			} else {
				output = append(output, min)
			}
		}
	}
	
	
	return output, nil
}