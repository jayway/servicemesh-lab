FROM golang AS builder

ADD combinedgen.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app .

FROM alpine
COPY --from=builder /app /
EXPOSE 8080
CMD ["/app"]
