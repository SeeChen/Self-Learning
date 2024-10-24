
import { daoCompany } from "../dao/daoCompany";
import { daoCompanyManage } from "../dao/daoCompanyManage"
import { daoUser } from "../dao/daoUser";

export const serviceCompanyManage = {

    async assign (

        companyId: number,
        userId: number,
        role: number
    ) {

        const company = await daoCompany.findByKey(companyId);
        const user = await daoUser.findByKey(userId);

        if (!company || !user) {

            throw new Error("Target company or target user not found.");
        }

        const userWithCompany = await daoCompanyManage.findByUserId(userId);
        if (userWithCompany) {

            throw new Error("The user is already assigned to a company.");
        }

        if (role !== 0 && role !== 1) {

            throw new Error("Invalid role.")
        }

        return await daoCompanyManage.create(companyId, userId, role);
    },

    async remove (

        companyId: number,
        userId: number
    ) {

        const userWithCompany = await daoCompanyManage.findByDoubleKey(companyId, userId);
        if (!userWithCompany) {

            throw new Error("Incorrect Company Id or User Id.");
        }

        await daoCompanyManage.delete(companyId, userId);
    },

    async modifyRole (

        companyId: number,
        userId: number,
        roleNew?: number
    ) {

        const userWithCompany = await daoCompanyManage.findByDoubleKey(companyId, userId);
        if (!userWithCompany) {

            throw new Error("Incorrect Company Id of User Id.");
        }

        roleNew = roleNew || userWithCompany.role;

        return await daoCompanyManage.modifyRole(

            companyId,
            userId,
            roleNew
        );
    }
}
