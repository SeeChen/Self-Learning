
import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export const daoUser = {

    create: async (data: any) => {

        return await prisma.user.create({

            data: data
        });
    },

    findByKey: async (id?: number, email?: string, params?: string[]) => {

        const allValue = ['id', 'name', 'email', 'address', 'password', 'userCompany']
        const selectObj: Record<string, boolean> = {};

        allValue.forEach(value => {

            selectObj[value] = params ? !params.includes(value) : true;
        });

        return await prisma.user.findUnique({

            where: {

                id: id ? Number(id) : undefined,
                email: email ? String(email) : undefined
            },

            select : selectObj
        });
    },

    update: async (id: number, data: any) => {

        return await prisma.user.update({

            where: { id },
            data
        });
    },

    delete: async (id: number) => {

        return await prisma.user.delete({

            where: { id: Number(id) }
        });
    },
}
