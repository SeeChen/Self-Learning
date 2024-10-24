
import express from 'express';

import { authMiddleware } from '../middleware/authMiddleware';
import { companyManageAssignUser, companyManageRemoveUser, companyManageModifyRole } from '../controllers/controllerManageUserCompany';

const routeCompanyManage = express.Router();

routeCompanyManage.post("/assign", companyManageAssignUser);
routeCompanyManage.delete("/remove", companyManageRemoveUser);
routeCompanyManage.put("/modify-role", companyManageModifyRole);

export default routeCompanyManage;
