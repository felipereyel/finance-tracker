FROM node:18-alpine as vueapp
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY ./public ./public
COPY index.html tsconfig.json tsconfig.node.json vite.config.ts ./

# src is mounted as volume

CMD ["npm", "run", "dev"]