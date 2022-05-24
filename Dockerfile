FROM golang:1.18

# Set environment variable
ENV APP_NAME where_to_stream
ENV CMD_PATH cmd/main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run Stage
FROM golang:1.18
 
# Set environment variable
ENV APP_NAME where_to_stream
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .
 
# Expose application port
EXPOSE 8080
 
# Start app
CMD ./$APP_NAME