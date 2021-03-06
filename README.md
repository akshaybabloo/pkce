# Proof Key for Code Exchange (PKCE)

[![Go](https://github.com/akshaybabloo/pkce/actions/workflows/go.yml/badge.svg)](https://github.com/akshaybabloo/pkce/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/akshaybabloo/pkce?status.svg)](https://pkg.go.dev/github.com/akshaybabloo/pkce)
[![codecov](https://codecov.io/gh/akshaybabloo/pkce/branch/main/graph/badge.svg?token=71W6D8Z5CX)](https://codecov.io/gh/akshaybabloo/pkce)

Generate PKCE verify and challenge code.

## Usage

There are two ways of doing this

### Your own string

You can put in your own string, example:

```go
package main

import (
	"fmt"

	"github.com/akshaybabloo/pkce"
)

func main() {
	p := pkce.Pkce{
		RandomString: "45d7820e694481f399e7fb9c444f0cb63301a7254d1401443835d9af2c9a6a5ec5b243c3470feb945336025964ef05c8d2f0e44baf76762ba6136914",
	}

	fmt.Println(p.VerifyCode())  // This is optional
	// Output: 45d7820e694481f399e7fb9c444f0cb63301a7254d1401443835d9af2c9a6a5ec5b243c3470feb945336025964ef05c8d2f0e44baf76762ba6136914
	
	fmt.Println(p.ChallengeCode())
	// Output: iQoF8w9kq5RnuMdisRXypyOoMCF7FGz-ro7dwHjC28U
}
```

When you enter your own string, calling `p.VerifyCode()` is optional

### Let us generate one for you

You can let the module generate a random string for you, but the length should be more than or equal to `43` and lest and or equal to `128`.

```go
package main

import (
	"fmt"

	"github.com/akshaybabloo/pkce"
)

func main() {
	p := pkce.Pkce{
		Length: 128,
	}

	fmt.Println(p.VerifyCode())
	// Output: 45d7820e694481f399e7fb9c444f0cb63301a7254d1401443835d9af2c9a6a5ec5b243c3470feb945336025964ef05c8d2f0e44baf76762ba6136914
	
	fmt.Println(p.ChallengeCode())
	// Output: iQoF8w9kq5RnuMdisRXypyOoMCF7FGz-ro7dwHjC28U
}
```

> Note: You have to call `p.VerifyCode()` before `p.ChallengeCode()` to generate a random string and its hash
