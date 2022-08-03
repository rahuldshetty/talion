
# Function

Functions are sequence of instruction grouped together as one executable unit for re-usability of logic. 

- Function in talion are first-class functions. 
- You can define closure in talion language. Function defined within other functional block can persist all the environment scoped variables which is used to define closure.

- Example:
    ```
    adder = fn(x){ return fn(y){y+x} }
    add_two = adder(2)
    add_two(10)                                // 12
    ```

- Variables declared within the functional block are only scoped within the function.