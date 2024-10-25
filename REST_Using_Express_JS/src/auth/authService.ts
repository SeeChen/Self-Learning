
import jwt from "jsonwebtoken";
import bcrypt from "bcryptjs";

interface TokenData {

    [key: string]: string | number | boolean | object;
}

export const hashPassword = async (password: string): Promise<string> => {

    return await bcrypt.hash(password, 10);
};

export const verifyPassword = async (password: string, hashedPassword: string): Promise<boolean> => {

    return await bcrypt.compare(password, hashedPassword);
};

export const generateToken = (data: TokenData): string => {

    if (!data || Object.keys(data).length === 0) {

        throw new Error("Data object must not be empty.");
    }

    const payload: { [key: string]: any } = {};
    for (const key in data) {

        if (data.hasOwnProperty(key)) {

            const value = data[key];

            if (typeof value === 'string' || typeof value === 'number' || typeof value === 'boolean') {

                payload[key] = value;
            } else {

                payload[key] = JSON.stringify(value);
            }
        }
    }

    return jwt.sign(payload, process.env.JWT_SECRET || "EiHeiHei", { expiresIn: "1h" });
};

export const verifyToken = (token: string): Promise<string> => {

    return new Promise((resolve, reject) => {

        jwt.verify(token, process.env.JWT_SECRET || "EiHeiHei", (err: any, decoded: any) => {

            if (err) {

                return reject(err);
            }

            resolve(decoded);
        });
    });
};
