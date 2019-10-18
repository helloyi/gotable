# The Table in go

[![GoDoc](https://godoc.org/github.com/helloyi/gotable?status.svg)](https://godoc.org/github.com/helloyi/gotable)

Put 'interface{}' on the Table, and to unfold, then, manipulates interface{} value, convenient and simple.

## Usage

Pick up some value from interface{}.
```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/helloyi/gotable"
)

func main() {
	resp, err := http.Get("https://api.alternative.me/fng/?limit=10")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var res interface{}
	if err := json.Unmarshal(data, &res); err != nil {
		log.Fatalln(err)
	}

	rest := table.New(res)

	fngs := map[string]string{}
	_ = rest.MustGet("data").EachDo(func(_, fng *table.Table) error {
		ts := fng.MustGet("timestamp").String()
		val := fng.MustGet("value").String()
		fngs[ts] = val
		return nil
	})

	data, err = json.MarshalIndent(fngs, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(data))
}
```

For a map:
```go
package main

import (
	"github.com/helloyi/gotable"
	"log"
)

func main() {
	var m interface {} = map[int]int{
		1: 1,
		2: 2,
	}
	t := table.New(m)

	sum := 0
	if err := t.EachDo(func(k, v *table.Table) error {
		sum += v.MustInt()
		return nil
	}); err != nil {
		log.Fatalln(err)
	}

	log.Println(sum) // print 3
}
```

Convert to struct
```go
package main

import (
	"fmt"
	"log"

	"github.com/helloyi/gotable"
)


func main() {
	x := map[string]interface{}{
		"a": 1,
		"A": 11,

		"B": "b",
		"b": "bb",

		"C": "c",
		"c": "cc",
	}

	var y struct {
		A int    `table:"a"` // set with name "a"
		B string `table:"_"` // passed
		C string // set with name "C"
	}

	t := table.New(x)
	if err := t.Conv(&y); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", y)
}
```

Set map key:
```go
package main

import (
	"github.com/helloyi/gotable"
	"log"
)

func main() {
  var m interface {} = map[int]int{
		1: 1,
		2: 2,
	}
	t := table.New(m)

	if err := t.Put(1, 100); err != nil {
		log.Fatalln(err)
	}

	log.Println(t.MustGet(1).MustInt()) // print 100
}
```
