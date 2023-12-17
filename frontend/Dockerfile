# Stage 1: Build stage
FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

# Stage 2: Production stage
FROM nginxinc/nginx-unprivileged:1.21.4-alpine

# Create a non-root user
RUN adduser -D -u 1001 nginxuser
USER nginxuser

# Set permissions for Nginx user
RUN chown -R nginxuser:nginxuser /usr/share/nginx/html
RUN chown -R nginxuser:nginxuser /var/run/nginx.pid

# Copy Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy built frontend files
COPY --from=build /app/dist/frontend /usr/share/nginx/html

EXPOSE 8081

# Switch to the non-root user

# Set the environment variable for API URL
CMD ["sh", "-c", "sed -i 's|API_URL_PLACEHOLDER|$API_URL|' /usr/share/nginx/html/main*.js && nginx -g 'daemon off;'"]
