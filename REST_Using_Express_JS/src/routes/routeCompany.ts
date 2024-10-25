
import express from 'express';

import { companyAuthVerify } from '../middleware/authCompany';

const routerCompany = express.Router();

routerCompany.post('/create', companyAuthVerify.create);
routerCompany.get('/info', companyAuthVerify.getInfo);
routerCompany.put('/update', companyAuthVerify.update);
routerCompany.delete('/delete', companyAuthVerify.delete);

export default routerCompany;
