# zerror

## Introduction
In the common **error** interface, there is only a string message. When a program collects the errors and processes them 
in a certain layer instead of handling right after they occur. There are some information lost, such as: 
the time, the place and the severity.

This library attempts adding the extra information to solve the issues.

## Examples

### Concepts
* **Time**: the creation timestamp of a zerror, and the default time format is RFC3339.
* **Severity**: there are three level severity in a zerror: Fatal, Warn, and Info. Default level is **Info**.
* **Message**: the error message set in **New** functions or overwritten in **SetMessage** function.
* **File**: the source file generating the error.
* **Function**: the function the error occurs.
* **Line**: the line number the error occurs.

In some cases, We need an original error instance generated in a third party function/library for the following **if** 
conditions, such as **sql.ErrNoRow**, and we could save them by **SetError** function. By default, **GetError** function 
returns an error instance with **Message**.

### Create and print

```go
e := zerror.New("example 1, %s", "zerror")
fmt.Printf("%v\n", e.String())
```
output:
```text
2019-07-27T11:17:34-04:00  Info : example 1, zerror(main.main:10)
```
We could also use **NewFatal/NewWarn/NewInfo** function to create different severity instance.

**String** function formats the error information, and generates a printable string. **Severity** is aligned in 5 letters.

### Change severity globally

```go
zerror.SetDefaultSeverity(zerror.SeverityWarn)
e = zerror.New("example 2, %s", "zerror")
fmt.Printf("%v\n", e.String())
```
output:
```text
2019-07-27T11:17:34-04:00  Warn : example 2, zerror(main.main:14)
```

### Change output format in **String** function return globally
```go
zerror.SetMessageFormat(fmt.Sprintf("%s:%s - %s: %s(%s)",
		zerror.NotationFile, zerror.NotationLine, zerror.NotationSeverity, zerror.NotationMessage, zerror.NotationTime))
e = zerror.New("example 3, %s", "zerror")
fmt.Printf("%v\n", e.String())
```
output:
```text
/Users/szqmtl/goproj/src/experience/zerrorTest.go:19 -  Warn : example 3, zerror(2019-07-27T11:17:34-04:00)
```

### Change time format globally
```go
zerror.SetTimeFormat(time.RFC822Z)
e = zerror.New("example 4, %s", "zerror")
fmt.Printf("%v\n", e.String())
```
output:
```text
/Users/szqmtl/goproj/src/experience/zerrorTest.go:23 -  Warn: example 4, zerror(27 Jul 19 12:09 -0400)
```

Take a look the test file for more and complete examples.