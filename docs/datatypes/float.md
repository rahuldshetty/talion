
# Float

Numeric data type that is used to store values with fractional or decimal values. They follow standard convention where the integer part is separated by decimal part by a decimal point (or dot).

- Under the hood, it is defined over float64 data type in Go using Float object type.
- Similar to integer datatype, floating point numbers support arithmetic & conditional operations.
- When performing arithmetic operator with a integer datatype, the integer value is first converted to floating point number (i.e 4 -> 4.0) and then operator action is performed. Returned value of these expression will always be of type Float

-   Example:
    ```
    one = 11
    two = 22
    one * two // 242
    one + two // 33
    one - two // -11
    one / two // 0 - division returns integer part.

    ```

- Floating point numbers can be used as key/value objects while defining HashMap. 
  Example:
  ```
  mp = {1.1: 1, 2.2: 2}
  mp[1.1] + mp[2.2]             // 3
  ```