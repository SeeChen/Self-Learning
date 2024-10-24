
import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export const daoCompany = {

    create: async (data: any) => {

        return await prisma.company.create({

            data: data
        });
    },

    findByKey: async (id?: number, name?: string, params?: string[]) => {

        const allValue = ['id', 'name', 'address', 'users'];
        const selectObj: Record<string, boolean> = {};

        allValue.forEach(value => {

            selectObj[value] = params ? !params.includes(value) : true;
        });

        return await prisma.company.findUnique({

            where: {

                id: id ? Number(id) : undefined,
                name: name ? String(name) : undefined
            },

            select : selectObj
        });
    },

    update: async (id: number, data: any) => {

        return await prisma.company.update({

            where: {id: Number(id)},
            data
        });
    },

    delete: async (id: number) => {

        return await prisma.company.delete({

            where: { id: Number(id) }
        });
    },
}
