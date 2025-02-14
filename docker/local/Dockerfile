# ARG service
FROM alpine 
RUN mkdir /ftgogo
# WORKDIR /app
COPY ./monolith /ftgogo/
# RUN ls -l /app/app/meta/build
ENTRYPOINT ["/ftgogo/monolith"]

