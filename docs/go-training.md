---
permalink: /index.html
title: Learning Go
description: A hands-on training for the Go programming language
theme: gaia
paginate: true
size: 16:9 1920px 1080px
backgroundImage: url(images/xebia-background.svg)
inlineSVG: true
html: true

---
<!-- class: default -->
<style>
section {
    font-size: 1.6em;
}
section.default h6 {
    text-align: right;
    position: relative;
    z-index: 1;
    color: #e2dbc4;
    font-family: "Roboto Mono", monospace;
    font-weight: normal;
    font-size: 0.7em;
    margin-bottom: -1.9em;
    margin-right: 0.3em;
}
p {
    margin-top: 0.5em;
}
pre {
    margin-top: 0.3em;
}
ul {
    margin-block-start: 0.3em;
    margin-block-end: 0.5em;
}
img[alt~="center"] {
  display: block;
  margin: 0 auto;
}
</style>

<!-- _class: lead -->

# Learning Go

![bg right:40%](images/wibautstraat1.jpg)

Laurence de Jong
Software Engineer

---
![bg left:30%](images/wibautstraat2.jpg)

# About me

The Netherlands

DjangoGirls, CouchSurfing, bouldering

Ticketing, trading, inventory, retail, healthcare

Go in production:
 - Blokker (May 2017 - May 2018)
 - Duxxie (June 2018 - January 2020)

---
<!-- _class: lead -->
![bg right:40%](images/wibautstraat3.jpg)

# You

---
<!-- _class: lead -->
![bg left:40%](images/wibautstraat4.jpg)

# Expectations and goals

---
<!-- _class: lead -->
![bg right:40%](images/wibautstraat5.jpg)

# Householding

---

# Agenda

- Startup
- Basics
- Advanced
- Exercises
- Evaluation

---

# Approach

- Learn by example
- See code, run code
- Associate with something you already know
- Interrupt for questions

---
<!-- _class: lead -->
![bg left:40%](images/wibautstraat7.jpg)

# Language

---

# Your experience

- Languages
- Go installed?
- Written Go code?
- A Tour of Go (tour.golang.org)?
- REST API in Go?
- Production Go code?

---

# Why was go created?

