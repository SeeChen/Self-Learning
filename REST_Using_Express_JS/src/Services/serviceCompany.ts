
import { daoCompany } from "../dao/daoCompany"
 
 export const serviceCompany = {

    async create (

        name: string,
        address: string
    ) {

        const companyExists = await daoCompany.findByKey(undefined, name);
        if (companyExists) {

            throw new Error("Company Exists.");
        }

        return await daoCompany.create({

            name,
            address
        });
    },

    async info (

        id?: number,
        name?: string,
        params?: string[]
    ) {

        if (!id && !name) {

            throw new Error("Please provide either Company Id or Company Name.");
        }

        const company = await daoCompany.findByKey(
            id ? Number(id) : undefined,
            name ? String(name) : undefined,
            params
        );

        if (!company) {

            throw new Error("Company Does not Exists");
        }

        return company;
    },

    async update (

        id: number,
        name?: string,
        address?: string
    ) {

        const company = await daoCompany.findByKey(id);
        if (!company) {

            throw new Error("Company Does not Exists.");
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

        id: number
    ) {

        const company = await daoCompany.findByKey(id);
        if (!company) {

            throw new Error("Company Does Not Exists.");
        }

        await daoCompany.delete(id);
    },
 }
