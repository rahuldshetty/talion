
# Integer

Integers are numerical data type that is used to store mathematical integer values. Depending on the platform size the range of integer value that is supported in talion could be 32bit-signed or 64bit signed integer based on talion-binary/OS that you run.

- Under the hood, it is defined over int64 data type in Go using Integer object type.
- Integers are truthy values, that mean when referring to integer in if/block condition as expression the condition will always be true.

-  Integer support arithmetic operators

    Example:
    ```
    one = 11
    two = 22
    one * two // 242
    one + two // 33
    one - two // -11
    one / two // 0 - division returns integer part.

    ```

-  Integer can be used with conditional operators

    Example:
    ```
    one = 11
    two = 22
    one == one // true
    one == two // false
    one + one == two // true
    ```

- Integer expression can be used to access element of [List](/datatypes/list.md) or be used as Key/Value in a [HashMap](/datatypes/hash.md).