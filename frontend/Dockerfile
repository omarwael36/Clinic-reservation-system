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

# Copy Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built frontend files
COPY --from=build /app/dist/frontend /usr/share/nginx/html

EXPOSE 80

# Set the environment variable for API URL
CMD ["nginx", "-g", "daemon off;"]
