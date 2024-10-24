
import { Request, Response } from "express";
import { serviceCompany } from "../Services/serviceCompany";

export const companyCreate = async (req: Request, res: Response) => {

    const { name, address } = req.body;
    try {

        const company = await serviceCompany.create (
            name,
            address
        );
        res.json(company);
    } catch (err) {

        res.status(400).send(err.message);
    }
};

// export const company_Create = async (req: Request, res: Response) => {

//     const { name, address } = req.body;
//     const creatorId = req.cookies["uid"];

//     const userWithCompany = await prisma.userCompany.findUnique({ where: { userId: Number(creatorId) } });
//     if (userWithCompany) {

//         return res.status(409).send("Already have company.");
//     }

//     const companyExist = await prisma.company.findUnique({ where: { name } });
//     if (companyExist) {

//         return res.status(409).send("Company Exists.");
//     }

//     const newCompany = await prisma.company.create({

//         data: {

//             name,
//             address,
//             users: {
//                 create: {

//                     userId: Number(creatorId),
//                     role: 0
//                 }
//             }
//         },

//         include: {

//             users: true
//         }
//     });

//     await prisma.user.update({

//         where: { id: creatorId },
//         data: {

//             userCompany: {

//                 connect: { userId: Number(creatorId) }
//             }
//         }
//     });

//     res.json(newCompany);
// };

export const companyGetByKey = async (req: Request, res: Response) => {

    const { id, name } = req.query;
    try {

        const company = await serviceCompany.info(
            id ? Number(id) : undefined,
            name ? String(name) : undefined,
            []
        );

        res.json(company);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

// export const company_GetByKey = async (req: Request, res: Response) => {

//     const { id, name } = req.query;
//     const uid = req.cookies["uid"];

//     if (!id && !name) {

//         return res.status(400).send("Please provide either Company ID or Company Name!");
//     }

//     const company = await prisma.company.findUnique({

//         where: {

//             OR: [
//                 { id: id ? Number(id) : undefined },
//                 { name: name ? String(name) : undefined }
//             ]
//         },

//         select: {

//             id: true,
//             name: true,
//             address: true,
//             users: uid ? {
//                 where: { userId: Number(uid) }
//             } : false
//         }
//     });

//     if (!company) {

//         return res.status(404).send("Company Not Found!");
//     }

//     if (!uid || !company.users.length) {

//         return res.json({ 

//             id: company.id,
//             name: company.name,
//             address: company.address
//         });
//     }

//     res.json({

//         id: company.id,
//         name: company.name,
//         address: company.address,
//         users: company.users
//     });
// };

export const companyUpdate = async (req: Request, res: Response) => {

    const { id } = req.query;
    const { name, address } = req.body;

    try {
        
        const company = await serviceCompany.update(

            Number(id),
            name,
            address
        );

        res.json(company);
    } catch (err) {

        res.status(401).send(err.message);
    }
};

// export const company_Update = async (req: Request, res: Response) => {

//     const { id } = req.params;
//     const { name, address } = req.body;

//     const uid = req.cookies["uid"];

//     const userCompany = await prisma.userCompany.findUnique({

//         where: {

//             userId_companyId: {

//                 userId: Number(uid),
//                 companyId: Number(id)
//             }
//         }
//     });

//     if (!userCompany || userCompany.role !== 0) {

//         return res.status(403).send("Permission Denied.");
//     }

//     const companyOriginal = await prisma.company.findUnique({

//         where: { id: Number(id) }
//     });

//     const updatedCompany = await prisma.company.update({

//         where: { id: Number(id) },
//         data: {

//             name: name ? name : companyOriginal.name,
//             address: address ? address : companyOriginal.address
//         }
//     });

//     res.json(updatedCompany);
// };

export const companyDelete = async (req: Request, res: Response) => {

    const { id } = req.query;
    try {

        await serviceCompany.delete(Number(id));
        res.status(200).send("Delete Successful.");
    } catch (err) {

        res.status(401).send(err.message);
    }
};

// export const company_Delete = async (req: Request, res: Response) => {

//     const { id } = req.params;
//     const { adminPassword } = req.body;

//     const uid = req.cookies["uid"];

//     const user = await prisma.user.findUnique({ where: { id: Number(uid) } });
//     const company = await prisma.company.findUnique({ where: { id: Number(id) } });
//     const userCompany = await prisma.userCompany.findUnique({

//         where: {

//             userId_companyId: {

//                 userId: Number(uid),
//                 companyId: Number(id)
//             }
//         }
//     });

//     if (!user || !company) {

//         return res.status(404).send("User or Company not found!");
//     }

//     if (userCompany.role !== 0 || !userCompany || !(await verifyPassword(adminPassword, user.password))) {

//         return res.status(403).send("Permission Denied.");
//     }

//     await prisma.company.delete({ where: { id: Number(id) } });
//     res.sendStatus(204);
// };
