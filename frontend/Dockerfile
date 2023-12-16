FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

FROM nginxinc/nginx-unprivileged:1.21.4-alpine

# Update Nginx configuration to listen on port 8080 instead of port 80
RUN sed -i 's/80;/8080;/g' /etc/nginx/conf.d/default.conf

# Set permissions for Nginx user
RUN chown -R nginx:nginx /var/cache/nginx /var/run/nginx

# Copy Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built frontend files
COPY --from=build /app/dist/frontend /usr/share/nginx/html

EXPOSE 8080

# Set the environment variable for API URL
CMD ["sh", "-c", "sed -i 's|API_URL_PLACEHOLDER|$API_URL|' /usr/share/nginx/html/main*.js && exec nginx -g 'daemon off;'"]