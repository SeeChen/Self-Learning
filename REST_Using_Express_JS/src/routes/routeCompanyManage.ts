
import express from 'express';

import { companyManageAuthVerify } from '../middleware/authManage';

const routeCompanyManage = express.Router();

routeCompanyManage.post("/assign", companyManageAuthVerify.assign);
routeCompanyManage.delete("/remove", companyManageAuthVerify.remove);
routeCompanyManage.put("/modify-role", companyManageAuthVerify.modifyRole);

export default routeCompanyManage;
