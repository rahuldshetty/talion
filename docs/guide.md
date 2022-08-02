
# Quick Start <!-- {docsify-ignore} -->

## Installation

Download talion binary from the Github release page: [alpha v1.0](https://github.com/rahuldshetty/talion/releases/tag/v1.0)

Run the interpreter to start using talion:

```bash
$ sh talion-linux-amd64
Hello USER! This is talion language!
You can now type commands
>>
```

## Say 'Hello World' in talion

Any coding example without "hello world" program is incomplete, so this is how you greet in talion:
```
>> greet = fn (name) { return "Hello " + name + "!" }
>> var name = "World"
>> print(greet(name))
Hello World!
```

## Factorial example

The above example was not really complicated. If you are still not conviced with this programming language, then I present to you RECURSION.

```
>> fact = fn(n){ if(n==1) { return 1 } else { return n * fact(n-1) } }
>> fact(10)
3628800
```
