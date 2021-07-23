# elasticsearch 安装
## 关闭防火墙
```
systemctl stop firewalld.service
systemctl disable firewalld.service
systemctl status firewalld.service
```

## 通过docker安装elasticsearch
```
#新建es的config配置文件夹
mkdir -p /data/elasticsearch/config
#新建es的data目录
mkdir -p /data/elasticsearch/data
#新建es的plugins目录
mkdir -p /data/elasticsearch/plugins
#给目录设置权限
chmod 777 -R /data/elasticsearch

echo "http.host: 0.0.0.0" >> /data/elasticsearch/config/elasticsearch.yml
#安装es
docker run --name elasticsearch -p 9200:9200 -p 9300:9300 \
	-e "discovery.type=single-node" \
  -e ES_JAVA_OPTS="-Xms128m -Xmx256m" \
  -v /data/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
  -v /data/elasticsearch/data:/usr/share/elasticsearch/data \
  -v /data/elasticsearch/plugins:/usr/share/elasticsearch/plugins \
  -d elasticsearch:7.10.1

```

## 通过docker安装 kibana
```
docker run -d --name kibana -e ELASTICSEARCH_HOSTS="http://192.168.0.104:9200" -p 5601:5601 kibana:7.10.1

```