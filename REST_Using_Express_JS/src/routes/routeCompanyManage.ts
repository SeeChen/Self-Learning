
import express from 'express';


import { companyManageAssign, companyManageRemove, companyManageModifyRole } from '../controllers/controllerCompanyManage';

const routeCompanyManage = express.Router();

routeCompanyManage.post("/assign", companyManageAssign);
routeCompanyManage.delete("/remove", companyManageRemove);
routeCompanyManage.put("/modify-role", companyManageModifyRole);

export default routeCompanyManage;