[YouTube: Why Learn Go?](https://www.youtube.com/watch?v=FTl0tl9BGdc)

<br>
<iframe width="700" height="400" src="https://www.youtube.com/embed/FTl0tl9BGdc" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

![bg height:300px fit right](images/gopher.png)
![bg right:30%](images/empty.png)

---

# Why was go created?

- 2007 (before Github existed)
- mainstream languages C++, Java, Python, JavaScript
- complex applications
- compilation, execution, development
- insufficient tooling and standards
- rise of multicore processors

---

# Requirements

- Simple to read and understand
- Scale: more teams, bigger infrastructure, more code
- Compile fast, start fast, run efficient
- Safe and reliable
- Modern: batteries included
- Optimum between C++ and Python

---

# Similarities with Java

- General purpose
- Statically typed (type strong)
- Curly braces
- Compiled
- Garbage collected
- Object-oriented but...

---

# Differences with Java

- No generics (but has slice and map)
- No classes (instead it has types)
- No inheritance (but has “embedding”)
- No constructors (but uses “constructor” functions)
- No exceptions
- No annotations/decorators
- No implicit type conversion
- No function overloading
- No JVM required (compiles directly to machine code)

---

# Differences with Java

- Built-in concurrency
- Rich standard library

---

# Useful resources

1. tour.golang.org
1. golangweekly.com
1. github.com/avelino/awesome-go
1. dave.cheney.net
1. calhoun.io
1. golang.org/doc/effective_go.html

---

<!-- _class: lead -->
![bg left:40%](images/wibautstraat8.jpg)

# Setup environment

---

# Installation

Git: git-scm.com/download

Go:
- MacOS `$ brew install go`
- Ubuntu 16.04 LTS, 18.04 LTS and 19.04
```shell script
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt update
$ sudo apt install golang-go
```
- Windows golang.org/dl

---
# Verify installation

```shell script
$ which go
/usr/local/bin
$ go version
go version go1.13.4 linux/amd64
$ go env
...
```

---
# More setup (MacOS, linux)

```shell script
$ # Is the "go"-executable in the PATH env-var?
$ echo $PATH
$ export PATH=${PATH}:/usr/local/bin      # in ~/.bash_profile

$ # Setup GOPATH
$ echo $GOPATH
$ export GOPATH=$HOME/go                  # in ~/.bash_profile

$ # Add the directory of your self-made executables to the PATH env-var
$ export PATH=${PATH}:${GOPATH}/bin       # in ~/.bash_profile
```

---
# Workspace

```shell script
$ ${GOPATH}/
├── bin # binaries
├── pkg # libraries
└── src # executables
```

---
<!-- _class: lead -->
![bg right:40%](images/wibautstraat2.jpg)

# Training material

---

# Get the training material

```shell script
$ go get -v github.com/ldej/go-training/...
```
Everything will end up in
```shell script
$ ${GOPATH}/src/github.com/ldej/go-training/
├── presentation/
├── examples/
└── exercises/
```
Switch to
```shell script
$ cd ${GOPATH}/src/github.com/ldej/go-training
```

---
<!-- _class: lead -->
![bg left:40%](images/wibautstraat3.jpg)

# Exercise: first program
---
# Dev tools
- IntelliJ IDEA or Goland
- Terminal or Windows command
- git

Tip:
- Make sure your editor runs `goimports` on save
- For IntelliJ IDEA / Goland use the File Watchers plugin

```shell script
$ go get golang.org/x/tools/cmd/goimports
```

---
# First program
```shell script
$ cd ${GOPATH}/src/github.com/ldej/go-training
$ mkdir -p hello
$ cd hello
```
Create file `first.go`
###### first.go
```go
package main

import "fmt"

func main() {
    fmt.Printf("Hi %s\n", "everybody")
}
```
```shell script
$ go fmt             # standard formatter (goimports is even better)

$ go run first.go    # compiles and runs right away
$ go build           # creates executable "hello" or hello.exe in .
$ go install         # creates executable "first" in ${GOPATH}/bin

$ hello
Hi everybody!
```
---
<!-- _class: lead -->
![bg left:40%](images/wibautstraat4.jpg)

# Basics
---

# Creating packages
- Group related stuff
- One package per directory
- More coarse-grained than Java: can contain multiple files
- Package name first line of source file
```go
package main // package that results in executable with same name as package
```
or
```go
package news // package that results in library that is accessible via 'news'
```
---
# Using other packages

###### packages.go
```go
package main

import (
    "fmt"  // package from stdlib
    "os"   // package from stdlib
    "time" // package from stdlib

    "github.com/google/uuid" // third-party package
)

func main() {
    u := uuid.New()   // use package-name as prefix
    now := time.Now() // use package-name as prefix

    fmt.Fprintf(os.Stdout, "uuid: %s\ntime: %s", u.String(),
        now.Format(time.RFC3339))
}
```
---
# Comments

Comments
```go
/* a comment */
// another one
```
Document your packages:
- Package level comment
- Every exported (capitalized) name in a program should have a comment

Verify documentation: 
```shell script
$ go doc -all
```
Enforce rules:
```shell script
$ go get -u golang.org/x/lint/golint
$ golint
```
helps you minimize your public exports

---
# Variables `var` `int` `string` `bool`

- Name and type swapped (from Java perspective)
- Have reasonable defaults (not nil)
```go
package main

const myConstString = "golang"

func main() {
    fmt.Printf("my-const-string: %s\n", myConstString)

    var status bool // uninitialized -> default (=false)
    fmt.Printf("status: %v\n", status)

    // := short notation: derives type from right-hand-side
    idx := 256
    fmt.Printf("idx: %d\n", idx)

    longString := `{
        "why": "Useful to embed json in source"
    }`
    fmt.Printf("my-long-string: %s\n", longString)
}
```

---

# Loops `for` `range`

for
```go
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
```
while-like
```go
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
```
iterate
```go
    values := []string{"a", "b", "c"}
    for idx, value := range values {
        fmt.Printf("%d:%s\n", idx, value)
    }
```

---

# If, else `if` `else`

```go
    num := 9
    if num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
```

---

# Switch `switch` `case` `fallthrough`

- On any type
- No fallthrough unless explicitly stated (`fallthrough`)

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}

