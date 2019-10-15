# magic

:tophat: Package `magic` is an auto-parsing library and competitive coding helper package with batteries included. You just need to initialize it with a `SrcProvider` (e.g. a string, file, stream, etc.) and subsequently whish a type to auto-parse. The library takes care of allocating and populating your memory. The sequence of inputs is defined by the memory layout of the reading-type.

**In no circumstance you should use this package for production code. Use it only during competetive coding, when coding time matters and input parsing is trivial. I originally built this for the [Catalysts Coding Contest (codingcontest.org)](http://codingcontest.org)**

[![GoDoc](https://godoc.org/github.com/harwoeck/magic?status.svg)](https://godoc.org/github.com/harwoeck/magic)
[![Go Report Card](https://goreportcard.com/badge/github.com/harwoeck/magic)](https://goreportcard.com/report/github.com/harwoeck/magic)

## Installation

To install `magic`, simly run:

```bash
$ go get -u github.com/harwoeck/magic
```

## Quick start

1. Create a `SrcProvider`

    ```go
    src := bufio.NewScanner(strings.NewReader("INPUT"))

    f, _ := os.Open("/path/to/file.txt")
    src := bufio.NewScanner(f)
    ```

2. Create a new `Manager` instance

    ```go
    m := magic.NewManager(src)
    ```

3. `Read`, either primitive and complex types, from the manager

    ```go
    parsedBool := m.ReadBool()
    parsedInt := m.ReadInt()
    parsedFloat := m.ReadFloat32()
    parsedString := m.ReadString()

    parsedStructList := m.Read([]myStruct).([]myStruct)
    ```

### Complete example

In the following example a complex type (slice of structs) is read and auto-parsed. Within the main struct there are uncaped slices. In these situations `magic` will dynamically try to infer the size by reading an integer from the `SrcProvider` and incremting the cursor position accordingly.

```go
type tx struct {
    id        int64
    authority string
    from      []txpart
    to        []txpart
}

type txpart struct {
    origin string
    amount float64
}

func main() {
    src := bufio.NewScanner(strings.NewReader(
        // Image the strings below are single lines within an input-file
        "1 FINANCE_MINISTRY 1 YourBankAccount 100 1 MyBankAccount 100\n" +
        "2 FINANCE_MINISTRY_DE 2 YourBusinessAcount1 50 YourBusinessAcount2 150 1 MyBusinessAcount 200\n"
    ))

    m := magic.NewManager(src)

    for _, item := range m.Read([]tx{}).([]tx) {
        fmt.Printf("%d %s\n", item.id, item.authority)
    }
}
```
