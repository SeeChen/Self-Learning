"""Python Operators Demonstration

This module demonstrates various categories of Python operators, including:

- Arithmetic Operators: `+`, `-`, `*`, `/`, `%`, `**`, `//`
- Comparison Operators: `==`, `!=`, `>`, `<`, `>=`, `<=`
- Assignment Operators: `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `**=`, `//=`
- Bitwise Operators: `&`, `|`, `^`, `~`, `<<`, `>>`
- Logical Operators: `and`, `or`, `not`
- Membership Operators: `in`, `not in`
- Identity Operators: `is`, `is not`

Each operator category is encapsulated in a dedicated class with methods
that demonstrate usage examples.
"""


# ------------------------------------------------------------
# Arithmetic Operators
# ------------------------------------------------------------

class ArithmeticOperators:
    """Demonstrates Python's arithmetic operators."""

    @staticmethod
    def add(a: int, b: int) -> None:
        """Demonstrates the addition operator (`+`).

        Args:
            a (int): The first operand.
            b (int): The second operand.
        """
        print("Operator: + (Addition)")
        print(f"{a} + {b} = {a + b}\n")

    @staticmethod
    def sub(a: int, b: int) -> None:
        """Demonstrates the subtraction operator (`-`).

        Args:
            a (int): The first operand.
            b (int): The second operand.
        """
        print("Operator: - (Subtraction)")
        print(f"{a} - {b} = {a - b}\n")

    @staticmethod
    def mul(a: int, b: int) -> None:
        """Demonstrates the multiplication operator (`*`)."""
        print("Operator: * (Multiplication)")
        print(f"{a} * {b} = {a * b}\n")

    @staticmethod
    def div(a: int, b: int) -> None:
        """Demonstrates the division operator (`/`).

        Handles division by zero gracefully.
        """
        print("Operator: / (Division)")
        try:
            print(f"{a} / {b} = {a / b}\n")
        except ZeroDivisionError:
            print("Error: Division by zero!\n")

    @staticmethod
    def mod(a: int, b: int) -> None:
        """Demonstrates the modulus operator (`%`)."""
        print("Operator: % (Modulus)")
        print(f"{a} % {b} = {a % b}\n")

    @staticmethod
    def pow(a: int, b: int) -> None:
        """Demonstrates the exponentiation operator (`**`)."""
        print("Operator: ** (Exponentiation)")
        print(f"{a} ** {b} = {a ** b}\n")

    @staticmethod
    def divisible(a: int, b: int) -> None:
        """Demonstrates the floor division operator (`//`)."""
        print("Operator: // (Floor Division)")
        print(f"{a} // {b} = {a // b}\n")


# ------------------------------------------------------------
# Comparison Operators
# ------------------------------------------------------------

class ComparisonOperators:
    """Demonstrates Python's comparison operators."""

    @staticmethod
    def equals(obj1, obj2) -> None:
        """Demonstrates equality operator (`==`)."""
        print("Operator: == (Equal To)")
        print(f"{obj1} == {obj2} : {obj1 == obj2}\n")

    @staticmethod
    def not_equals(obj1, obj2) -> None:
        """Demonstrates inequality operator (`!=`)."""
        print("Operator: != (Not Equal To)")
        print(f"{obj1} != {obj2} : {obj1 != obj2}\n")

    @staticmethod
    def gt(obj1, obj2) -> None:
        """Demonstrates greater-than operator (`>`)."""
        print("Operator: > (Greater Than)")
        print(f"{obj1} > {obj2} : {obj1 > obj2}\n")

    @staticmethod
    def lt(obj1, obj2) -> None:
        """Demonstrates less-than operator (`<`)."""
        print("Operator: < (Less Than)")
        print(f"{obj1} < {obj2} : {obj1 < obj2}\n")

    @staticmethod
    def gt_equals(obj1, obj2) -> None:
        """Demonstrates greater-than-or-equal operator (`>=`)."""
        print("Operator: >= (Greater Than or Equal To)")
        print(f"{obj1} >= {obj2} : {obj1 >= obj2}\n")

    @staticmethod
    def lt_equals(obj1, obj2) -> None:
        """Demonstrates less-than-or-equal operator (`<=`)."""
        print("Operator: <= (Less Than or Equal To)")
        print(f"{obj1} <= {obj2} : {obj1 <= obj2}\n")


