# Packages

## Definition

*Definition according to `The Go Programming Lanugage` book for packages`*

> The purpose of any package system is to make the design and maintenance of large programs practical by grouping related features together into units that can be easily understood and changed, independent of the other packages of the program. This modularity allows packages to be shared and reused by different projects, distributed within an organization, or made available to the wider world.

* Packages provide a way to encapsulate functions, structs, types, etc by controlling the visibility of these types in packages
* Restricting the visibility of package members hides helper functions and types behind the packages' API
    * Allowing the package maintainer to change the implementation and not affecting code outside the package
* Restricting visibility hides variables so that clients can access and update them through:
    * exported functions that preserve internal invariants

## Usage

* A package declaration is required at the start of every Go source file.
* Its main purpose is to determine the default identifier for that package called the package name when
    * it is imported by another package.
* A Go source file may contain zero or more import declarations immediately after the package
    * declaration and before the first non-import declaration.
* Each import declaration may specify the import path of a single package
* or multiple packages in a parenthesized list

```go
import "fmt"
import "os"

import (
    "fmt"
    "os"
)
```

**Notice above the 2 different styles and both are valid, albeit the latter is more concise**

```go
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Code-Craftsmanship-Saturdays/software-security/model"
)
```

* You can also group packages by adding blank lines.
* The first group is part of the standard library and last group is part of a package I wrote
* These packages must be used are there will be `unused import error`

`import _ "image/png" // register PNG decoder` this is known as a blank import and it doesn't have to be used in the package

* Any function, types, etc that you capitalize will be made Public outside of the package
    * otherwise they are private
