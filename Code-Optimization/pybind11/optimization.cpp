#include <pybind11/pybind11.h>

int times(int m, int n) {

    return m * n;
}

PYBIND11_MODULE(times, m) {
    m.def("times", &times, "This is a test function.");
}