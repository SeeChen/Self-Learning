from setuptools import setup, Extension
import pybind11

ext_modules = [
    Extension(
        "myModule",
        [
            "module/bindings.cpp",
            "module/myTestFunc.cpp"
        ],
        include_dirs=[pybind11.get_include()],
        language="C++"
    )
]

setup(
    name='myModule',
    ext_modules=ext_modules,
    setup_requires=["pybind11"],
)