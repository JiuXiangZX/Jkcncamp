kubectl create ns test
sleep 10
kubectl apply -n test -f httpserver-configmap.yaml
kubectl apply -n test -f deployment.yaml
kubectl apply -n test -f service.yaml
kubectl apply -n test -f nginx-ingress.yaml
sleep 3m
kubectl apply -n test -f ingress.yaml