func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

---

# Exercise 1: Control structures

Create a program that calculates:

- Sum of all values from 1 to 100
- Sum incremental values until their sum exceeds 1000
- Put calculation logic in separate package (library)

---

# Functions `func` `return`

- Core building block
- Scope: based on case
- Java: static methods
```go
func ConvertIt(arg int) string { // public
  return convertInternal(arg)
}

func convertInternal(arg in ) string { // private
  return fmt.Sprintf("My integer value as string: %d", arg)
}
```
Can return multiple values
```go
func swap(x, y string) (string, string) { 
 return y, x 
}
```
(More on functions later)

---

# Defer `defer`

- Cleanup of file-handles, mutexes, channels and connections
- Debugging: log "enter" and "leave" of function
- Unit-testing: "setup" and "teardown"
```go
func enter(name string) string {
    log.Printf("enter %s", name)
    return name
}

func leave(name string) {
    log.Printf("leave %s", name)
}

func main() {
    defer leave(enter("main"))
    log.Printf("in main")
}
```

---

# Error handling

- Multiple return values
- if error is nil, the call worked
```go
resp, err := doSomethingThatCanFail(arg1, arg2)
if err != nil {
    return fmt.Errorf("Error doing something that can fail: %s", err) // early return to minimize indentation
}
// continue with success path

// use _ (=blank) if you don't care
resp, _ := doit(arg1, arg2)
```
- Keep indentation low

---

# Error handling

- Function signature tells that things can go wrong

```go
func doSomethingThatCanFail(arg1 string, arg2 int) (string, error) {
    if arg1 == "" {
        return arg1, fmt.Errorf("arg1 is empty")
    }
    return arg1, nil
}
```

- All your own API's should use this pattern
- Do not use panic and recover

---

# Exercise 2: Error handling

File access using `io/ioutil`:

- Read a file
- Capitalize the content of the file
- Write this capitalized content to a new file with different name
- Use proper error-handling
- Use `defer` to close


```go
import "io/ioutil"
```

---
<!-- _class: lead -->
![bg left:40%](images/wibautstraat5.jpg)

# Data
---

# Struct `struct`

- No constructor
- Case of variable determines accessibility (private, public)

```go
type Student struct { // public
    Name     string   // public
    password string   // private
    teacher  teacher  // private
}

type teacher struct { // not accessible outside package
    Name string
}

func main() {
    student := Student{ // constructor like
        Name:     "John",
        password: "secret",
        teacher: teacher{
            Name: "Laurence",
        },
    }
    fmt.Printf("%+v", student) // %+v: convenience debugging
}
```

---

# Struct methods

```go
type Patient struct {
    Name        string
    YearBorn    int
    IsHealthy   bool
    LastChecked time.Time
}

func (p Patient) HasHighRiskOnDisease() bool { // no side effect
    return (time.Now().Year() - p.YearBorn) > 70 // p => "this"
}

func (p *Patient) MarkHealthy() { // has side effect
    p.IsHealthy = true
    p.LastChecked = time.Now()
}

func main() {
    dada := Patient{
        Name:     "Inder",
        YearBorn: 1940,
    }
    log.Printf("high-risk: %+v\n", opa.HasHighRiskOnDisease())
    dada.MarkHealthy()
    log.Printf("after: %+v\n", opa)
}
```

---

# Struct methods: value or pointer?

> 1. Use the same receiver type for all your methods. This isn't always feasible, but try to.
> 2. Methods defines a behavior of a type; if the method uses a state (updates / mutates) use pointer receiver.
> 3. If a method don't mutate state, use value receiver.
> 4. Functions operates on values; functions should not depend on the state of a type.

<!-- _footer: Source: https://dev.to/chen/gos-method-receiver-pointer-vs-value-1kl8 -->

---

# Pointers `*` `&`

