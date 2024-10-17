
import jwt from 'jsonwebtoken';
import bcrypt from 'bcryptjs';

import { PrismaClient } from '@prisma/client';
import { Request, Response } from 'express';

const prisma = new PrismaClient();

export const userCreate = async (req: Request, res: Response ) => {

    const { name, email, password, address } = req.body;
    console.log(req);

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

    const { id } = req.params;
    const { name, email, password, address } = req.body;

    const hashedPassword = password ? await bcrypt.hash(password, 10) : undefined;

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

    const { id } = req.params;

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

    res.json({ token });
};
