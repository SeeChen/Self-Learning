
import { daoCompany } from "../dao/daoCompany";
import { daoCompanyManage } from "../dao/daoCompanyManage"
import { daoUser } from "../dao/daoUser";

export const serviceCompanyManage = {

    async assign (

        adminId: number,
        companyId: number,
        userId: number,
        role: number
    ) {

        const [company, user] = await Promise.all([

            daoCompany.findByKey(companyId),
            daoUser.findByKey(userId)
        ]);
        if (!company || !user) {

            throw new Error(`Target company (ID: ${companyId}) or target user (ID: ${userId}) not found.`);
        }

        const userWithCompany = await daoCompanyManage.findByUserId(userId);
        if (userWithCompany) {

            throw new Error("The user is already assigned to a company.");
        }

        const adminWithCompany = await daoCompanyManage.findByDoubleKey(companyId, adminId);
        if (!adminWithCompany || adminWithCompany.role !== 0) {

            throw new Error("Permission Denied.");
        }

        const isValidRole = (role: number) => role === 0 || role === 1;
        if (!isValidRole(role)) {

            throw new Error(`Invalid role: ${role}.`);
        }

        return await daoCompanyManage.create(companyId, userId, role);
    },

    async remove (

        adminId: number,
        companyId: number,
        userId: number
    ) {

        const userWithCompany = await daoCompanyManage.findByDoubleKey(companyId, userId);
        if (!userWithCompany) {

            throw new Error(`User (ID: ${userId}) is not associated with Company (ID: ${companyId}).`);
        }

        const adminWithCompany = await daoCompanyManage.findByDoubleKey(companyId, adminId);
        if (!adminWithCompany || adminWithCompany.role !== 0) {

            throw new Error(`Permission Denied for user (ID: ${adminId}).`);
        }

        await daoCompanyManage.delete(companyId, userId);
    },

    async modifyRole (

        adminId: number,
        companyId: number,
        userId: number,
        roleNew?: number
    ) {

        const userWithCompany = await daoCompanyManage.findByDoubleKey(companyId, userId);
        if (!userWithCompany) {

            throw new Error("Incorrect Company Id of User Id.");
        }

        const adminWithCompany = await daoCompanyManage.findByDoubleKey(companyId, adminId);
        if (!adminWithCompany || adminWithCompany.role !== 0) {

            throw new Error(`Permission Denied for user (ID: ${adminId}).`);
        }

        const isValidRole = (role: number) => role === 0 || role === 1;
        roleNew = roleNew !== undefined ? roleNew : userWithCompany.role;
        if (!isValidRole(roleNew)) {

            throw new Error(`Invalid role: ${roleNew}.`);
        }

        return await daoCompanyManage.modifyRole(

            companyId,
            userId,
            roleNew
        );
    }
}
