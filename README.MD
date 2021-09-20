# EXAMPLE APPLICATION

The Application for teting perpose only.
The application consist of few api backend. You can read more on [README.md](./web/README.md).

## Usage 
You can run application on your local machine, on docker or kubernets cluster
Before using the application, make sure that database is accessable from local, docker or k8s cluster.

### Local PC
[makefile](./makefile) already got scripts and you run it to test application

```bash
make dev
```

### Docker Deployment
To run application on docker, first you are going to build container image
```bash
docker build . -t example-app
```
In the image list you will see the image name
```bash
docker image ls
```
to run container, 

```bash
docker run -d -p 8080:8080 --name appContainer example-app
```

### Upload to Registry

To run container inside of k8s pod, we need build and push image to registry. 
Please meke sure that you have sufficient privileges to push the image to registry.

```bash
docker build . -t REGISTRY/APPNAME:TAG
```

After building the image, push it to registry
```bash
docker push REGISTRY/APPNAME:TAG
```
### Helm Deployment

To run application on k8s, we are going to you [helm-chart](./helm-chart/)
To create pod, k8s worker nodes need to pull images from registry, and in the [values](./helm-chart/values.yaml) you need to change variables values accrdingly
```bash
image:
  repo: REGISTY
  name: APPNAME
  tag: TAG
```

After editing values file, you can deploy application:

```bash
helm upgrade [-n NAMESPACE] --install [RELEASENAME] .
```

## Recommendations
1. JSON format to generate logs. Its easy to store and query.
2. Use external logging system, like ElasticSeach to collect application log.
3. Do not store data(state) in container, if you have to, use external storage, like bucket
4. Do not store credential data on git or container, use Vault or Consul.
5. Use external metrics and alerting system, to monitor application status. Like Prometheus and Grafana
6. Do not run statefull application on pod, like database application