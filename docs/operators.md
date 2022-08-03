
# Operators

Operator are special symbols that has functional purpose to perform action on given data inputs/objects.

talion supports the following operators for its data types:

## Unary Operators

Operators that takes in one operand are called Unary Operators.

- Unary Not (!x): Negatives boolean value
- Unary Minus (-x): Denote negative of the given numeric value.

## Arithmetic Operators

These operators are used with integer objects to perform basic mathematical operations.
These are infix operator that takes in two operand values.

- Addition (+): To perform mathematical addition of two operators
- Subtraction (-): Subtracts the first operand value by the second operand value.
- Multiplication (*): Mathematical multiplication of two values.
- Integer Division (/): Divides first operand by the seond. This performs integer division.

## Conditional Operators

Operators used to compare magnitude of two operands. 

- Less Than <
- Less Than or Equal To <=
- Greater Than >
- Greater Than or Equal To >=
- Equal To ==
- Not Equal To !=

## Special Operators

These are operators symbols work for specific expression or data types.

- Assignment Operator (=): Used to bind identifier with an expression value.
- String Concatenation (+): Adds one string to end of another string and return the new string object.

## Example Usage:
```
a = 3
b = 4
c = a*a + b*b
c == 25                 // true
!true                   // false
!false                  // true
(1!=1) == false         // true
"a" + "b"               // "ab"
```