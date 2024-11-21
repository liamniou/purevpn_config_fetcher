FROM alpine:3.19
RUN apk add --no-cache \
      chromium \
      chromium-swiftshader

ARG TARGETOS TARGETARCH

COPY --chmod="111" dist/purevpnwg-${TARGETOS}-${TARGETARCH} /bin/purevpnwg

ENV PUREVPN_USERNAME=""
ENV PUREVPN_PASSWORD=""

# Amsterdam
ENV PUREVPN_SERVER_COUNTRY="NL"
ENV PUREVPN_SERVER_CITY="2902"

ENTRYPOINT ["/bin/purevpnwg", "full"]
