#ifndef VERTOR_MATH_H
#define VERTOR_MATH_H

#include <string>

// Class of Vertor 2D
struct Vector2D {

    double x, y;

    // Constructor
    Vector2D(
        double x = 0,
        double y = 0
    );

    double magnitude () const;

    // Vector Addition
    Vector2D operator + (
        const Vector2D& other
    ) const;

    // Scalar multiplication
    Vector2D operator * (
        double factor
    ) const;

    std::string toString() const;
};

Vector2D operator * (
    double factor,
    const Vector2D& vec
);

#endif
