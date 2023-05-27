FROM node:latest as build

WORKDIR /app
COPY package*.json /app/
RUN yarn install
COPY . /app
RUN yarn build


FROM nginx
COPY --from=build /app/build /usr/share/nginx/html
COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf