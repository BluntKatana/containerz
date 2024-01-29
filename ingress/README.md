```sudo microk8s kubectl create secret tls ca-tls-secret --cert=cert.pem --key=key.pem -n cert-manager```

Add to `/etc/hosts`:
- `127.0.0.1    y.com`
- `127.0.0.1    api.y.com`