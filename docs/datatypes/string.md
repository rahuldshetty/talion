
# String

Sequence of character enclosed by double quotes.

- String characters can be accessed using zero-indexed or negative-index integer value.
- Supports concatenation using addition operator.
- Can be used as a key/value in a HashMap.
- Length of the string can be calculated by using len() method.

- As of version alpha v1.0 there is no support for escape characters. This means that "\n" would be evaluated as "\n" but not a new line in the interpreter.

- Example Usage:
    ```
    h = "hello"
    h[0]                    // h
    w = "world"

    print(h + " " + w)      // "hello world"
    print(len(h))           // 5
    ```