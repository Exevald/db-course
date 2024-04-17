FROM node:18-alpine
WORKDIR /app/

COPY public/ /app/public
COPY src/ /app/src
COPY package.json /app/

COPY . .
#RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]