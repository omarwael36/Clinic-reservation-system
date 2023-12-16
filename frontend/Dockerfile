FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

FROM nginx:alpine
RUN mkdir -p /var/cache/nginx/client_temp
RUN chown -R nginx:nginx /var/cache/nginx
RUN chmod 777 /usr/share/nginx/html

# Copy Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built frontend files
COPY --from=build /app/dist/frontend /usr/share/nginx/html

EXPOSE 80

# Set the environment variable for API URL
CMD ["sh", "-c", "sed -i 's|API_URL_PLACEHOLDER|$API_URL|' /usr/share/nginx/html/main*.js && nginx -g 'daemon off;'"]