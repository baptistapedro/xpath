package myfuzz
import "github.com/antchfx/xpath"

func Fuzz(data []byte) int  {
	_, err := xpath.Compile(string(data))
	if err != nil {
		return 1
	}
	return 0
}