- Default value: nil
```go
var ho *HugeObject = &HugeObject{} // ho := &HugeObject{}
ho := new(HugeObject)
insuranceService.CalculateRisk( hu )
```
- For methods that mutate data
```go
func (p Patient)MarkDeceased() { // won't adjust patient
  p.Deceased = true
}
func (p *Patient)MarkDeceased() { // will work
  p.Deceased = true
}
```
- Indicate Optional (poor mans)
```go
type Person struct {
  Name string
  Child *Person // optional
}
```

---

# Pointers: Are pointers a performance optimization?

The short answer: __No__

- Stack (for function local data)
- Heap (for shared data)

When __not__ to use:
- When you __*think*__ it __*might*__ give you better performance

When to use:
- For data mutation
- As optional
- After profiling indicates that copying is a problem

Many types, such as slices, strings, and maps, contain pointers to underlying data, passing pointers to these types rarely makes sense.

<!-- _footer: Source: https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85 --> 

---
 
 # Enumerations
 
 ```go
type Color int

const (
    Unknown    Color = iota // 0 (=default)
    Red                     // 1
    Green                   // 2
    Blue                    // 3
)

func (c Color) String() string {
    switch c {
    case Green:
        return "green"
    case Blue:
        return "blue"
    case Red:
        return "red"
    default:
        return "unknown"
    }
}

func main() {
    var myColor Color // uses default
    otherColor := Green
    fmt.Printf("my-color: %v (%d), other-color: %v (%d)\n",
        myColor, myColor, otherColor, otherColor)
}
```

---

# Exercise 3: Data modeling

- Model your business domain using structs, enums and pointers
- Might need slices (see next section)

If you can't think of anything:
- A school with teachers, students, lessons, rooms

---

# Containers
- array and slice
- map

---

# Slice

- Can contain everything: primitives, structs, slices, maps etc
- Like Java ArrayList
- Sortable
- Supported operations: `append`, `replace`, `[idx]`, `[idx-from:idx-to]`, iterate

Fixed length immutable
```go
numbers := [4]int{10, 20, 30, 40}
s := [...]string{"Cheese", "Coffee"} // idiomatic: let compiler count
```
Dynamic size
```go
var slice0 []string = []string{}   // empty
slice1 := []string{}               // empty
slice2 := []string{"a", "b", "c"}  // initialize with data
slice3 := make([]string, 0, 5)     // optimization: empty with reserved capacity
```
Not thread safe (combine with Mutex)

---

# Slices in action

```go
func main() {
    letters := []string{"a", "b", "c", "d"}
    fmt.Printf("before: %v:   length: %d, capacity: %d (%p)\n",
        letters, len(letters), cap(letters), letters)

    // add items
    // append(letters, "e") // wrong!!!
    letters = append(letters, "e") // why? realloc when no longer fits
    fmt.Printf("after:  %v: length: %d, capacity: %d (%p)\n\n",
        letters, len(letters), cap(letters), letters) // pointer has changed

    // access items
    fmt.Printf("first:   %v\n", letters[0])              // a
    fmt.Printf("nothing: %v\n", letters[2:2])            // []
    fmt.Printf("begin:   %v\n", letters[:2])             // [a b]
    fmt.Printf("middle:  %v\n", letters[1:3])            // [b c]
    fmt.Printf("end:     %v\n", letters[3:])             // [d e]
    fmt.Printf("last:    %v\n", letters[len(letters)-1]) // e

    // iterate
    for idx, value := range letters {
        fmt.Printf("values[%d] = %s\n", idx, value)
    }
}
```

---

# Map
- Store key-value pairs (like Java HashMap)
- Typically key is primitive, value can be everything: primitives, structs, slices, maps etc
- Supported operations:
`get-on-key, put-on-key, replace-on-key, delete-on-key, iterate`
- initialization:
```go
var m1 map[string]int = make(map[string]int)
m2 := make(map[string]int)
m3 := map[string]int{}
m4 := map[string]int{
     "route": 66,
}
```
Random iteration order
Not thread safe (combine with Mutex)

---

# Maps in action

