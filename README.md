[![Code Scanning - Action](https://github.com/ka1i/golang-cli-app/actions/workflows/codeScan.yml/badge.svg)](https://github.com/ka1i/golang-cli-app/actions/workflows/codeScan.yml)

# golang cli app

## Build
```bash
make
```

## Usage
```bash
./bin/app-cli -h
```

## Cloud build & deploy
```bash
PrivateRegistry=<private registry>

# Docker Build
docker build --progress=plain -f build/local/Dockerfile -t ${PrivateRegistry}/gca/cli:latest .

# Docker Deploy
docker stack deploy -c deployment/docker/docker-compose.yml gca

# Kubernetes Deploy
kubectl create configmap gca-configs --from-file configs.yaml -n app    # 请提前编辑好配置文件
./deployment/kubernetes/autogen.sh 1 ${PrivateRegistry}                 # 参数1：指定副本集数量。 参数2：指定私有镜像仓库地址
kubectl apply -f deployment/kubernetes/cli-app.yaml -n app              # 开始部署
```