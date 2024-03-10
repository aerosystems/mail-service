FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs

COPY ./mail-service.bin /app
COPY ./templates /templates

# Run the server executable
CMD [ "/app/mail-service.bin" ]