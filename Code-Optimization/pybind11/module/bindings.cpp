#include <pybind11/pybind11.h>

namespace py = pybind11;

int testFunction(void);

PYBIND11_MODULE(myModule, m) {
    m.def("my_test", &testFunction, "This is My Test function.");
}