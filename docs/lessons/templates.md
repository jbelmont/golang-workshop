# Templates

* Template golang package is used to help handle templates (i.e. html).

Template functions:

* Name: returns the name of the template
* Parse: Parse parses text as a template body for t.
* Lookup: returns the template with the given name that is associated with t.


*Find many more functions in the official Golang package below*
[Templates](https://golang.org/pkg/text/template/)


```go
package main

import (
        "log"
        "os"
        "text/template"
)

type Person struct {
        Name   string
}

const tmpl = `Hello {{.Name}}`

func main() {
    person := Person{
        Name:   "John Rambo",
    }

    t := template.New("Person template")

    t, err := t.Parse(tmpl)
    if err != nil {
            log.Fatal("Parse: ", err)
            return
    }

    err = t.Execute(os.Stdout, person)
    if err != nil {
            log.Fatal("Execute: ", err)
            return
    }
}

// outputs "Hello John Rambo"
```

Notice here that we declare a const variable of `tmpl`

* `template.New` allocates a new HTML template associated with the given one and with the same delimiters.
* The association, which is transitive, allows one template to invoke another with a `{{template}}` action.

[Templates Go Playground](https://play.golang.org/p/2pO4koslxq)
