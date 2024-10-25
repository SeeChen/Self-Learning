
import express from 'express';

import { userAuthVerify } from '../middleware/authUser';
import { userCreate, userLogin} from '../controllers/controllerUser';

const routerUser = express.Router();

routerUser.post('/create', userCreate);
routerUser.post('/login', userLogin);
routerUser.get('/info', userAuthVerify.info);
routerUser.put('/update', userAuthVerify.update);
routerUser.delete('/delete', userAuthVerify.delete);

export default routerUser;
