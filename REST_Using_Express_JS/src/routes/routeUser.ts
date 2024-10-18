
import express from 'express';

import { authMiddleware } from '../middleware/authMiddleware';
import { userCreate, userGetById, userUpdateById, userDeleteById, userLogin } from '../controllers/controllerUser';

const routerUser = express.Router();

routerUser.post('/create', userCreate);
routerUser.post('/login', userLogin);
routerUser.get('/info/:id', authMiddleware, userGetById);
routerUser.put('/update', authMiddleware, userUpdateById);
routerUser.delete('/delete', authMiddleware, userDeleteById);

export default routerUser;
