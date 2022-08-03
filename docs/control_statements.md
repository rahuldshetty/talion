
# Control Statements

Control statement are provided in programming language to control the flow of execution.

talion supports one control statement structure:

## If/Else Expression

Statement is used to take a branch based on condition statement. Optionally you can provide alternative branch to take when the condition is not satisfied.

- Syntax:
    ```
    if (condition){
        <statements>
    } else {
        <alternate_statements>
    }
    ```
- Else block is optional.
- Condition should evaluate to Truthy values. Truthy values in talion are non-null and true values which evaluate to some definite value.

- Example:
    ```
    large_value = 10000000
    small_value = 1
    cond = true
    if(large_value>small_value){ cond = false }
    print(cond)                                             // false
    ```