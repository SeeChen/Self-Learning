from setuptools import setup, Extension
import pybind11

setup(
    name='times',
    ext_modules=[Extension("times", ["optimization.cpp"], include_dirs=[pybind11.get_include()])],
    setup_requires=["pybind11"],
)