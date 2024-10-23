
import express from 'express';

import { authMiddleware } from '../middleware/authMiddleware';
import { companyCreate, companyGetById, companyGetByName, companyUpdate, companyDelete } from '../controllers/controllerCompany';

const routerCompany = express.Router();

routerCompany.post('/create', companyCreate);
routerCompany.get('/info/:id', companyGetById);
routerCompany.get('/info/:name', companyGetByName);
routerCompany.put('/update', companyUpdate);
routerCompany.delete('/delete', companyDelete);

export default routerCompany;
