
import { PrismaClient } from "@prisma/client";
import { Request, Response } from "express";

const prisma = new PrismaClient();

export const companyCreate = async (req: Request, res: Response) => {

    const { name } = req.body;
    
    const companyExist = await prisma.company.findUnique({ where: { name } });
    if (companyExist) {
        
        return res.status(409).json({ message: "Company exists." });
    }

    const company = await prisma.company.create({ data: { name } });

    res.json(company);
};

export const companyGetById = async (req: Request, res: Response) => {

    const { id } = req.params;
    const company = await prisma.company.findUnique({ where: { id: Number(id) } });
  
    if (!company) {

      return res.status(404).json({ message: "Company not found." });
    }
  
    res.json(company);
};

export const companyGetByName = async (req: Request, res: Response) => {

    const { name } = req.params;
    const company = await prisma.company.findUnique({ where: { name } });
  
    if (!company) {
        
        return res.status(404).json({ message: "Company not found." });
    }
  
    res.json(company);
};

export const companyUpdate = async (req: Request, res: Response) => {
    
    const { id, name } = req.body;

    const company = await prisma.company.update({
      
        where: { id: Number(id) },
        data: { name }
    });
    
    res.json(company);
};

export const companyDelete = async (req: Request, res: Response) => {
    
    const { id } = req.body;
    
    await prisma.company.delete({ where: { id: Number(id) } });
    
    res.sendStatus(204);
};