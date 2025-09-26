FROM nginx
RUN apt-get update && apt-get install -y vim silversearcher-ag
COPY static-site /usr/share/nginx/html
