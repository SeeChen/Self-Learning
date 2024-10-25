
import { verifyPassword, verifyToken } from "../auth/authService";
import { Request, Response} from "express";
import { companyCreate, companyDelete, companyGetByKey, companyUpdate } from "../controllers/controllerCompany";

export const companyAuthVerify = {

    create: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.status(403);
        }

        try {

            const auth = await verifyToken(token);
            req.uAuth = auth;

            companyCreate(req, res);
        } catch (err) {

            return res.status(403);
        }
    },

    getInfo: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (token) {

            try {

                const auth = await verifyToken(token);
                req.uAuth = auth;
            } catch (err) {
    
                return res.status(403);
            }
        }

        companyGetByKey(req, res);
    },

    update: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.status(403);
        }

        try {

            const auth = await verifyToken(token);
            const { id } = req.query;

            if (Number(auth.com) !== Number(id) || Number(auth.com_auth) !== 0) {

                return res.status(403).send("Permission Denied.")
            }
            
            req.uAuth = auth;
            companyUpdate(req, res);
        } catch (err) {

            return res.status(403);
        }
    },

    delete: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.status(403);
        }

        try {

            const auth = await verifyToken(token);
            const { id } = req.query;
            const { adminEmail, adminPassword } = req.body;

            if (Number(auth.com) !== Number(id) || Number(auth.com_auth) !== 0) {

                return res.status(403).send("Permission Denied..")
            }

            if (adminEmail !== auth.email || !(await verifyPassword(adminPassword, auth.pwd))) {

                return res.status(403).send("Permission Denied...")
            }
            
            req.uAuth = auth;
            companyDelete(req, res);
        } catch (err) {

            return res.status(403);
        }
    },
}