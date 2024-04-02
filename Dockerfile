FROM alpine
RUN apk add --no-cache \
      chromium \
      chromium-swiftshader

COPY dist/purevpnwg-${BUILDOS}-${BUILDARCH} /bin/purevpnwg

ENV PUREVPN_USERNAME=""
ENV PUREVPN_PASSWORD=""
ENV PUREVPN_SERVER_COUNTRY="DE"
ENV PUREVPN_SERVER_CITY="2762"
ENV PUREVPN_WIREGUARD_FILE="/out/wg0.conf"

ENTRYPOINT ["/bin/purevpnwg", "full"]