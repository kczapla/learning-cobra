package calc

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func convertToFloat64(elems []string) []float64 {
	result := make([]float64, len(elems))
	for i, e := range elems {
		f, err := strconv.ParseFloat(e, 64)
		if err != nil {
			panic(err)
		}
		result[i] = f
	}
	return result
}

func sum(elems []float64) float64 {
	sum := 0.0
	for _, e := range elems {
		sum += e
	}
	return sum
}

func multiply(elems []float64) float64 {
	sum := 1.0
	for _, e := range elems {
		sum *= e
	}
	return sum
}

func subtract(elems []float64) float64 {
	sum := 0.0
	for _, e := range elems {
		sum -= e
	}
	return sum
}

func divide(elems []float64) float64 {
	sum := 0.0
	fmt.Println(elems)
	for _, e := range elems {
		if e == 0.0 {
			panic("div by zero")
		}
		sum /= e
	}
	return sum
}

var (
	CalcCmd = &cobra.Command{
		Use:   "calc",
		Short: "calc ops",
		Long:  "calc ops on pos args",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			elems := convertToFloat64(args)
			if add {
				fmt.Println(sum(elems))
			} else if mul {
				fmt.Println(multiply(elems))
			} else if sub {
				fmt.Println(subtract(elems))
			} else if div {
				fmt.Println(divide(elems))
			}
		},
	}

	mul bool
	div bool
	sub bool
	add bool
)

func init() {
	CalcCmd.Flags().BoolVar(&mul, "mul", false, "multiply")
	CalcCmd.Flags().BoolVar(&div, "div", false, "divide")
	CalcCmd.Flags().BoolVar(&sub, "sub", false, "subtract")
	CalcCmd.Flags().BoolVar(&add, "add", false, "add")
	CalcCmd.MarkFlagsMutuallyExclusive("mul", "div", "sub", "add")
	CalcCmd.MarkFlagsOneRequired("mul", "div", "sub", "add")
}
