#!/bin/bash --posix

replicas=$1
repository=$2

replicas=${replicas:=1}
repository=${repository:=empty.repository}

cat > ./deployment/kubernetes/cli-app.yaml << EOF
kind: ConfigMap
apiVersion: v1
metadata:
  name: gca-environment
data:
  GCA_APP_MODE: "debug"
  GCA_APP_PORT: "10086"

---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: gca-data
spec:
  accessModes:
    - ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 32Gi

---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: gca-cli
  labels:
    app: gca-cli

spec:
  replicas:  ${replicas}
  selector:
    matchLabels:
      app: gca-cli
  template:
    metadata:
      labels:
        app: gca-cli
    spec:
      containers:
        - name: gca-cli
          image: ${repository}/gca/cli:latest
          ports:
            - name: web
              containerPort: 10086
          imagePullPolicy: Always
          envFrom:
          - configMapRef:
              name: gca-environment
          volumeMounts:
          - name: configs
            readOnly: true
            subPath: "app.yaml"
            mountPath: "/opt/configs/app.yaml"
          - name: data
            mountPath: "/opt/share"
          readinessProbe:
            tcpSocket:
              port: 8080
            timeoutSeconds: 10
            periodSeconds: 10
            successThreshold: 1
            failureThreshold: 3
      volumes:
        - name: data
          persistentVolumeClaim:
            claimName: gca-data
        - name: configs
          configMap:
            name: gca-configs
            items:
              - key: "app.yaml"
                path: "app.yaml"

---
apiVersion: v1
kind: Service
metadata:
  name: gca-cli
spec:
  ports:
    - name: http
      port: 10086
      nodePort: 30801
  selector:
    app: gca-cli
  type: NodePort
EOF