```go
func main() {
	studentsOnSchool := map[string][]string{
		"Cambridge": []string{"Raj", "Alice"},
		"MIT":       []string{"Bob"},
	}
	fmt.Printf("1: %+v\n", studentsOnSchool)   // %+v debugging convenience

	studentsOnSchool["DU"] = []string{"Abhi"}  // add map entry
	fmt.Printf("2: %+v\n", studentsOnSchool)   // %+v debugging convenience

	delete(studentsOnSchool, "MIT")            // remove map entry
	fmt.Printf("3: %+v\n", studentsOnSchool)   // %+v debugging convenience

	cambridgeStudents, found := studentsOnSchool["Cambridge"] // get map entry
	if !found {
		cambridgeStudents = []string{}
	}
	cambridgeStudents = append(cambridgeStudents, "Neha")
	studentsOnSchool["Cambridge"] = cambridgeStudents // put map entry

	for key, value := range studentsOnSchool { // iterate map
		fmt.Printf("4: %s - %v\n", key, value)
	}
}
```

---

# Exercise 4: slices and maps

Use maps and slices to group the following people on hobby:
```go
Julia:  cricket, drawing
Sophie: drawing
Mila:   drawing
Emma:   tennis, kabaddi
Neha:   running
Abhi:   photography, cricket
Noor:   cricket
Elin:   hockey
Sara:   cricket, kabaddi
Yara:   tennis
```

---

<!-- _class: lead -->
![bg right:40%](images/wibautstraat6.jpg)

# Interfaces

---

# Interface `interface`

- Duck-typing "If it walks like a duck and it quacks like a duck, then it must be a duck"
- no explicit "implements"
- Good to improve testability
example from stdlib
```go
package fmt
// Accepts anything that implements the "Writer"-interface:
// Examples of Writers: file, buffer, stdout, network, http-response, zip-file etc
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) { ... }
```
other example
```go
// The business logic layer can return this EnrichedError as regular 'error' .
// The http layer convert this error into an appropriate http-response (200, 400, 403, 500 etc)
type EnrichedError struct {
     Kind ErrorKind // invalid-input, not-authorized, internal-error etc
     Message string
}
func (e HttpError) Error() string { // implement Error-interface
   return e.Message
}
```

---

# Interface `interface`

- Naming convention: ends with "er"
- Keep them small. Why?
- Composeable

```go
type Datastorer interface {
    Put(key string, value interface{}) error
    Get(key string) (interface{}, bool, error)
    Remove(key string) error
}
```

---

# Example usage of interface

- For dependency injection
- The business logic of PatientService is testable without a "real"-datastore

```go
func main() {
    // Inject Datastorer into business logic service
    patientService := NewPatientService(NewSimplisticDatastore())

    patient := Patient{UID: "patient-12345", FullName: "Abhi Kumar", Allergies: []string{"peanuts"}}

    // Initialize with data
    err := patientService.Create(patient) // uses Datastorer.Put
    if err != nil {
        log.Fatalf("Error creating patient: %s", err)
    }

    // Adjust patient
    err = patientService.MarkAllergicToAntiBiotics(patient.UID) // uses Datastorer.Get and Put
    if err != nil {
        log.Fatalf("MarkAllergicToAntiBiotics error: %s", err)
    }
}
```

---

# Exercise 5: interfaces

Implement a simple in-memory database that implements the following interface:

```go
type Datastorer interface {
    Put(key string, value interface{}) error
    Get(key string) (interface{}, bool, error)
    Remove(key string) error
}
```

---

<!-- _class: lead -->
![bg left:40%](images/wibautstraat7.jpg)

# Testing

---

# Testing

Essential for software with a long predicted lifetime

Why:

- Need “Safety-net" so you dare to keep improving and extending your software

How:

- Do from beginning
- Tests should be easy and fast to run
- Test against the API
- Only test against internals in specific cases
- Prefer HTTP (as the ultimate API) to trigger your business logic, so you have freedom to change the internals and still keep your safety net intact

---

# Unit testing

- Part of toolchain
- In same package, dedicated file
- Filename convention: <file_name>_test.go

