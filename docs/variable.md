
# Variables

Just like any other programming languages, talion supports naming your data objects through identifers. talion follows dynamic typing when it comes to assigning data type to variables. 

- Identifiers are the names associtated with variable in memory. Identifier should alwasy start with alphabetical characters followed by alphabets/numerica/underscore. 
- Variable binding to identifier can be done with VARiable keyword or just by directly providing the identifer name. 
    Syntax:
    ```
    VAR <identifier> = <value>
    <identifer> = <value>
    ```
- Identifer values can be passed as parameter to function calls, or used in dictonary or list items.
- Example Usage:
    ```
    a = 10
    print(a+10)                     // 20
    a = a * a                       // 100
    var b = 123                     
    b = b + a                       // b = 223
    ```