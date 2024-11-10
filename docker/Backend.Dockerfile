FROM node:22.9.0

WORKDIR /app

COPY . /app/

RUN ["npm", "install"]

EXPOSE 5001

CMD [ "node", "src/api/app.mjs", "--", "--host", "0.0.0.0"]
