Program that takes source input text or tokens and generates Abstract Syntax Tree.
It is abstract because the output removed details like brackets/comma/etc to geenrate tree that conforms to the structure of program.

Parser Generator: Tool to generate Parser for given set of rules. Like YACC or Bison
They read Context-Free Grammar (CFG) text to generate parser. 

Talion uses Recursive Descent Parser - Top Down Parser.

# var Statements

Format: var <identifer> = <expressopm>;

Variable Binding:
var x = 5;
var y = 10;
var foobar = add(5, 5);
var barfoo = 5 * 5 / 10 + 18 - add(5, 5) + multiply(124);
var anotherName = barfoo;

Expression: Something that produces value
In the example: 5, 10, add(5,5)..

Statement: Don't produce value.
E.g: var a = 5;

