FROM scratch
COPY task /
ENTRYPOINT ["/task"]