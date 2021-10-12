# purepkg - pure packages

When developing in Go it can be desirable at times to limit the imports in a package. For reasons like:
* Avoid the likelihood of future circular dependencies when you have a certain package structure in mind.
* Having the intention to move the package to its own repository later as an independent library.
* Enforcement of a certain package architecture.
* Prevent accidently importing the wrong package: think of `math/rand` instead of `crypto/rand` in cryptographic code.
* Caring about the purity of a package, eg. only using the standard library.

`purepkg` can help to achieve that by verifying that only allowed packages are imported.

## Installation

You can install `purepkg` directly from the `main` branch with `go install github.com/the/purepkg`. Or run it directly by invoking `go run github.com/the/purepkg`.

## Usage

```
purepkg [-v] [-stdlib] [-notest] [-allow pkg1,pkg2,...] <package>

Options:
  -allow packages
    	comma separated list of allowed packages
  -notest
    	ignore tests
  -stdlib
    	include all standard library packages
  -v    verbose output
```

When no unallowed package is found `purepkg` exits with exit code `0`. Otherwise it will exit with `2` which can be useful to abort a build process or checks executed in a CI/CD pipeline. By default test code is included. Adding the `-notest` argument ignores the tests.

## Examples

```
$ purepkg -stdlib github.com/the/purepkg

package github.com/the/purepkg
  ✗ import golang.org/x/tools/go/packages
```

```
$ purepkg -v -stdlib -allow "golang.org/*" github.com/the/purepkg

package github.com/the/purepkg
  ✔ import log
  ✔ import os
  ✔ import strings
  ✔ import flag
  ✔ import fmt
  ✔ import golang.org/x/tools/go/packages
```

```
$ purepkg -v -stdlib -allow "golang.org/x/text*" golang.org/x/text/language/...

package golang.org/x/text/language
  ✔ import sort
  ✔ import strconv
  ✔ import strings
  ✔ import errors
  ✔ import fmt
  ✔ import golang.org/x/text/internal/language
  ✔ import golang.org/x/text/internal/language/compact
package golang.org/x/text/language/display
  ✔ import sort
  ✔ import strings
  ✔ import fmt
  ✔ import golang.org/x/text/internal/format
  ✔ import golang.org/x/text/language
```
