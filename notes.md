
# Compiling and running

```zsh
go build <filename>.go

./<filename>
```

# Just Running go files

```zsh
go run <filename>.go
```

# Setting up Go files

Start by adding

```go
package main
```

Add "fmt" package for printing

```go
import "fmt"
```

Go to [fmt-docs](www.github.com/evandabest/go-notes/fmt)

# The main function

In the file you will make a function

Keyword: func
 - declares a function

```go
func main() {}
```

# Importing time and aliases

```go
import (
    t "time"
)
```

Characters before the package name are used to represent that package in a different name

Now I can reference the time package with just "t"

```go
fmt.Println(t.Now())
```

# Data type

Variables can be many different types of data

```go
var cap bool = true 
var num int8 = -127
var pnum uint8 = 127
var nums int16 = -32768
var pnums uint16 = 32768
//more for the int storage

var deci float = 23.242
```



# Declaring Variables

You can use the following to declare variables

Const

```go
const num = 2;
```

Var
```go 
var num uint16 = 2;
```









