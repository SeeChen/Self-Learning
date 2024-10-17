
import jwt from 'jsonwebtoken';

import { Request, Response, NextFunction } from 'express';

export const authMiddleware = (req: any, res: any, next: NextFunction) => {

    const token = req.header['Authorization']?.split(' ')[1];

    if (!token) {

        return res.sendStatus(403);
    }

    jwt.verify(token, process.env.JWT_SECRET || "EiHeiHei", (err: any, user: any) => {

        if (err) {

            return res.sendStatus(403)
        }

        req.user = user;
        next();
    });
};
