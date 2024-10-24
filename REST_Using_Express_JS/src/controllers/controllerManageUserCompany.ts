
import { PrismaClient } from "@prisma/client";
import { Request, Response } from "express";

const prisma = new PrismaClient();

export const companyManageAssignUser = async (req: Request, res: Response) => {

    const uid = req.cookies["uid"];
    const { companyId } = req.params;

    const { userId, role } = req.body;

    const userWithCompany = await prisma.userCompany.findUnique({ where: { userId: Number(userId) } });
    if (userWithCompany) {

        return res.status(409).send("The user is already assigned to a company.");
    }

    const adminCompany = await prisma.userCompany.findUnique({

        where: {

            userId_companyId: {

                userId: Number(uid),
                companyId: Number(companyId)
            }
        }
    });


    if (!adminCompany || adminCompany.role !== 0) {

        return res.status(403).send("Permission Denied.");
    }

    if (role !== 0 && role !== 1) {

        return res.status(400).send("Invalid role.");
    }

    const userCompany = await prisma.userCompany.create({

        data: {

            user: { connect: { id: userId } },
            company: { connect: { id: companyId } },
            role
        }
    });

    return res.status(201).json(userCompany);
};

export const companyManageRemoveUser = async (req: Request, res: Response) => {

    const uid = req.cookies["uid"];
    const { companyId } = req.params;
    const { userId } = req.body;

    const adminCompany = await prisma.userCompany.findUnique({

        where: {

            userId_companyId: {

                userId: Number(uid),
                companyId: Number(companyId)
            }
        }
    });

    if (!adminCompany || adminCompany.role !== 0) {

        return res.status(403).send("Permission Denied.");
    }

    const userCompany = await prisma.userCompany.findUnique({

        where: {

            userId_companyId: {

                userId: Number(userId),
                companyId: Number(companyId)
            }
        }
    });

    if (!userCompany) {

        return res.status(404).send("User not found in the specified company.");
    }

    await prisma.userCompany.delete({

        where: {

            userId_companyId: {

                userId: Number(userId),
                companyId: Number(companyId)
            }
        }
    });

    return res.sendStatus(204);
};

export const companyManageModifyRole = async (req: Request, res: Response) => {

    const uid = req.cookies["uid"];
    const { companyId } =  req.params;

    const { userId, newRole } = req.body;

    const adminCompany = await prisma.userCompany.findUnique({

        where: {

            userId_companyId: {

                userId: Number(uid),
                companyId: Number(companyId)
            }
        }
    });

    if (!adminCompany || adminCompany.role !== 0) {

        return res.status(403).send("Permission Denied.");
    }

    const userCompany = await prisma.userCompany.findUnique({

        where: {

            userId_companyId: {

                userId: Number(userId),
                companyId: Number(companyId)
            }
        }
    });

    if (!userCompany) {

        return res.status(404).send("User not found in the specified company.");
    }

    if (![0, 1].includes(Number(newRole))) {
        
        return res.status(400).send("Invalid role.");
    }

    const updatedUserCompany = await prisma.userCompany.update({

        where: {

            userId_companyId: {

                userId: Number(userId),
                companyId: Number(companyId)
            }
        },

        data: {

            role: Number(newRole)
        }
    });

    return res.status(200).json(updatedUserCompany);
};
