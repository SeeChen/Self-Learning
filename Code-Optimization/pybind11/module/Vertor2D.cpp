#include "VectorMath.h"

#include <cmath>
#include <sstream>

Vector2D::Vector2D(
    double x, 
    double y
) : 
    x(x), y(y) 
{}

double Vector2D::magnitude() const {
    return std::sqrt(
        x * x + y * y
    );
}

Vector2D Vector2D::operator + (const Vector2D& other) const {
    return Vector2D(
        x + other.x,
        y + other.y
    );
}

Vector2D Vector2D::operator * (double factor) const {
    return Vector2D(
        x * factor,
        y * factor
    );
}

std::string Vector2D::toString() const {
    std::stringstream stream;
    stream << "Vector2D(" << x << ", " << y << ")";
    return stream.str();
}

Vector2D operator * (
    double factor,
    const Vector2D& vec
) {
    return Vector2D(
        vec.x * factor,
        vec.y * factor
    );
}

