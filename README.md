# obscure

package obscure provides a command that will obscure any text passed in via standard input

obscure replaces all the text passed into standard input, save the first five characters,
with random symbols. It was developed as a way to show sensitive information (auth tokens,
etc.) to the audience during presentations. It lets you prove, e.g., that an environment
variable is set without disclosing the full contents of the variable.

## Example usage:

```sh
$ pwd
/home/paddy/go/src/github.com/paddyforan/obscure
$ pwd | obscure
/home*&%&&@$$&#!#&?&??@!***$*!&&@@?$?&?$@*#&&?#**
```
