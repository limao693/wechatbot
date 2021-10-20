FROM alpine:3.13.6
COPY wechatbot /src/
COPY entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 6001
WORKDIR /src/