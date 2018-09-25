# BenchMarkExporter
A Prometheus exporter example and Generate rand data

### Use Docker

	docker run -e n=24 -e m=100 -d -p 9120:9120 daocloud.io/daocloud/benchmark-exporter:v1

	m为Metrics的总数
	n为每个Metrics的name长度 
	
