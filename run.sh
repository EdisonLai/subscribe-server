docker run -d --restart=always --network v2ray -p 443:443 -p 80:80 --name nginx -v /etc/nginx/conf.d/:/etc/nginx/conf.d/ -v /www:/www -v /ssl:/ssl nginx
docker run -d --restart=always --network v2ray -v /etc/v2ray:/etc/v2ray --name v2ray v2fly/v2fly-core
docker run -d --restart=always --network v2ray -p 8080:8080 -v ~/store:/root/store --name sub-server edisonlai/v-server:beta.v.2
docker run -d --restart=always --network v2ray -p 25500:25500 --name convertor tindy2013/subconverter:latest



docker run -d --restart=always -p 443:443 -p 80:80 --name nginx -v /etc/nginx/conf.d/:/etc/nginx/conf.d/ -v /www:/www -v /ssl:/ssl nginx
docker run -d --restart=always -p 8080:8080 -v ~/store:/root/store --name sub-server edisonlai/v-server:beta.v.2
docker run -d --restart=always -p 25500:25500 --name convertor tindy2013/subconverter:latest
