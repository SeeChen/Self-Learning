
import { Request, Response } from "express";
import { serviceCompany } from "../Services/serviceCompany";

export const companyCreate = async (req: Request, res: Response) => {

    const creatorId = req.uAuth.uid;
    const { name, address } = req.body;
    try {

        const company = await serviceCompany.create (
            Number(creatorId),
            name,
            address
        );
        res.json(company);
    } catch (err) {

        res.status(400).send(err.message);
    }
};

export const companyGetByKey = async (req: Request, res: Response) => {

    const UID = req.uAuth.uid;
    const { id, name } = req.query;
    try {

        const company = await serviceCompany.info(

            UID ? Number(UID) : undefined,
            id ? Number(id) : undefined,
            name ? String(name) : undefined
        );

        res.json(company);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

export const companyUpdate = async (req: Request, res: Response) => {

    const adminId = req.uAuth.uid;
    const { id } = req.query;
    const { name, address } = req.body;
    try {
        
        const company = await serviceCompany.update(

            Number(adminId),
            Number(id),
            name,
            address
        );

        res.json(company);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

export const companyDelete = async (req: Request, res: Response) => {

    const uid = req.uAuth.uid;
    const { id } = req.query;
    const { companyName, adminEmail, adminPassword } = req.body;

    try {

        await serviceCompany.delete(Number(uid), String(adminEmail), String(adminPassword), Number(id), String(companyName));
        res.status(200).send("Delete Successful.");
    } catch (err) {

        res.status(401).send(err.message);
    }
};
