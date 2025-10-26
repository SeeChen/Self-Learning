<div align=center>

# Chapter 0: Hello Python!

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This is the initial step in the **Python Environment Validation Journey**.
It demonstrates:
- Checking if Python runs correctly.
- Printing essential environment information including:
    - Python Version
    - Executable Path
    - Operating System Details

## 2. Technology
- **Language**: Python 3.11
- **IDE**: VS Code
- **OS**: Linux / Windows
- **Packages Used**:
  - `sys` (system-specific parameters and functions)
  - `platform` (system environment querying)

## 3. Key Concepts
- **Entry Check**:
    - `if __name__ == "__main__": ...`
- **Environment Information**:
    - `sys.version` -> Python version  
    - `sys.executable` -> Python executable path  
    - `platform.system()` / `platform.release()` -> OS info  
    - `platform.architecture()` -> CPU architecture  
    - `platform.python_build()` -> Python build info  

## 4. Run & Output
Run the program using:
```bash
python main.py
```

Expected output:
```bash
Hello Python!
----------------------------------------
Python Version : 3.x.x (xxxx...) 
Executable Path: /path/to/python
Platform       : Windows/Linux/macOS 10/11/...
Architecture   : 64bit
Python Build   : ('default', 'xxx-xx-xxxx')
----------------------------------------
```

## 5. Practice
1. Modify the greeting message.
2. Print additional system details from the `platform` module.
3. Save the environment information to a text file.

## 6. References
1. [Python Official Documentation](https://docs.python.org/3/)
2. [Platform Module Docs](https://docs.python.org/3/library/platform.html)
3. [sys Module Docs](https://docs.python.org/3/library/sys.html)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 26-OCT-2025 19:03 UTC +08:00*
</div>