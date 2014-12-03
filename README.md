## Go FileSystem

fs shortcuts in Golang.

### Installation
```bash
$ go get github.com/turingou/go-fs
```

### Example
```go
package main

import "fmt"
import "github.com/turingou/go-fs"

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  // Create a tmp file
  err := fs.WriteFile("./demo.txt", "demo text")
  check(err)

  // Read from tmp file we just created
  fileString, err := fs.ReadFile("./demo.txt")
  check(err)
  
  // => `demo text`
  fmt.Println(fileString) 
  
  data := make(map[string]interface{})
  data["name"] = "John"
  data["age"] = 12

  // Create JSON file from a map
  err := fs.WriteJSON("./demo.json", data)
  check(err)

  // Read JSON file from tmp json file
  json, err := fs.ReadJSON("./demo.json")

  fmt.Println(json)
}
```

### Tests

```bash
$ go test
```

### Contributing
- Fork this repo
- Clone your repo
- Install dependencies
- Checkout a feature branch
- Feel free to add your features
- Make sure your features are fully tested
- Open a pull request, and enjoy <3

### MIT license
Copyright (c) 2014 turing &lt;o.u.turing@gmail.com&gt;

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the &quot;Software&quot;), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED &quot;AS IS&quot;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

---
![docor](https://raw.githubusercontent.com/turingou/docor/master/docor.png)
built upon love by [docor](https://github.com/turingou/docor.git) v0.2.0
