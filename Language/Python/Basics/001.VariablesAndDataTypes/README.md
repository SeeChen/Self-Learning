<div align=center>

# Chapter 1: Variables and Basic Data Types

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **Python variables** and **Basic  built-in data types**

Topic include:
- Variable assignment
- Multiple assignment
- Primitive data types: `int`, `float`, `complex`, `str`, `bool`, `NoneType`
- Type checking with `type()` and `isinstance()`
- Explicit type conversion
- String indexing, slicing, and escape sequences
- Relationship between `bool` and `int`
- Constants and dynamic typing

## 2. Technology
- **Language**: Python 3.11
- **IDE**: VS Code
- **OS**: Linux / Windows
- **Packages Used**:
  - Built-in types

## 3. Key Concepts
| Concept | Description |
|---------|----------------|
| **Variable** | Reference to stored data |
| **Dynamic Typing** | Type changes during runtime |
| **Constants** | Uppercase naming convention |
| **Number Types** | `int`, `float`, `complex` |
| **Boolean** | `True` or `False`, subclass of int |
| **NoneType** | Represents no value |
| **Type Conversion** | `int()`, `float()`, `str()`, `bool()` |
| **String Indexing & Slicing** | `"Hello"[0:3]` etc. |
| **Escape Sequences** | `\n`, `\t`, `\"` |

## 4. Run & Output
Run the program using:
```bash
python main.py
```

Expected output:
```bash
1
1.0
string
<class 'float'>
<class 'int'>
<class 'str'>
<class 'bool'>
<class 'int'>
<class 'float'>
<class 'complex'>
The variable 'a' (748) is of type int: True
Complex number (4+3j) -> Real part: 4.0, Imag part: 3.0
True + True = 2
bool is subclass of int: True
The variable empty_value has value: None
Type of empty_value: <class 'NoneType'>
Now empty_value = 42, type: <class 'int'>
Convert '3.14' -> float: 3.14
Convert 5.9 -> int: 5
Convert 100 -> str: '100'
Convert 0 -> bool: False
T
!
his
This
This
is a String!
Line 1
Line 2
A       B       C
He said: "Hello!"
This is a raw string: \n won't create a new line
This is
a multiline
string.
value = 100, type = <class 'int'>
value = Now I am a string, type = <class 'str'>
Constant PI = 3.14159
Constant E = 2.71828
```

## 5. Practice
1. Change the value of constants and observe results.
2. Demonstrate type conversion errors with invalid input.
3. Create your own multi-line string with formatting.
4. Verify more `issubclass` relationships.

## 6. References
1. [Python Official Documentation](https://docs.python.org/3/)
2. [Built-in Types](https://docs.python.org/3/library/stdtypes.html)
3. [Variables in Python](https://wiki.python.org/moin/Variables)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 27-OCT-2025 23:03 UTC +08:00*
</div>