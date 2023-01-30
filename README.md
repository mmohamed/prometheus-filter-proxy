# Prometheus filter & authentication proxy

This reverse-proxy was developed to resolve a multi-tenancy case using one mutual Grafana and Prometheus.
We need to secure the prometheus endpoint behind a username/password and we need to add a global custom filtering series for each tenant. 

## Set up

1- Configure and deploy the proxy in front of your Prometheus server instances, to add the proxy as Sidecar to intercept queries ( we cant add an extra container with Helm Chart values):
```yaml
spec:
  template:
    spec:
      containers:
        ...
        - name: reverse-proxy
          image: medinvention/prometheus-filter-proxy:0.0.1
          args:
            - "run"
            - "--port=9091"
            - "--prometheus-server=http://localhost:9090"
            - "--auth=username:password"
          ports:
            - name: http
              containerPort: 9091
              protocol: TCP
          resources:
            limits:
              cpu: 250m
              memory: 200Mi
```

Then, patch service (prometheus-server, and change port):
```yaml
spec:
    ports:
      - name: http
        port: 80
        protocol: TCP
        targetPort: 9091
```


*For services we need to use the merge patch command of kubectl*

Availabel options:
- `--port`: Proxy port.
- `--prometheus-server`: Prometheus server URL.
- `--auth`: (Optional) username:password authentication informations.

2- Configure Grafana datasource, to set the username and the password, then add your custom filter into the custom query field :
```yaml
filter={<labelname>=<value>}
```

## Build yours

If you want to build it from this repository, follow the instructions below:

```bash
docker build --tag prometheus-filter-proxy:local . -f build/Dockerfile
# For multi-plateform 
# docker buildx build --push --platform linux/arm64,linux/amd64 --tag medinvention/prometheus-filter-proxy:local . -f build/Dockerfile
```