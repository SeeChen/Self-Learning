
import express from 'express';

import { userAuthVerify } from '../middleware/authUser';
import { userCreate, userUpdateById, userDeleteById, userLogin, userGetByKey } from '../controllers/controllerUser';

const routerUser = express.Router();

routerUser.post('/create', userCreate);
routerUser.post('/login', userLogin);
routerUser.get('/info', userAuthVerify, userGetByKey);
routerUser.put('/update', userAuthVerify, userUpdateById);
routerUser.delete('/delete', userAuthVerify, userDeleteById);

export default routerUser;
