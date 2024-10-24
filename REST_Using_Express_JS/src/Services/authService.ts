
import jwt from "jsonwebtoken";
import bcrypt from "bcryptjs";

export const hashPassword = async (password: string): Promise<string> => {

    return await bcrypt.hash(password, 10);
};

export const verifyPassword = async (password: string, hashedPassword: string): Promise<boolean> => {

    return await bcrypt.compare(password, hashedPassword);
};

export const generateToken = (userId: number): string => {

    return jwt.sign({ userId }, process.env.JWT_SECRET || "EiHeiHei", { expiresIn: "1h" });
};

export const verifyToken = (token: string): any => {

    return jwt.verify(token, process.env.JWT_SECRET || "EiHeiHei");
};
