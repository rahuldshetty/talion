
# Hash

Hash is used to store unordered collection of elements as key/value pair. Data type is similar to dictionary or map object in other programming languages where there is a one-to-one mapping between a key object and value.

- Following syntax is used to create hash map:
    ```
    { <hashable-key1>:<value1>, <hashable-key2>:<value2> ... }
    ```
- To access a given value for a key or to update the value, we can use the following annotations:
    ```
    a = {1:1, 2:"two"}
    a[0]                // null
    a[1]                // 1
    a[2] = 2            // Value is updated from "two" to 2
    a[2]                // 2
    ```
    If the key element is not present in the hash, it will be newly created.

- For accessing a given key element in hash it takes constant time.
- The data type for a key object should be Hashable type - String, Boolean, Integer data types.


- Example Usage:
  ```
  a = {0: fn(x){x+100}, 1: 2}
  a[0](10)                      // 110
  len(a)                        // 2
  ```