# -*- coding: utf-8 -*-
"""
A simple Python environment check script.

This script verifies that Python is correctly installed and running.
It prints a greeting message, the current Python version, and
basic system information such as the platform and executable path.
"""

import sys
import platform


def show_environment_info():
    """Print a greeting message and environment information.

    This function serves as a basic test to confirm that Python is properly
    installed and running on the user's system. It displays a welcome message,
    the current Python version, and the operating system details.

    Returns:
        None
    """
    print("Hello Python!")
    print("-" * 40)
    print(f"Python Version : {sys.version}")
    print(f"Executable Path: {sys.executable}")
    print(f"Platform       : {platform.system()} {platform.release()}")
    print(f"Architecture   : {platform.architecture()[0]}")
    print(f"Python Build   : {platform.python_build()}")
    print("-" * 40)


if __name__ == "__main__":
    show_environment_info()
