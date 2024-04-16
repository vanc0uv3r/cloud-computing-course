#/bin/bash
docker build -f ./db/Dockerfile -t redis-k8s ./db
docker build -f ./server/Dockerfile -t server-k8s ./server

kubectl apply -f redis-db.yaml
kubectl apply -f go-server.yaml
