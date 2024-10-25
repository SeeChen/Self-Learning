
import { verifyPassword, verifyToken } from "../auth/authService";
import { Request, Response} from "express";
import { companyManageAssign, companyManageModifyRole, companyManageRemove } from "../controllers/controllerCompanyManage";

export const companyManageAuthVerify = {

    assign: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.status(403);
        }

        try {

            const auth = await verifyToken(token);
            const { userId } = req.body;
            
            if (Number(auth.com_auth) !== 0 || Number(userId) == Number(auth.uid)) {

                return res.status(403).send("Permission Denied.");
            }

            req.uAuth = auth;
            companyManageAssign(req, res);
        } catch (err) {

            return res.status(403);
        }
    },

    remove: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.status(403);
        }

        try {

            const auth = await verifyToken(token);
            const { userId } = req.body;

            if (Number(auth.com_auth) !== 0 || Number(userId) == Number(auth.uid)) {

                return res.status(403).send("Permission Denied.");
            }

            req.uAuth = auth;
            companyManageRemove(req, res);
        } catch (err) {

            return res.status(403);
        }
    },

    modifyRole: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.status(403);
        }

        try {

            const auth = await verifyToken(token);
            const { userId } = req.body;

            if (Number(auth.com_auth) !== 0 || Number(userId) == Number(auth.uid)) {

                return res.status(403).send("Permission Denied.");
            }
            
            req.uAuth = auth;
            companyManageModifyRole(req, res);
        } catch (err) {

            return res.status(403);
        }
    },
}
