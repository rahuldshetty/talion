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
    my_var = 10
    my_n = 10 + my_var
    var a = 10;
    var b = "hello"
    var c = true;
    var j = "hello" + " " + "world"
    "hello"[1]
    ```
- Data Structures:
    - List: 
        Ordered sequence which can store elements of any datatypes. Mutable by nature. 

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
        push(ls, 1) // push to end of array
        ls[-1] // return 1
        ```
        
    - Hash: Dictionary/Map based datastructure that stores key, value pair object. 
      Datatype supported to be hashed as key - Integer, Boolean, String
      Supports Get/Set through simple index & assignment operators.
      ```
      a = {1:"one", 2:"two", "3": 3, 3: "three"}
      a[1]
      a[2]
      a[3]
      a["3"]

      a["three"] // NULL if key doesn't exist
      a["three"] = 102 // create new key
      a[3] = 1   // update existing key
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

 - print(<object>) 
   Outputs the data to console.  