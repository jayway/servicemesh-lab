FROM golang AS builder

ADD namegen.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app .

FROM scratch

COPY --from=builder /app /
EXPOSE 8080
CMD ["/app"]
