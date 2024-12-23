
import { Request, Response } from "express";
import { serviceCompanyManage } from "../Services/serviceCompanyManage";

export const companyManageAssign = async (req: Request, res: Response) => {

    const adminId = req.uAuth.uid;
    const companyId = req.uAuth.com;
    const { userId, role } = req.body;

    try {

        const assign = await serviceCompanyManage.assign(

            Number(adminId),
            Number(companyId),
            Number(userId),
            Number(role)
        );

        res.json(assign);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

export const companyManageRemove = async (req: Request, res: Response) => {

    const adminId = req.uAuth.uid;
    const companyId= req.uAuth.com;
    const { userId } = req.body;

    try {

        await serviceCompanyManage.remove(
            
            Number(adminId), 
            Number(companyId), 
            Number(userId)
        );
        res.status(200).send("Remove Successful.");
    } catch (err) {

        res.status(401).send(err.message);
    }
};

export const companyManageModifyRole = async (req: Request, res: Response) => {

    const adminId = req.uAuth.uid;
    const companyId= req.uAuth.com;
    const { userId, roleNew } = req.body;

    try {

        const userWithCompany = await serviceCompanyManage.modifyRole(

            Number(adminId),
            Number(companyId),
            Number(userId),
            Number(roleNew)
        );

        res.json(userWithCompany);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

