FROM alpine:3.19
RUN apk add --no-cache \
      chromium \
      chromium-swiftshader

ARG TARGETOS TARGETARCH

COPY --chmod="111" ./purevpnwg-${TARGETOS}-${TARGETARCH} /bin/purevpnwg

ENV PUREVPN_USERNAME=""
ENV PUREVPN_PASSWORD=""
ENV PUREVPN_SERVER_COUNTRY="US"
# New York
ENV PUREVPN_SERVER_CITY="8778"
ENV PUREVPN_WIREGUARD_FILE="/out/wg0.conf"

ENTRYPOINT ["/bin/purevpnwg", "full"]