###### reverse_test.go
```go
import (
    "testing"
)

// Naming convention:  starts with Test and has "t *testing.T" as parameter
func TestReverse(t *testing.T) {
	value := Reverse("ecnerual")
	if value != "laurence" {
		t.Errorf("Reverse() = %v, want %v", value, "laurence")
	}
}
```
```shell script
$ go test
--- FAIL: TestReverse (0.00s)
    reverse_test.go:10: Reverse() = ecnerual, want laurence
FAIL
exit status 1
```

---

# Table driven testing

- Used a lot in the stdlib
- Very readable
- Easy to be complete
```go
func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{input: "abc", sep: "/", want: []string{"abc"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Split(tt.input, tt.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

---

# Design for testability: Use dependency injection

```go
package uniqueid

// Generator provides an interface for creating uuid's.
type Generator interface {
    Generate() string
}
```
Service must internally create uuid's:
```go
package patient

type Service struct {
    uuidGenerator uniqueid.Generator
}

func New(uuidGenerator uniqueid.Generator) *Service {
    return &Service{
        uuidGenerator: uuidGenerator,
    }
}
```
This pattern is suitable for managing access to databases

---

# Mocks and stubs

- Create by hand
- Generate from interface using mockgen (https://github.com/golang/mock)
```go
package uniqueid

//go:generate mockgen -source=uniqueid.go -destination=uniqueid_mock.go -package=uniqueid Generator

type Generator interface {
    Generate() string
}
```

---

# Benchmarking
- Premature optimization is ...
- Never optimize before measuring first

```go
// trigger benchmark with: go test -bench=.

// Naming convention: starts with "Benchmark" and has "b *testing.B" as parameter
func BenchmarkDoCalculationByValue(b *testing.B) {
    // run the function b.N times
    for n := 0; n < b.N; n++ {
        bs.DoCalculationByValue()
    }
}

func BenchmarkDoCalculationByReference(b *testing.B) {
    for n := 0; n < b.N; n++ {
        (&bs).DoCalculationByReference()
    }
}
```

---

# More on testing

- Standard library offers utils for testing http-clients and http-servers
- Code coverage: IntellijIDEA
- Race condition detection
- Continuous integration

---

# Exercise 6: tests

- Write a table driven test for validating email-addresses
- Make the 'nontestable'-package testable
- Run the benchmark and find out at what 'size' by value and by reference have equal performance

###### examples/nontestable/nontestable.go

```go
package nontestable

import (
	"io/ioutil"
	"time"

	"github.com/google/uuid" // third-party package
)

func Write() error {
	u := uuid.New()
	ft := time.Now().Format(time.RFC3339)
	return ioutil.WriteFile(u.String()+".txt", []byte(ft), 0644)
}
```

---

# What about assertions?

https://golang.org/doc/faq#assertions
> Go doesn't provide assertions. They are undeniably convenient, but our experience has been that programmers use them as a crutch to avoid thinking about proper error handling and reporting. Proper error handling means that servers continue to operate instead of crashing after a non-fatal error. Proper error reporting means that errors are direct and to the point, saving the programmer from interpreting a large crash trace. Precise errors are particularly important when the programmer seeing the errors is not familiar with the code.
> 
>  We understand that this is a point of contention. There are many things in the Go language and libraries that differ from modern practices, simply because we feel it's sometimes worth trying a different approach.

---

# Assertions!

https://github.com/stretchr/testify

```go
package yours

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")
  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")
  // assert for nil (good for errors)
  assert.Nil(t, object)
  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {
    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)
  }
}
```

---

<!-- _class: lead -->
![bg right:40%](images/wibautstraat8.jpg)

# Concurrency

---

# Concurrency

Overemphasized
Most of your code is synchronous
Concurrent-style not forced upon you, used selectively

- built-in "channels" and "goroutines"
- Goroutines: think very, very lightweight threads
- Channels: think pipe or queue to communicate with goroutine(s)
- "select"-loop: UNIX-like: wait for events from multiple channels

---

# Channels `chan`

Do not communicate by sharing memory. Instead, share memory by communicating.

```go
func sum(a []int, resultChannel chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    resultChannel <- sum // send result back over channel
}

