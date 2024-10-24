
import express from 'express';

import { companyCreate, companyGetByKey, companyUpdate, companyDelete } from '../controllers/controllerCompany';

const routerCompany = express.Router();

routerCompany.post('/create', companyCreate);
routerCompany.get('/info', companyGetByKey);
routerCompany.put('/update', companyUpdate);
routerCompany.delete('/delete', companyDelete);

export default routerCompany;
