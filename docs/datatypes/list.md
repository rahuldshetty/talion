
# List

List data type are used to store collection of other data type elements. Unlike arrays, they support storing any type of elements in a single list.

- Elements in list can be accessed using zero-indexed or negative-indexed integer value.
- Supports len() method to count the number of elements present in the list.
- Supports push() method to insert new element in to the list.
- Lists are mutable data type in talion, i.e elements pushed will update to the original object in memory.

- Example Usage:
    ```
    a=[1, 2, fn(x){x+1}, 4]
    a[0]                            // 1
    a[-1]                           // 4
    push(a, a[2](2) )               // Run the method definition present in 2nd index of list and push the value into the same list
    a[3]                            // 3
    len(a)                          // 5
    ```