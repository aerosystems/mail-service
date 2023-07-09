FROM alpine:latest
RUN mkdir /app

COPY ./mail-service/mail-service.bin /app

# Run the server executable
CMD [ "/app/mail-service.bin" ]