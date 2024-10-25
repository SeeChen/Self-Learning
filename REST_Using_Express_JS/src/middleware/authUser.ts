
import { verifyPassword, verifyToken } from "../auth/authService";
import { Request, Response} from "express";
import { userDeleteById, userGetByKey, userUpdateById } from "../controllers/controllerUser";

export const userAuthVerify = {

    info: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (token) {

            try {

                const auth = await verifyToken(token);
                req.uAuth = auth;
            } catch (err) {

                res.status(403).send(err.message);
            }
        }

        userGetByKey(req, res);
    },

    update: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.sendStatus(403);
        }

        try {

            const auth = await verifyToken(token);
            req.uAuth = auth;

            if (req.body.passwordOld && !(await verifyPassword(req.body.passwordOld, String(auth.pwd)))) {

                return res.status(403).send("Incorrect Password.")
            }

            userUpdateById(req, res);
        } catch (err) {

            return res.status(403).send(err.message);
        }
    },

    delete: async (req: Request, res: Response): Promise<any> => {

        const token = req.cookies["auth"];
        if (!token) {

            res.sendStatus(403);
        }

        try {

            const auth = await verifyToken(token);

            const { name, email, password } = req.body;
            if (name !== auth.name || email !== auth.email || !(await verifyPassword(password, auth.pwd))) {

                return res.status(403).send("Delete Not Complete.");
            }

            req.uAuth = auth;

            userDeleteById(req, res);
        } catch (err) {

            return res.status(403).send(err.message);
        }
    },
}
