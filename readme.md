# HCL Parser

## Overview

This package is used to generate Javascript parser for HCL (HashiCorp Configuration Language).

Thank you [phamtai97](https://github.com/phamtai97) for making this feasible

## Installation

```sh
# Install GopherJS first
$ go get -u github.com/gopherjs/gopherjs

# Build JS package
$ gopherjs build github.com/anhldbk/hcl-parser -o index.js -m

# Require & enjoy
```

## Example JS code

```js
import {parse} from "hcl-parser"
const input = `
# only root or owners to manage groups
allow {
  subject = "iam:user/anhldbk"
  action = "group/\*"
  resource = [
    "iam:group/admin"
  ]
}

# allow any other subjects to read
allow {
  subject = "\*"
  action = [
    "group/read",
    "group/list"
  ]
  resource = [
    "iam:group/zalopay/pma"
  ]
}  
`

console.log(
  JSON.stringify(
    parse(input)
  )
)
```