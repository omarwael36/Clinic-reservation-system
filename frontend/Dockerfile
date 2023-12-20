FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/

ARG API_URL
RUN sed -i "s|'YOUR_API_URL_HERE'|'$API_URL'|g" src/environments/environment.ts

RUN npm run build --prod

FROM nginxinc/nginx-unprivileged:alpine3.18-perl

COPY --chown=nginx:nginx nginx.conf /etc/nginx/nginx.conf

COPY --chown=nginx:nginx --from=build /app/dist/frontend /var/www/html/
