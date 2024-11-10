FROM node:22.9.0

WORKDIR /app

COPY . /app/

RUN ["npm", "install"]

EXPOSE 5173

RUN ["npm", "run", "build"]

CMD [ "npm", "start", "--", "--host", "0.0.0.0"]
