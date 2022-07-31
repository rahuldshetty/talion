# Talion

Interpreter programming language written with Go language.

- Complete code written with native Go modules that includes Lexer, Parser & Evaluator.
- Tree walking interpreter
- Operator:
    - Unary: Not(!), Minus (-)
      ```
      !true 
      !false
      -10
      -(-10)
      ```

    - Binary Operator: Addition, Subtraction, Division, Multiplication
        ```
        100+(-100)
        5*4+(1-2)
        42 - 5/5
        ```

    - Conditional Operator: Equal To(==), Not Equal To(!=), Greater Than (<), Lesser Than(<), Less Than or Equal to (<=), Greater Than or Equal To(>=)
        ```
        true == true
        1 >= 1
        5 < 95
        false == (1>1)
        ```

- Support Datatype: Integer, String, Boolean, Null Types, List
    ```
    var a = 10;
    var b = "hello"
    var c = true;
    var j = "hello" + " " + "world"
    ```
- Data Structures:
    - List: 
        Ordered sequence of elements of any datatypes. 

        Lists in talion support zero-indexed and negative indexing. 

        Out of bound index values will return null object in response.
        ```
        var l = [1, 2, 3, 4]
        var b = 4
        var ls = [1, 2, b, "hello"] 
        ls[0]
        ls[2]
        ls[-1]
        len(ls) // return 4
        ```

- Null/truth based if else conditional statements.
- Variable binding: 
    ```
    var a = 100;
    var b = a + (100 - 20)
    ```
- Functions: 
    ```   
    var add = fn(x, y) { return x + y; }
    add(1, add(1, 1))
    ```

- Closure Support & Higher Order function

    ```
    var adder = fn(x){ fn(y){ x + y } }
    var add_two = adder(2)
    add_two(20) // Result is 22
    ```
- Garbage Collector: Leverages Go's GC to manage memory in talion language.

- Builtin Functions:
  - len(<string_object>)
    ```
    var s = "Hello world"
    len(s) // returns 11
    ```