FROM node:20.10.0 as build

RUN mkdir -p /app
WORKDIR /app

COPY package*.json /app/
RUN npm install

COPY . /app/
RUN npm run build --prod

FROM nginx:alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/dist/frontend /usr/share/nginx/html

# Set proper permissions
RUN chmod -R 755 /usr/share/nginx/html

# Expose port
EXPOSE 80

# Copy startup script
COPY start.sh /start.sh
RUN chmod +x /start.sh

CMD ["/start.sh"]
