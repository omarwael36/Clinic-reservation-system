# Stage 1: Build
FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

# Stage 2: Run with NGINX
FROM nginx:alpine

# Copy NGINX configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the built Angular app to NGINX HTML directory
COPY --from=build /app/dist/frontend /usr/share/nginx/html

# Ensure proper permissions for NGINX directories (optional)
RUN chown -R nginx:nginx /usr/share/nginx/html

# Expose port
EXPOSE 80

# Start NGINX using an unprivileged user (nginxinc/nginx-unprivileged style)
USER nginx

# Set the environment variable for API URL and start NGINX
CMD ["sh", "-c", "sed -i 's|API_URL_PLACEHOLDER|$API_URL|' /usr/share/nginx/html/main*.js && exec nginx -g 'daemon off;'"]
