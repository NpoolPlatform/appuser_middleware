FROM jrei/systemd-centos:7

COPY AppUserMiddleware.viper.yaml /etc/AppUserMiddleware/AppUserMiddleware.viper.yaml
COPY appuser-middleware /usr/local/bin/appuser-middleware

CMD ["appuser-middleware", "run"]