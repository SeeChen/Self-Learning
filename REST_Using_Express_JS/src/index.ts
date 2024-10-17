import express from 'express';
import { PrismaClient } from '@prisma/client';
import jwt from 'jsonwebtoken';
import bcrypt from 'bcryptjs';

const app = express();
const prisma = new PrismaClient();
const PORT = process.env.PORT || 3000;

app.use(express.json());

// 中间件用于保护路由
const authenticateJWT = (req: any, res: any, next: any) => {
  const token = req.header('Authorization')?.split(' ')[1];

  if (token) {
    jwt.verify(token, 'your_jwt_secret', (err: any, user: any) => {
      if (err) {
        return res.sendStatus(403);
      }
      req.user = user;
      next();
    });
  } else {
    res.sendStatus(401);
  }
};

// 创建用户
app.post('/users', async (req, res) => {
  const { name, email, password, address } = req.body;
  const hashedPassword = await bcrypt.hash(password, 10);
  const user = await prisma.user.create({
    data: { name, email, password: hashedPassword, address },
  });
  res.json(user);
});

// 根据 ID 获取用户
app.get('/users/:id', authenticateJWT, async (req, res) => {
  const { id } = req.params;
  const user = await prisma.user.findUnique({ where: { id: Number(id) } });
  res.json(user);
});

// 根据 ID 更新用户
app.put('/users/:id', authenticateJWT, async (req, res) => {
  const { id } = req.params;
  const { name, email, password, address } = req.body;
  const data: any = { name, email, address };

  if (password) {
    data.password = await bcrypt.hash(password, 10);
  }

  const user = await prisma.user.update({
    where: { id: Number(id) },
    data,
  });
  res.json(user);
});

// 根据 ID 删除用户
app.delete('/users/:id', authenticateJWT, async (req, res) => {
  const { id } = req.params;
  await prisma.user.delete({ where: { id: Number(id) } });
  res.sendStatus(204);
});

// 使用密码登录
app.post('/login', async (req, res) => {
  const { email, password } = req.body;
  const user = await prisma.user.findUnique({ where: { email } });

  if (user && (await bcrypt.compare(password, user.password))) {
    const token = jwt.sign({ userId: user.id }, 'your_jwt_secret', { expiresIn: '1h' });
    res.json({ token });
  } else {
    res.sendStatus(403);
  }
});

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
