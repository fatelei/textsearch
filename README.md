# textsearch
Text search algorithm implement in golang

## Implemented algorithm.

- [Kmp](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm)
- [Rabin karp](https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm)

## Usage

### Kmp

```
package main

import (
  "fmt"
  "github.com/fatelei/textsearch/kmp"
)

func main() {
  fmt.Printf("%v\n", kmp.Search("abcd", "abc"))
}
```

### Rabin karp

```
package main

import (
  "fmt"
  "github.com/fatelei/textsearch/rabin_karp"
)

func main() {
  // You should set base and module.
  rk = RabinKarp{base: 101, module: 1 << 64}
  fmt.Printf("%v\n", rk.search("abcd", "bc"))
}
```