FROM alpine:3.7

ADD run.sh /tmp/run.sh
RUN chmod +x /tmp/run.sh

# Add bad alpine packages
RUN apk add samba=4.7.6-r3

CMD ["/tmp/run.sh"]
