package hoosh

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/hooss-only/stack"
)

func Encode(text string) string {
	texts := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890`~-_=+]}[{\\|'\"?.>,</;:!@#$%^&*() ", "")
	numbers := stack.Stack{}
	hooshed := 0
	for i := range texts {
		rand.Seed(int64(i))
		numbers.Push(strconv.Itoa(rand.Int()))
	}
	for index, i := range texts {
		if strings.Contains(text, i) {
			f, err := strconv.Atoi(numbers.GetStack()[index])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			hooshed += f
		}
	}
	hooshNumber := strings.Split("1234567890", "")
	hooshString := strings.Split("@!&$#@=/+'", "")
	result := strconv.Itoa(hooshed)
	for index, i := range hooshNumber {
		result = strings.ReplaceAll(result, i, hooshString[index])
	}
	return result
}
