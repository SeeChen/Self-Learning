<div align=center>

# Chapter 2: Operators in Python

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **Python operators** across multiple categories, explaining their syntax, semantics, and use cases.  
Each operator type is encapsulated in a dedicated class that prints demonstration results.

Covered categories:
- Arithmetic Operators   
- Comparison Operators   
- Assignment Operators   
- Bitwise Operators   
- Logical Operators   
- Membership Operators   
- Identity Operators 

## 2. Technology
- **Language**: Python 3.11
- **IDE**: VS Code
- **OS**: Linux / Windows
- **Packages Used**:
  - Built-in types

## 3. Key Concepts
| Category | Operators | Description |
|----------------|------------------|------------------|
| **Arithmetic** | `+`, `-`, `*`, `/`, `%`, `**`, `//` | Basic mathematical operations |
| **Comparison** | `==`, `!=`, `>`, `<`, `>=`, `<=` | Comparing values for size or equality |
| **Assignment** | `=`, `+=`, `-=`, `*=`, `/=`, `%=` | Assigning calculation results |
| **Bitwise** | `&`, `|`, `^`, `~`, `<<`, `>>` | Processing integers in binary form |
| **Logical** | `and`, `or`, `not` | Boolean logic-based control judgments |
| **Membership** | `in`, `not in` | Determining if an element belongs to a container |
| **Identity** | `is`, `is not` | Determining if two objects are the same reference |

## 4. Run & Output
Run the program using:
```bash
python main.py
```

Expected output:
```bash
=== Arithmetic Operators ===
Operator: + (Addition)
10 + 5 = 15

Operator: - (Subtraction)
10 - 5 = 5

Operator: * (Multiplication)
10 * 5 = 50

Operator: / (Division)
10 / 5 = 2.0

Operator: % (Modulus)
10 % 5 = 0

Operator: ** (Exponentiation)
10 ** 5 = 100000

Operator: // (Floor Division)
10 // 5 = 2

=== Comparison Operators ===
Operator: == (Equal To)
10 == 5 : False

Operator: != (Not Equal To)
10 != 5 : True

Operator: > (Greater Than)
10 > 5 : True

Operator: < (Less Than)
10 < 5 : False

Operator: >= (Greater Than or Equal To)
10 >= 5 : True

Operator: <= (Less Than or Equal To)
10 <= 5 : False

=== Assignment Operators ===
Operator: = (Assignment)
val_int = 10, val_str = 'Python'

Operator: += (Add and Assign)
a = 10 → a += 5 → a = 15

Operator: -= (Subtract and Assign)
a = 10 → a -= 5 → a = 5

Operator: *= (Multiply and Assign)
a = 10 → a *= 5 → a = 50

Operator: /= (Divide and Assign)
a = 10 → a /= 5 → a = 2.0

Operator: %= (Modulus and Assign)
a = 10 → a %= 3 → a = 1

Operator: **= (Exponentiate and Assign)
a = 2 → a **= 5 → a = 32

Operator: //= (Floor Divide and Assign)
a = 10 → a //= 3 → a = 3

=== Bitwise Operators ===
Operator: & (Bitwise AND)
10 & 5 = 0

Operator: | (Bitwise OR)
10 | 5 = 15

Operator: ^ (Bitwise XOR)
10 ^ 5 = 15

Operator: ~ (Bitwise NOT)
~10 = -11

Operator: << (Left Shift)
10 << 2 = 40

Operator: >> (Right Shift)
10 >> 2 = 2

=== Logical Operators ===
Operator: and (Logical AND)
True and False = False

Operator: or (Logical OR)
True or False = True

Operator: not (Logical NOT)
not True = False

=== Membership Operators ===
Operator: in (Membership)
3 in [1, 2, 3, 4] : True

Operator: not in (Membership)
x not in Python : True

=== Identity Operators ===
Operator: is (Identity)
[1, 2] is [1, 2] : True

Operator: is not (Identity)
[1, 2] is not [1, 2] : True
```

## 5. Practice
1. Modify operands `a` and `b` to observe different results.
2. Add a new operator demonstration for floor division with negative numbers.
3. Explore how `is` differes from `==` in object comparison.

## 6. References
1. [Python Official Documentation](https://docs.python.org/3/)
2. [Python Operator Precedence](https://docs.python.org/3/reference/expressions.html#operator-precedence)
3. [Built-In Types and Operators](https://docs.python.org/3/library/stdtypes.html)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-Nov-2025 23:03 UTC +08:00*
</div>