- Initialize an application with the command: go mod init booking-app
- All code must belong to some package
- Run code using: go run

Data Type in GO:
- STRINGS   -> Textual Data
- INTEGERS  -> Numeric Data
- FLOAT     -> Decimal Data
- BOOLEANS
- MAPS
- ARRAYS    -> Arrays in GoLang have a fixed size and we can't mix datatypes in arrays
            New arrays can be declared in the following ways
            1. var bookings [50]string
            2. var bookings = [50]string{"Vedansh"}

- SLICES    -> Basically a list [No size declaration, same as python list]
            New slices can be declared in the following ways
            1. var bookings []string
            2. var bookings := []string{"Vedansh"}


Loops:

- GoLang only has FOR loops
    - Infinite Loop        for{}
    - For Each Loop    for ind, val := values {}
    - Execute code until some condition is met     for a > 0 && b < 0 {}  


Conditionals:

- if a < b {

    }  else if c < d {

    } else {

    }
- switch city {
    case "New York": // do something and break
    case "US", "USA": // do somethin and break
    default: // do something
}


Functions:

- Groups logic that belongs together
- Reuse logic to avoid duplication of code
- func doSomething() <return type>{

}
 doSomething()


 Package Level Variables:

 - Defined outside of all the functions
 - Can be accessed inside any function and in all files which are in the same package