# ------------------------------------------------------------
# Assignment Operators
# ------------------------------------------------------------

class AssignmentOperators:
    """Demonstrates Python's assignment operators."""

    @staticmethod
    def equals() -> None:
        """Demonstrates simple assignment operator (`=`)."""
        print("Operator: = (Assignment)")
        val_int = 10
        val_str = "Python"
        print(f"val_int = {val_int}, val_str = '{val_str}'\n")

    @staticmethod
    def add_equals() -> None:
        """Demonstrates addition assignment operator (`+=`)."""
        print("Operator: += (Add and Assign)")
        a = 10
        a += 5
        print(f"a = 10 → a += 5 → a = {a}\n")

    @staticmethod
    def sub_equals() -> None:
        """Demonstrates subtraction assignment operator (`-=`)."""
        print("Operator: -= (Subtract and Assign)")
        a = 10
        a -= 5
        print(f"a = 10 → a -= 5 → a = {a}\n")

    @staticmethod
    def mul_equals() -> None:
        """Demonstrates multiplication assignment operator (`*=`)."""
        print("Operator: *= (Multiply and Assign)")
        a = 10
        a *= 5
        print(f"a = 10 → a *= 5 → a = {a}\n")

    @staticmethod
    def div_equals() -> None:
        """Demonstrates division assignment operator (`/=`)."""
        print("Operator: /= (Divide and Assign)")
        a = 10
        a /= 5
        print(f"a = 10 → a /= 5 → a = {a}\n")

    @staticmethod
    def mod_equals() -> None:
        """Demonstrates modulus assignment operator (`%=`)."""
        print("Operator: %= (Modulus and Assign)")
        a = 10
        a %= 3
        print(f"a = 10 → a %= 3 → a = {a}\n")

    @staticmethod
    def pow_equals() -> None:
        """Demonstrates exponentiation assignment operator (`**=`)."""
        print("Operator: **= (Exponentiate and Assign)")
        a = 2
        a **= 5
        print(f"a = 2 → a **= 5 → a = {a}\n")

    @staticmethod
    def divisible_equals() -> None:
        """Demonstrates floor division assignment operator (`//=`)."""
        print("Operator: //= (Floor Divide and Assign)")
        a = 10
        a //= 3
        print(f"a = 10 → a //= 3 → a = {a}\n")


# ------------------------------------------------------------
# Bitwise Operators
# ------------------------------------------------------------

class BitOperators:
    """Demonstrates Python's bitwise operators."""

    @staticmethod
    def and_op(a: int, b: int) -> None:
        """Demonstrates bitwise AND (`&`)."""
        print("Operator: & (Bitwise AND)")
        print(f"{a} & {b} = {a & b}\n")

    @staticmethod
    def or_op(a: int, b: int) -> None:
        """Demonstrates bitwise OR (`|`)."""
        print("Operator: | (Bitwise OR)")
        print(f"{a} | {b} = {a | b}\n")

    @staticmethod
    def xor_op(a: int, b: int) -> None:
        """Demonstrates bitwise XOR (`^`)."""
        print("Operator: ^ (Bitwise XOR)")
        print(f"{a} ^ {b} = {a ^ b}\n")

    @staticmethod
    def not_op(a: int) -> None:
        """Demonstrates bitwise NOT (`~`)."""
        print("Operator: ~ (Bitwise NOT)")
        print(f"~{a} = {~a}\n")

    @staticmethod
    def left_shift(a: int, n: int) -> None:
        """Demonstrates left shift (`<<`)."""
        print("Operator: << (Left Shift)")
        print(f"{a} << {n} = {a << n}\n")

    @staticmethod
    def right_shift(a: int, n: int) -> None:
        """Demonstrates right shift (`>>`)."""
        print("Operator: >> (Right Shift)")
        print(f"{a} >> {n} = {a >> n}\n")


