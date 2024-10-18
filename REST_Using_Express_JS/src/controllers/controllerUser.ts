
import jwt from 'jsonwebtoken';
import bcrypt from 'bcryptjs';

import { PrismaClient } from '@prisma/client';
import { Request, Response } from 'express';

const prisma = new PrismaClient();

export const userCreate = async (req: Request, res: Response ) => {

    const { name, email, password, address } = req.body;

    const userExist = await prisma.user.findUnique({ where: { email } });
    if (userExist) {

        return res.status(409).json({ message: "User Exist." });
    }

    const hashedPassword = await bcrypt.hash(password, 10);

    const user = await prisma.user.create({

        data: {

            name,
            email,
            password: hashedPassword,
            address
        }
    });

    res.json(user);
};

export const userGetById = async (req: Request, res: Response) => {

    const { id } = req.params;
    const user = await prisma.user.findUnique({ where: { id: Number(id) } });

    if (!user) {

        return res.status(404).send("User Not Found");
    }

    res.json(user);
};

export const userUpdateById = async (req: Request, res: Response) => {

    const id = req.cookies["uid"];
    const { name, email, passwordOld, passwordNew, address } = req.body;

    const userOriginal = await prisma.user.findUnique({ where: { id: Number(id) } });
    if (!(await bcrypt.compare(passwordOld, userOriginal.password))) {

        return res.status(401).send("Invalid Password");
    }

    const hashedPassword = passwordNew ? await bcrypt.hash(passwordNew, 10) : passwordOld;

    const user = await prisma.user.update({

        where: { id: Number(id) },
        data: {

            name,
            email,
            password: hashedPassword,
            address
        }
    });

    res.json(user);
};

export const userDeleteById = async (req: Request, res: Response) => {

    const id = req.cookies["uid"];
    const { name, email, password } = req.body;

    const user = await prisma.user.findUnique({ where: { id: Number(id) } });
    if (!(await bcrypt.compare(password, user.password))) {

        return res.status(401).send("Invalid Password");
    }

    if (user.name != name) {

        return res.status(401).send("Incorrect Name.");
    }
    if (user.email != email) {

        return res.status(401).send("Incorrect Email.")
    }

    await prisma.user.delete({ where: { id: Number(id) } });
    
    res.sendStatus(204);
  };

export const userLogin = async (req: Request, res: Response) => {

    const { email, password } = req.body;
    const user = await prisma.user.findUnique({ where: { email } });

    if (!user || !(await bcrypt.compare(password, user.password))) {
        
        return res.status(401).send("Invalid credentials");
    }

    const token = jwt.sign({ userId: user.id }, process.env.JWT_SECRET || 'EiHeiHei', { expiresIn: '1h' });

    res.cookie("auth", token, {

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
        user: {

            id: user.id,
            name: user.name,
            email: user.email
        }
    });
};