func doit() {
    responseChannel := make(chan int) // construct channel
    defer close(responseChannel)      // prevent resource leak

    go sum([]int{1, 2, 3}, responseChannel)      // 1 + 2 + 3 = 6
    go sum([]int{4, 5, 6}, responseChannel)      // 4 + 5 + 6 = 15
    x, y := <-responseChannel, <-responseChannel // receive from channel

    fmt.Printf("one=%d\nanother=%d", x, y) // order undefined
}
```

---

# Select `select`

Wait for events from multiple channels

```go
func sendMsg(c chan string) {
    time.Sleep(100 * time.Millisecond)
    c <- "Put your helmet on"
}

func main() {
    tick := time.Tick(800 * time.Millisecond)
    boom := time.After(3 * time.Second)
    msgChannel := make(chan string)
    go sendMsg(msgChannel)
    for {
        select { // blocking until msg received on one of its channels
        case msg := <-msgChannel:
            fmt.Printf("msg: %s\n", msg) // stay in loop
        case <-tick:
            fmt.Println("tick.") // stay in loop
        case <-boom:
            fmt.Println("BOOM!")
            return // abort loop
        }
    }
}
```

---

# Concurrency example usage
- Devide and conquer: Chunk your problem in parts that can be processed concurrently in isolation
- Send out multiple requests concurrently: Continue when all responses have been received
- Cache eviction: Scheduled background "thread" removes old entries at regular intervals
- Fan out: Fire and forget
- Respond faster by moving non-criticals (like notification) off the main thread
- Synchronized state-machine: select loop is the heart

NB:
- When multiple go-routines access the same data, you will need to synchronize (with mutexes or channels)
- Cleanup: don't leave dangling channels and goroutines

---

# Exercise 7: concurrency

Execute a slow action (see function below) concurrently (100x) and wait for all the results for no more than 1 second. Report the number of results received.

```go
func SimulateSlowAction( a, b int) int {
    sleepDurationInMillsec := 500 + (rand.Intn(1000)) // what does this do?
    time.Sleep(time.Duration(sleepDurationInMillsec) * time.Millisecond)
    return a*b
}
```

Create a "thread-safe" in-memory cache (Get, Put) that automatically removes entries older than 10 seconds

---

<!-- _class: lead -->
![bg left:40%](images/wibautstraat2.jpg)

# More on functions

---

# Package initialisation

- Executed only once at startup
- Can be multiple
- For advanced initialisations
```go
func init() {
  // your global initialisations here
}
```

---

# Variadic functions

```go
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

---

# Closures

- Lambda's in Java
- Alternative to single method interfaces

example: binary search

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{1, 11, -5, 8, 2, 0, 12}
    sort.Ints(numbers)
    fmt.Println("Sorted:", numbers)

    index := sort.Search(len(numbers), func(i int) bool {
        return numbers[i] >= 7
    })
    fmt.Printf("The first number >= 7 is at %d and has value %d", index, numbers[index])
}
```

---

# More examples usages of closures

```go
func main() {
    a := 1
    b := 1004223
    go func() {
        // this block runs in background within a go-routine
        result := doWork1(a, b)
        result = doWork2(result)
        result = doWork3(result)
        log.Printf("Result:%s\n", result)
    }()
    log.Printf("Continue without waiting for result\n")

    time.Sleep(time.Second * 5) // Why is this needed?
    log.Printf("main terminates\n")
}
```

---

<!-- _class: lead -->
![bg left:40%](images/wibautstraat3.jpg)

# The standard library

---

# Rich standard libraries
- flags
- file I/O
- sync
- logging
- os
- sort
- networking
- http, http2: client and server
- encoding (json, xml, mime)
- compression
- crypto
- templates
- sql

---

# Serialisation

- Based on struct tags
- xml, json etc in stdlib

example:
```go
type Person struct {
   Name      string   `json:"name"      xml:"PersonName"`
   Interests []string `json:"interests" xml:"PersonInterests"`
   Children  []Child  `json:"children"  xml:"Person_Children"`
}

