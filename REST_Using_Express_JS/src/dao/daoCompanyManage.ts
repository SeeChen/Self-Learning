
import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient();

export const daoCompanyManage = {

    create: async (companyId: number, userId: number, role: number) => {

        return await prisma.userCompany.create({

            data: {

                user: { connect: { id: userId } },
                company: { connect: { id: companyId } },
                role
            }
        });
    },

    findByUserId: async (userId: number) => {

        return await prisma.userCompany.findUnique({

            where: {

                userId: Number(userId)
            }
        });
    },

    findByDoubleKey: async (companyId: number, userId: number) => {

        return await prisma.userCompany.findUnique({

            where: {

                userId_companyId: {

                    userId: Number(userId),
                    companyId: Number(companyId)
                }
            }
        });
    },

    modifyRole: async (companyId: number, userId: number, roleNew: number) => {

        return await prisma.userCompany.update({

            where: {

                userId_companyId: {

                    userId: Number(userId),
                    companyId: Number(companyId)
                }
            },

            data: {

                role: Number(roleNew)
            }
        });
    },

    delete: async (companyId: number, userId: number) => {

        return await prisma.userCompany.delete({

            where: {

                userId_companyId: {

                    userId: Number(userId),
                    companyId: Number(companyId)
                }
            }
        });
    },
}