
import { verifyToken } from "../Services/authService";
import { Request, Response, NextFunction } from "express";

export const userAuthVerify = async (req: Request, res: Response, next: NextFunction) => {

    const tokenUser = req.cookies["auth-user"];
    if (!tokenUser) {

        return res.sendStatus(403);
    }

    try {

        const user = await verifyToken(tokenUser);
        req.user = user;
        next();
    } catch (err) {

        return res.sendStatus(403);
    }
};
