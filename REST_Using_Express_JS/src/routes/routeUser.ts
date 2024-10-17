
import express from 'express';

import { authMiddleware } from '../middleware/authMiddleware';
import { userCreate, userGetById, userUpdateById, userDeleteById, userLogin } from '../controllers/controllerUser';

const routerUser = express.Router();

routerUser.post('/create', userCreate);
routerUser.get('/:id', authMiddleware, userGetById);
routerUser.put('/:id', authMiddleware, userUpdateById);
routerUser.delete('/:id', authMiddleware, userDeleteById);
routerUser.post('/login', userLogin);

export default routerUser;
