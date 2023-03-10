# envreader

Simple environment variable reader for Go.

## Installation

```bash
go get github.com/boatware/envreader
```

## Usage

```go
package main

import (
    "os"

    "github.com/boatware/envreader"
)

func main() {
    // Set environment variables
    os.Setenv("MY_VAR", "my value")
    os.Setenv("MY_INT", "123")
    os.Setenv("MY_BOOL", "true")

    // Declare variables and their types
    var myVar string
    var myInt int
    var myBool bool
	
    // Read environment variables
    myVar = envreader.ReadString("MY_VAR")
    myInt = envreader.ReadInt("MY_INT")
    myBool = envreader.ReadBool("MY_BOOL")
}
```