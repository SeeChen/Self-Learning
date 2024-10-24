
import { PrismaClient } from '@prisma/client';
import { Request, Response } from 'express';
import { generateToken } from '../Services/authService';
import { serviceUser } from '../Services/serviceUser';

const prisma = new PrismaClient();

export const userCreate = async (req: Request, res: Response ) => {

    const { name, email, password, address } = req.body;
    try {

        const user = await serviceUser.create (

            name,
            email,
            password,
            address
        );

        res.json(await serviceUser.info(user.id, undefined, ["password"]));
    } catch (err) {

        res.status(409).send(err.message);
    }
};

export const userLogin = async (req: Request, res: Response) => {

    const { email, password } = req.body;
    try {

        const user = await serviceUser.login(email, password);

        res.cookie("auth-user", generateToken(user.id), {

            httpOnly: true,
            sameSite: 'strict',
            maxAge: 60 * 60 * 1000,
        });

        res.cookie("uid", user.id, {

            httpOnly: true,
            sameSite: 'strict',
            maxAge: 60 * 60 * 1000,
        });

        
        res.json({

            message: "Login Successful",
            data: await serviceUser.info(
                user.id,
                user.email,
                ["password", "userCompany"]
            )
        });

    } catch (err) {

        res.status(401).send(err.message);
    }
};

export const userGetByKey = async (req: Request, res: Response) => {

    const uid = req.cookies["uid"];
    const { id, email } = req.query;
    try {

        const user = await serviceUser.info(

            id ? Number(id) : undefined,
            email ? String(email) : undefined,
            uid === id ? [] : ["password", "userCompany"]
        );

        res.json(user);
    } catch (err) {

        res.status(403).send(err.message);
    }
};

export const userUpdateById = async (req: Request, res: Response) => {

    const id = req.cookies["uid"];
    const { name, email, passwordOld, passwordNew, address } = req.body;

    try {

        const newUser = await serviceUser.update(

            Number(id),
            name,
            email,
            passwordOld,
            passwordNew,
            address
        );

        res.json(newUser);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

export const userDeleteById = async (req: Request, res: Response) => {

    const id = req.cookies["uid"];
    const { name, email, password } = req.body;

    try {

        await serviceUser.delete(id, name, email, password);
        res.status(200).send("Delete Successful.");
    } catch (err) {

        res.status(401).send(err.message);
    }
};
