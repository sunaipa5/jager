
# Jager
This library was created to generate `JSON` in various and fast ways and send it via `http.responseWrite.`

[![Go Reference](https://pkg.go.dev/badge/github.com/sunaipa5/jager.svg)](https://pkg.go.dev/github.com/sunaipa5/jager)
## Examples

### String To JSON
```go

package main

import (
	"net/http"
    "github.com/sunaipa5/jager"
)

func main(){

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	jager.String(w, `{"type":"jager string","name":"john","surname":"doe","age":30, "isStudent":true}`)
    })

    http.ListenAndServe(":3000", nil)

}
```

### Map To JSON
```go
package main

import (
	"net/http"
    "github.com/sunaipa5/jager"
)

func main(){

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    	jager.Map(w, map[string]interface{}{
    		"name":      "John",
    		"surname":   "Doe",
    		"age":       30,
    		"isStudent": true,
    	})
    })
    
    http.ListenAndServe(":3000", nil)

}
```

### Struct To JSON
```go
package main

import (
	"net/http"
    "github.com/sunaipa5/jager"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        person := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "johndoe@example.com",
        }

	   jager.Struct(w,person)
    })

    http.ListenAndServe(":3000", nil)

}
```