# ------------------------------------------------------------
# Logical Operators
# ------------------------------------------------------------

class LogicOperators:
    """Demonstrates Python's logical operators."""

    @staticmethod
    def logic_and(a: bool, b: bool) -> None:
        """Demonstrates logical AND operator (`and`)."""
        print("Operator: and (Logical AND)")
        print(f"{a} and {b} = {a and b}\n")

    @staticmethod
    def logic_or(a: bool, b: bool) -> None:
        """Demonstrates logical OR operator (`or`)."""
        print("Operator: or (Logical OR)")
        print(f"{a} or {b} = {a or b}\n")

    @staticmethod
    def logic_not(a: bool) -> None:
        """Demonstrates logical NOT operator (`not`)."""
        print("Operator: not (Logical NOT)")
        print(f"not {a} = {not a}\n")


# ------------------------------------------------------------
# Membership Operators
# ------------------------------------------------------------

class MembersOperators:
    """Demonstrates Python's membership operators."""

    @staticmethod
    def in_op(element, container) -> None:
        """Demonstrates `in` operator."""
        print("Operator: in (Membership)")
        print(f"{element} in {container} : {element in container}\n")

    @staticmethod
    def not_in_op(element, container) -> None:
        """Demonstrates `not in` operator."""
        print("Operator: not in (Membership)")
        print(f"{element} not in {container} : {element not in container}\n")


# ------------------------------------------------------------
# Identity Operators
# ------------------------------------------------------------

class IdentityOperators:
    """Demonstrates Python's identity operators."""

    @staticmethod
    def is_op(obj1, obj2) -> None:
        """Demonstrates `is` operator."""
        print("Operator: is (Identity)")
        print(f"{obj1} is {obj2} : {obj1 is obj2}\n")

    @staticmethod
    def is_not_op(obj1, obj2) -> None:
        """Demonstrates `is not` operator."""
        print("Operator: is not (Identity)")
        print(f"{obj1} is not {obj2} : {obj1 is not obj2}\n")


# ------------------------------------------------------------
# Main Demonstration
# ------------------------------------------------------------

if __name__ == "__main__":
    print("=== Arithmetic Operators ===")
    a, b = 10, 5
    ArithmeticOperators.add(a, b)
    ArithmeticOperators.sub(a, b)
    ArithmeticOperators.mul(a, b)
    ArithmeticOperators.div(a, b)
    ArithmeticOperators.mod(a, b)
    ArithmeticOperators.pow(a, b)
    ArithmeticOperators.divisible(a, b)

    print("=== Comparison Operators ===")
    ComparisonOperators.equals(a, b)
    ComparisonOperators.not_equals(a, b)
    ComparisonOperators.gt(a, b)
    ComparisonOperators.lt(a, b)
    ComparisonOperators.gt_equals(a, b)
    ComparisonOperators.lt_equals(a, b)

    print("=== Assignment Operators ===")
    AssignmentOperators.equals()
    AssignmentOperators.add_equals()
    AssignmentOperators.sub_equals()
    AssignmentOperators.mul_equals()
    AssignmentOperators.div_equals()
    AssignmentOperators.mod_equals()
    AssignmentOperators.pow_equals()
    AssignmentOperators.divisible_equals()

    print("=== Bitwise Operators ===")
    BitOperators.and_op(a, b)
    BitOperators.or_op(a, b)
    BitOperators.xor_op(a, b)
    BitOperators.not_op(a)
    BitOperators.left_shift(a, 2)
    BitOperators.right_shift(a, 2)

    print("=== Logical Operators ===")
    LogicOperators.logic_and(True, False)
    LogicOperators.logic_or(True, False)
    LogicOperators.logic_not(True)

    print("=== Membership Operators ===")
    MembersOperators.in_op(3, [1, 2, 3, 4])
    MembersOperators.not_in_op('x', 'Python')

    print("=== Identity Operators ===")
    x = [1, 2]
    y = x
    z = [1, 2]
    IdentityOperators.is_op(x, y)
    IdentityOperators.is_not_op(x, z)
