# Seely Golang Client

Golang client of Seely platform.

# Usage

## Install

```
go get -u github.com/teletraan/go-seely
```

## Example

```
package main

import (
    seely github.com/teletraan/go-seely
)

func main() {
    c := seely.New("http://localhost:2021/graphql/", "a@t.io", "password")
    c.MemberService.Search
}
```