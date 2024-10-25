
import { Request, Response } from 'express';
import { generateToken } from '../auth/authService';
import { serviceUser } from '../Services/serviceUser';

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

        res.cookie("auth", generateToken({

            "uid": user.id,
            "name": user.name,
            "pwd": user.password,
            "email": user.email,
            "com": user.userCompany ? user.userCompany.companyId : -1,
            "com_auth": user.userCompany ? user.userCompany.role : -1,
        }), {

            httpOnly: true,
            sameSite: 'strict',
            maxAge: 60 * 60 * 1000,
        })
        
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

    const uid = Number(req.uAuth.uid) || -1;
    const { id, email } = req.query;
    try {

        const user = await serviceUser.info(

            id ? Number(id) : undefined,
            email ? String(email) : undefined,
            Number(uid) === Number(id) ? [] : ["password", "userCompany"]
        );

        res.json(user);
    } catch (err) {

        res.status(403).send(err.message);
    }
};

export const userUpdateById = async (req: Request, res: Response) => {

    const id = req.uAuth.uid;
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

        res.cookie("auth", generateToken({

            "uid": newUser.id,
            "name": newUser.name,
            "pwd": newUser.password,
            "email": newUser.email,
            "com": newUser.userCompany ? newUser.userCompany.companyId : -1,
            "com_auth": newUser.userCompany ? newUser.userCompany.role : -1,
        }), {

            httpOnly: true,
            sameSite: 'strict',
            maxAge: 60 * 60 * 1000,
        });

        res.json(newUser);
    } catch (err) {

        res.sendStatus(401);
    }
};

export const userDeleteById = async (req: Request, res: Response) => {

    const id = req.uAuth.uid;
    const { name, email, password } = req.body;

    try {

        await serviceUser.delete(id, name, email, password);
        res.status(200).send("Delete Successful.");
    } catch (err) {

        res.status(401).send(err.message);
    }
};
