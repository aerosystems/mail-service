FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs

COPY ./mail-service/mail-service.bin /app
COPY ./mail-service/templates /templates

# Run the server executable
CMD [ "/app/mail-service.bin" ]