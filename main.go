// package obscure provides a command that will obscure any text passed in via standard input
//
// obscure replaces all the text passed into standard input, save the first five characters,
// with random symbols. It was developed as a way to show sensitive information (auth tokens,
// etc.) to the audience during presentations. It lets you prove, e.g., that an environment
// variable is set without disclosing the full contents of the variable.
//
// Example usage: $ echo $MYVAR | obscure
package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const ranChars = `$@!*#%&?`

func main() {
	rand.Seed(time.Now().Unix())
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Error reading standard in: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s%s\n", in[:5], randomChars(len(in)-5))
}

func randomChars(n int) string {
	var chars []byte
	for i := 0; i < n; i++ {
		pos := rand.Intn(len(ranChars))
		chars = append(chars, ranChars[pos])
	}
	return string(chars)
}
