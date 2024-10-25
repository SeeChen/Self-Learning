
import { verifyToken } from "../auth/authService";
import { Request, Response, NextFunction } from "express";

export const companyAuthVerify = async (req: Request, res: Response, next: NextFunction) => {

    const tokenCompany = req.cookies["auth-company"];
    if (!tokenCompany) {

        return res.sendStatus(403);
    }

    try {

        const user = await verifyToken(tokenCompany);
        console.log(user.keyNumber);
        req.user = user;
        next();
    } catch (err) {

        return res.sendStatus(403);
    }
};
