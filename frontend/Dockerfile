FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

FROM nginxinc/nginx-unprivileged:alpine3.18-perl

# Create the target directory with the correct ownership
USER root
RUN mkdir -p /var/www/html && chown -R nginx:nginx /var/www

# Switch back to non-root user
USER nginx

# Copy Nginx configuration
COPY --chown=nginx:nginx nginx.conf /etc/nginx/nginx.conf

# Copy built frontend files
COPY --chown=nginx:nginx --from=build /app/dist/frontend /var/www/html/

EXPOSE 8080

# Replace placeholders without creating temporary files
CMD ["sh", "-c", "find /var/www/html -type f -name 'main*.js' -exec sed -i 's|API_URL_PLACEHOLDER|$API_URL|' {} + && nginx -g 'daemon off;'"]