type Child struct {
   Name string `json:"name"          xml:"name"`
   Age  int    `json:"age,omitempty" xml:"age,omitempty"`
}
```

<!-- _footer: https://github.com/golang/go/wiki/Well-known-struct-tags -->

---

# Exercise 8: json and xml

- Create some structs with different primitives and types
- Experiment with json and xml tags and directives

Tips:
- https://mholt.github.io/json-to-go/
- https://godoc.org/encoding/json#Marshal
- https://godoc.org/encoding/xml#Marshal
- https://github.com/golang/go/wiki/Well-known-struct-tags

---

# Flags: read command-line arguments

```go
func printUsage() {
    fmt.Fprintf(os.Stderr, "\nUsage:\n")
    fmt.Fprintf(os.Stderr, " %s [flags]\n", path.Base(os.Args[0]))
    flag.PrintDefaults()
    fmt.Fprintf(os.Stderr, "\n")
    os.Exit(1)
}

func main() {
    login := flag.String("login", "", "GitHub login of user")
    once := flag.Bool("once", false, "Perform action once")
    reps := flag.Int("reps", 10, "Number of reps")
    flag.Parse()

    if *login == "" {
        printUsage()
    }

    for idx:=0;idx<*reps; idx++ {
        log.Printf("Looking up GitHub user: %s (once:%v)", *login, *once)
    }
}
```

---

<!-- _class: lead -->
![bg right:40%](images/wibautstraat4.jpg)

# HTTP

---

# HTTP client

Request and response:
- Request.Method: `POST`, `PUT`, `DELETE` and `GET`
- Request.Url: REST-ful?
- *.Headers: Content-Type, Accept, Authorization (Basic or Bearer)
- *.Payload: Json, XML
- Timeout
- Response.StatusCode: 200, 201, ... 400, 401, 403, 404, ... 500, 501, 503

Based on Swagger/OpenApi-spec?
Standard library provides API

---

# HTTP client example

###### example/httpClient/httpClient.go
```go
type GetPatientResponse struct {
	UID       string   `json:"uid"`
	FullName  string   `json:"full_name"`
	Allergies []string `json:"allergies"`
}

type Client struct {
	Hostname string
}

func (cl *Client) GetPatientOnUID(patientUid string) (*GetPatientResponse, error) {
	client := http.Client{}
	httpResponse, err := client.Get(fmt.Sprintf("%s/api/v1/patients/%s", cl.Hostname, patientUid))
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching patient: http-status %d", httpResponse.StatusCode)
	}
	var resp GetPatientResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
```

---

# Exercise 9: HTTP client POST

- Create a patient using HTTP POST to https://patient-store.appspot.com/api/v1/patients

```shell script
$ curl -X POST \
    --data '{"full_name": "Laurence", "allergies": ["coffee"]}' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    https://patient-store.appspot.com/api/v1/patients
```
- Play around at https://patient-store.appspot.com/swagger/index.html
- Create a unit test for your HTTP client
    - simulate server behaviour using `httptest`
    - create one for success and one for not found
    
---

# HTTP server

- Methods: POST, PUT, DELETE, GET AND HEAD
- Url: REST-ful?
- Payload: Json, XML
- Middleware for non-functionals (CORS, Auth, Monitoring)
- Response status-code and error-message

Based on Swagger/OpenApi-spec?
Standard library provides API:
- HTTP/2 capable
- Each request runs in its own goroutine
- File server and reverse proxy included

---

# HTTP server

Clean Architecture

```go
func main() {
    var router *mux.Router = mux.NewRouter()
    
    storeInstance := store.NewInMemoryStore()

    patientRepository := repository.NewPatientRepository(storeInstance)

    patientService := services.NewPatientService(patientRepository)
    
    web.NewPatientHandler(router, patientService)
    
    http.Handle("/", router)
}
```

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047

---

# Clean Architecture

![center](images/clean-arch.png)

---

# Clean Architecture

__Delivery__

- Translates HTTP to our models and passes them to service
- Does not contain business logic

__Service__

- Contains our business logic
- Has nothing with http
- Interacts with repository

__Repository__

- Does not contain any business logic
- Translates our models so they can be passed on (to databases, to microservice)

__Models__

- Does not interact with anything
- They are used in our business logic

---

