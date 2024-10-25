
import { daoCompany } from "../dao/daoCompany"
import { daoCompanyManage } from "../dao/daoCompanyManage";
import { daoUser } from "../dao/daoUser";
import { verifyPassword } from "../auth/authService";
 
 export const serviceCompany = {

    async create (

        creatorId: number,
        name: string,
        address: string
    ) {

        const userWithCompany = await daoCompanyManage.findByUserId(creatorId);
        if (userWithCompany) {

            throw new Error("Already have company.");
        }

        const companyExists = await daoCompany.findByKey(undefined, name);
        if (companyExists) {

            throw new Error("Company Exists.");
        }

        return await daoCompany.create({

            name,
            address,
            users: {

                create: {

                    userId: Number(creatorId),
                    role: 0
                }
            }
        });
    },

    async info (

        UID?: number,
        id?: number,
        name?: string
    ) {

        if (!id && !name) {

            throw new Error("Please provide either Company Id or Company Name.");
        }

        const company = await daoCompany.findByKey(
            id ? Number(id) : undefined,
            name ? String(name) : undefined,
            [UID && id ? (await daoCompanyManage.findByDoubleKey(Number(id), Number(UID)) ? "" : "users") : "users"]
        );

        if (!company) {

            throw new Error("Company Does not Exists");
        }

        return company;
    },

    async update (

        adminId: number,
        id: number,
        name?: string,
        address?: string
    ) {

        const company = await daoCompany.findByKey(id);
        if (!company) {

            throw new Error("Company Does not Exists.");
        }

        const userWithCompany = await daoCompanyManage.findByDoubleKey(Number(id), Number(adminId));
        if (!userWithCompany || userWithCompany.role !== 0) {

            throw new Error("Permission Denied.");
        }

        name = name || company.name;
        address = address || company.address;

        return await daoCompany.update(
            id, {

                name,
                address
            }
        );
    },

    async delete (

        adminId: number,
        adminEmail: string,
        adminPassword: string,
        companyId: number,
        companyName: string
    ) {

        const user = await daoUser.findByKey(Number(adminId));
        const company = await daoCompany.findByKey(companyId);
        if (!user || !company) {

            throw new Error("Target not found.");
        }

        const userWithCompany = await daoCompanyManage.findByDoubleKey(companyId, adminId);
        if (!userWithCompany || userWithCompany.role !== 0 || (companyName !== company.name) || (adminEmail !== user.email) || !await verifyPassword(adminPassword, user.password)) {

            throw new Error("Permission Denied.");
        }

        await daoCompany.delete(companyId);
    },
 }
