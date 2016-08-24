FROM scratch

COPY app /

ENTRYPOINT ["/app"]

EXPOSE 8080
