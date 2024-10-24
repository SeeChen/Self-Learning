
import express from 'express'
import routerUser from './routes/routeUser';
import routerCompany from './routes/routeCompany';
import routeCompanyManage from './routes/routeCompanyManage';

const app = express();
const cookieParser = require('cookie-parser')
const PORT = process.env.PORT || 1234;

app.use(express.json());
app.use(cookieParser());
app.use('/api/users', routerUser)
app.use('/api/company', routerCompany);
app.use('/api/company/manage', routeCompanyManage);

app.get('/', async (req, res) => {

  res.send("<html><head><title>TITLE</title></head><body><h1>SEECHEN</h1></body></html>")
})

app.listen(PORT, () => {

  console.log(`Server is running on port: ${PORT}`);
})
