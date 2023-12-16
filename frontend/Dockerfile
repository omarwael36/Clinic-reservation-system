FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

FROM nginxinc/nginx-unprivileged:1.21.4-alpine

# Set permissions for Nginx user
USER root
RUN chown -R nginx:nginx /usr/share/nginx/html
RUN chmod -R 777 /usr/share/nginx/html

# Add 'cap_net_bind_service' capability to the Nginx binary
RUN apk add --no-cache libcap
RUN setcap 'cap_net_bind_service=+ep' /usr/sbin/nginx

# Remove unnecessary packages
RUN apk del libcap

# Copy Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built frontend files
COPY --from=build /app/dist/frontend /usr/share/nginx/html

EXPOSE 8080

# Set the environment variable for API URL
CMD ["sh", "-c", "sed -i 's|API_URL_PLACEHOLDER|$API_URL|' /usr/share/nginx/html/main*.js && nginx -g 'daemon off;'"]