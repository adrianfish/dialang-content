FROM nginx:stable-alpine3.21
COPY static-site /usr/share/nginx/html
