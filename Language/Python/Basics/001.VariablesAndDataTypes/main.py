"""Python Variables and Basic Data Types

This module demonstrates how variables work in Python, how to assign values,
and how to use basic built-in data types such as:
- int
- float
- complex
- str
- bool
- NoneType

It also demonstrates:
- Multiple assignment
- Type checking with type() and isinstance()
- Type conversion
- String indexing, slicing, and escape sequences
- The relationship between bool and int
- Constants and dynamic typing
"""


# ------------------------------------------------------------
# Constants (best practice: uppercase variable names)
# ------------------------------------------------------------
PI: float = 3.14159
E: float = 2.71828


def demo_variables() -> None:
    """Demonstrate Python variable assignment and basic data types."""
    number: int = 1
    double: float = 1.0
    string: str = "string"

    print(number)
    print(double)
    print(string)

    # Multiple assignment: assign the same value to multiple variables
    number1 = number2 = number3 = number

    # Multiple assignment: assign different values to different variables
    double1, integer1, string1 = 1.414, 5, "Str"

    # Using type() to check variable types
    print(type(double1))
    print(type(integer1))
    print(type(string1))

    # Boolean data type
    is_active: bool = True
    print(type(is_active))

    # Python standard built-in data types include:
    # Number, String, Boolean, List, Tuple, Set, Dictionary


def demo_numbers() -> None:
    """Demonstrate Python number types and isinstance() usage."""
    a: int = 748
    b: float = 520.1314
    c: complex = 4 + 3j

    print(type(a))
    print(type(b))
    print(type(c))

    # Use isinstance() to check a variable's data type
    is_number: bool = isinstance(a, int)
    print(f"The variable 'a' ({a}) is of type int: {is_number}")

    # Access complex number attributes
    print(f"Complex number {c} -> Real part: {c.real}, Imag part: {c.imag}")

    # Boolean is a subclass of int (True == 1, False == 0)
    print(f"True + True = {True + True}")
    print(f"bool is subclass of int: {issubclass(bool, int)}")


def demo_none() -> None:
    """Demonstrate the NoneType object."""
    empty_value: None | int = None
    print(f"The variable empty_value has value: {empty_value}")
    print(f"Type of empty_value: {type(empty_value)}")

    # Assign a value later (dynamic typing)
    empty_value = 42
    print(f"Now empty_value = {empty_value}, type: {type(empty_value)}")


def demo_type_conversion() -> None:
    """Demonstrate explicit type conversions in Python."""
    num_str: str = "3.14"
    num_float: float = float(num_str)
    num_int: int = int(5.9)
    num_str_again: str = str(100)
    bool_value: bool = bool(0)  # False because 0 is falsy

    print(f"Convert '3.14' -> float: {num_float}")
    print(f"Convert 5.9 -> int: {num_int}")
    print(f"Convert 100 -> str: '{num_str_again}'")
    print(f"Convert 0 -> bool: {bool_value}")


def demo_strings() -> None:
    """Demonstrate string operations, indexing, slicing, and escape sequences."""
    my_str: str = "This is a String!"

    # Indexing (starts at 0)
    print(my_str[0])  # Prints first character: T

    # Negative indexing (starts at -1)
    print(my_str[-1])  # Prints last character: !

    # String slicing: [start_index : end_index]
    print(my_str[1:5])  # Prints characters from index 1 to 4

    # Slice from the start
    print(my_str[:5])
    print(my_str[0:5])

    # Slice to the end
    print(my_str[5:])

    # Copy the entire string
    print(my_str[:])

    # Escape sequences
    new_line_str: str = "Line 1\nLine 2"
    tab_str: str = "A\tB\tC"
    quote_str: str = "He said: \"Hello!\""
    print(new_line_str)
    print(tab_str)
    print(quote_str)

    # Raw string: ignores escape sequences
    raw_str: str = r"This is a raw string: \n won't create a new line"
    print(raw_str)

    # Multiline string
    multiline_str: str = """This is
a multiline
string."""
    print(multiline_str)


def demo_dynamic_typing() -> None:
    """Demonstrate dynamic typing in Python."""
    value: int | str = 100
    print(f"value = {value}, type = {type(value)}")

    # Python allows reassigning variables with a new type (not recommended in large codebases)
    value = "Now I am a string"
    print(f"value = {value}, type = {type(value)}")


def demo_constants() -> None:
    """Showcase constants and best practices."""
    print(f"Constant PI = {PI}")
    print(f"Constant E = {E}")


def main() -> None:
    """Entry point: Run all demonstrations."""
    demo_variables()
    demo_numbers()
    demo_none()
    demo_type_conversion()
    demo_strings()
    demo_dynamic_typing()
    demo_constants()


if __name__ == "__main__":
    main()
