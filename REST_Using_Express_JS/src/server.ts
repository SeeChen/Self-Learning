
import express from 'express'
import routerUser from './routes/routeUser';

const app = express();
const PORT = process.env.PORT || 1234;

app.use(express.json());
app.use('/api/users', routerUser)

app.get('/', async (req, res) => {

  res.send("<html><head><title>TITLE</title></head><body><h1>SEECHEN</h1></body></html>")
})

app.listen(PORT, () => {

  console.log(`Server is running on port: ${PORT}`);
})
