FROM node:14

WORKDIR /usr/src/app

COPY vuln_app/package.json ./
COPY vuln_app/server.js ./

RUN npm install

EXPOSE 8080

CMD ["node", "server.js"]
