FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN sed -i "s|environment.apiUrl = '';|environment.apiUrl = '${API_URL}';|g" src/environments/environment.ts && npm run build --prod

RUN npm run build --prod

FROM nginxinc/nginx-unprivileged:alpine3.18-perl
COPY --chown=nginx:nginx nginx.conf /etc/nginx/nginx.conf
COPY --chown=nginx:nginx --from=build /app/dist/frontend /var/www/html/

