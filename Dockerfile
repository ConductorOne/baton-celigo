FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-celigo"]
COPY baton-celigo /