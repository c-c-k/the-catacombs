package main

import "fmt"

func main() {
    // Basic variable types and default values
    // var var-name var-type  // default value
    var i int  // 0
    var f float64  // 0
    var b bool  // false
    var s string  // ""
    fmt.Println("var i int\nvar f float64\nvar b bool\nvar s string")
    fmt.Println(i, f, b, s)


    // Full form variable declaration and assignment
    var demoInt1 int = 1
    fmt.Println("var demoInt1 int = 1")
    fmt.Println(demoInt1)

    // Variable type can be dropped when it can be clearly inferred
    // from the initial value
    var demoInt2 = 2
    fmt.Println("var demoInt2 = 2")
    fmt.Println(demoInt2)

    // The declare and assign operator `:=` can be used instead 
    // of the var keyword when the type can be inferred
    demoInt3 := 3
    fmt.Println("demoInt3 := 3")
    fmt.Println(demoInt3)

    // The `:=` operator is specifically intended for declaring and assigning 
    // and not just assigning, trying to use it for assignment only
    // will cause an error.
    // demoInt3 := 33  // error: no new variables on left side of :=

    // An assignment to an existing variable will work if it is done
    // alongside the declaration and assignment of another variable though
    demoInt3, demoInt4 := 33, 4
    fmt.Println("demoInt3, demoInt4 := 33, 4")
    fmt.Println(demoInt3, demoInt4)

    // Multiple variables of the same type can be declared with a single
    // regular var statement but must all either be assigned or not assigned
    // an initial value.
    var demoInt5, demoInt6, demoInt7 int  // All are equal to 0
    fmt.Println("var demoInt5, demoInt6, demoInt7 int")
    fmt.Println(demoInt5, demoInt6, demoInt7)
    var demoInt8, demoInt9, demoInt10 int = 8, 9, 10
    fmt.Println("var demoInt8, demoInt9, demoInt10 int = 8, 9, 10")
    fmt.Println(demoInt8, demoInt9, demoInt10)
    // var demoInt11, demoInt12, demoInt13 int = 5, 6 // error: missing init expr for demoInt13 

    // To declare multiple values of different types or/and with some being
    // assigned and some not being assigned an initial value a block of
    // variables can be used
    var (
        demoInt14 int = 14
        demoInt15 int
        demoBool1 bool
        demoFloat1 = 3.14
        demoString1 = "hello 1"
    )
    fmt.Println("var (\ndemoInt14 int = 14\ndemoInt15 int\ndemoBool1 bool\ndemoFloat1 = 3.14\ndemoString1 = \"hello 1\"\n)")
    fmt.Println(demoInt14, demoInt15, demoBool1, demoFloat1, demoString1)

    // Constants are declared with the `const` keyword and like variables can
    // be declared in blocks and drop the type if it can clearly be inferred
    // from the assigned value
    const demoFloat2 float64 = 2
    const demoInt16 = 16
    const (
        demoInt17 = 17
        demoBool2 = true
        demoFloat3 = 3.14
        demoString2 = "hello 2"
    )
    fmt.Println("const demoFloat2 float64 = 2\nconst demoInt16 = 16\nconst (\ndemoInt17 = 17\ndemoBool2 = true\ndemoFloat3 = 3.14\ndemoString2 = \"hello 2\"\n)\n")
    fmt.Println(demoFloat2, demoInt16, demoInt17, demoBool2, demoFloat3, demoString2)

    // A special semi-inumeration keyword in Golang is `iota`.
    // It's simple and full syntax usage is like
    const (
        salesID = iota
        productionID = iota
        managmentID = iota
    )
    fmt.Println("const = (\nsalesID = iota\nproductionID = iota\nmanagmentID = iota\n)\n")
    fmt.Println(salesID, productionID, managmentID)

    // Only the first `iota` keyword is needed and it can be used to change the
    // starting count value
    // Also an underscore can be used to skip a value
    const (
        sunday = iota + 1
        _
        tuesday
        _
        thursday
        _
        satureday
    )
    fmt.Println("const (\nsunday = iota + 1\n_\ntuesday\n_\nthursday\n_\nsatureday\n)\n")
    fmt.Println(sunday, tuesday, thursday, satureday)

    // A common use case for the iota is bitmasks
    const (
        ExecuteBit = 1 << iota
        WriteBit
        ReadBit
    )
    fmt.Println("oonst (\nExecuteBit = 1 << iota\nWriteBit\nReadBit\n)\n")
    fmt.Println(ExecuteBit, WriteBit, ReadBit)

    // Another unusual aspect of Golang is the rune type which is an alias for the int32 type and is usually denoted by a single character surrounded by apostrophes e.g. `'A'` and resolves to the characters ASCII Unicode value
    fmt.Println("'A' = ", 'A')
    fmt.Println("'啐' = ", '啐')

    // A pointer is created by prefixing a variable with "&" and resolved by
    // prefixing the pointer variable with "*"
    target := "target"
    pointer := &target
    fmt.Println("target := \"target\"\npointer := &target")
    fmt.Println(pointer, *pointer)
}



