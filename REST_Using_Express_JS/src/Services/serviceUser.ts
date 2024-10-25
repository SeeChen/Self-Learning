
import { daoUser } from "../dao/daoUser";
import { hashPassword, verifyPassword } from "../auth/authService";

export const serviceUser = {

    async create (
        
        name: string,
        email: string,
        password: string,
        address: string
    ) {

        const userExists = await daoUser.findByKey(undefined, email);
        if (userExists) {

            throw new Error("User Exists.");
        }

        const hashedPassword = await hashPassword(password);
        return await daoUser.create({

            name,
            email,
            password: hashedPassword,
            address
        });
    },

    async login (

        email: string,
        password: string
    ) {

        const user = await daoUser.findByKey(undefined, email);
        if (!user) {

            throw new Error("User Does not Exists.");
        }

        if (!(await verifyPassword(password, user.password))) {

            throw new Error("Incorrect Password!");
        }

        return user;
    },

    async verify (

        name: string,
        email: string,
        password: string
    ) {

        try {

            const user = await serviceUser.login(email, password);
            if (name !== user.name) {

                throw new Error("Incorrect Name");
            }

            return user;

        } catch (err: any) {

            throw new Error(err.message)
        }
    },

    async info (

        id?: number,
        email?: string,
        params?: string[]
    ) {

        if (!id && !email) {

            throw new Error("Please provide either User Id or User Email.");
        }

        const user =  await daoUser.findByKey(
            id ? Number(id) : undefined,
            email ? String(email) : undefined,
            params
        );

        if (!user) {

            throw new Error("User Not Found.");
        }

        return user;
    },

    async update (

        id: number,
        name?: string,
        email?: string,
        pwdOld?: string,
        pwdNew?: string,
        address?: string
    ) {

        const user = await daoUser.findByKey(id);
        if (!user) {

            throw new Error("User Not Fouund.");
        }

        name = name || user.name;
        email = email || user.email;
        pwdNew = pwdNew ? await hashPassword(pwdNew) : user.password;
        address = address || user.address;

        if (pwdOld && !(await verifyPassword(pwdOld, user.password))) {

            throw new Error("Incorrect Password.");
        }

        return await daoUser.update(

            id, {

                name,
                email,
                password: pwdNew,
                address
            }
        );
    },

    async delete (

        id: number,
        name: string,
        email: string,
        password: string
    ) {

        try {

            const user = await serviceUser.verify(name, email, password);
            await daoUser.delete(id);
        } catch (err) {

            throw new Error(err.message);
        }
    },
}
