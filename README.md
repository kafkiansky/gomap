## Generic Map for Go.


[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)

## Contents

- [Installation](#installation)
- [Usage](#usage)
  - [From](#from)
  - [FromSlice](#fromslice)
  - [M](#m)
  - [Filter](#filter)
  - [FilterValues](#filtervalues)
  - [FilterKeys](#filterkeys)
  - [Chunk](#chunk)
  - [Diff](#diff)
  - [Join](#join)
  - [Only](#only)
  - [Each](#each)
- [Testing](#testing)
- [License](#license)


## Installation

```bash
go get -u github.com/kafkiansky/gomap
```

## Usage

### From

`From` creates the `Map` from the given map:

```go
package main

import "github.com/kafkiansky/gomap"

func main() {
  m := gomap.From(map[string]int{}) // Map[string, int]
}
```

### FromSlice

`FromSlice` creates the `Map` from the given slice of values:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
  fmt.Println(gomap.FromSlice([]string{"x", "y"}).Map()) // map[1:x, 2:y]
}
```

### M

Alias to `From`.

### Filter

Allows to filter both keys and values of the given map:

```go
package main

import (
  "github.com/kafkiansky/gomap"
)

func main() {
  m := gomap.M(map[string]int{"x": 1, "y": 2})
  m = m.Filter(func(k string, v int) bool {
	  // do filter logic here
  })
}
```

### FilterValues

Allows to filter just values of the given map:

```go
package main

import (
  "github.com/kafkiansky/gomap"
)

func main() {
  m := gomap.M(map[string]int{"x": 1, "y": 2})
  m = m.FilterValues(func(v int) bool {
	  // do filter logic here
  })
}
```

### FilterKeys

Allows to filter just keys of the given map:

```go
package main

import (
  "github.com/kafkiansky/gomap"
)

func main() {
  m := gomap.M(map[string]int{"x": 1, "y": 2})
  m = m.FilterKeys(func(k string) bool {
	  // do filter logic here
  })
}
```

### Chunk

Allows to chunk given `Map[K, V]` to the slice `[]Map[K, V]`:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
  m := gomap.M(map[string]int{"x": 1, "y": 2, "z": 3})
  maps := m.Chunk(2)
  fmt.Println(len(maps)) // 2

  for _, m := range maps {
	  fmt.Println(m.Map()) // [x: 1, y:2], [z:3]
  }
}
```

### Diff

Return `Map[K, V]` with only this values that does not exist in other maps by keys. 

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
  fmt.Println(
	  gomap.
		  M(map[string]int{"x": 1, "y": 2, "z": 3}).
		  Diff(gomap.M(map[string]int{"y": 2})).
		  Map(),
	  ) // [x:1, z:3]
}
```

### Join

Join maps together.

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
  fmt.Println(
	  gomap.
		  M(map[string]int{"x": 1, "y": 2, "z": 3}).
		  Join(gomap.M(map[string]int{"q": 4})).
		  Map(),
	  ) // [x:1, y:2, z:3, q:4]

  fmt.Println(gomap.Join(
	  gomap.M(map[string]int{"a": 1}),
	  gomap.M(map[string]int{"b": 2}),
	  gomap.M(map[string]int{"c": 3}),
	  ).Map(),
  ) // [a:1, b:2, c:3]
}
```

### Only

Create `Map[K, V]` just for provided keys.

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
  fmt.Println(
	  gomap.
		  M(map[string]int{"x": 1, "y": 2, "z": 3}).
		  Only("x").
		  Map(),
	  ) // [x:1]
}
```

### Each

Apply callback for each element of `Map[K, V]` and return new `Map[K, E]`.

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
	m := gomap.M(map[string]int{"x": 1, "y": 2, "z": 3})

	fmt.Println(gomap.Each(m, func(v int) int64 { return int64(v) * 2 }).Map()) // [x:2, y:4, z:6]
}
```

Or you can apply `Each` function on the `Map` structure. But in this case the mapper function can return only a value of the same type as the map value type, because methods cannot have type parameters:

```go
package main

import (
  "fmt"
  "github.com/kafkiansky/gomap"
)

func main() {
	m := gomap.M(map[string]int{"x": 1, "y": 2, "z": 3})

	fmt.Println(m.Each(func(v int) int { return v * 2 }).Map()) // [x:2, y:4, z:6]
}
```

## Testing

``` bash
$ make check
```  

## License

The MIT License (MIT). See [License File](LICENSE) for more information.