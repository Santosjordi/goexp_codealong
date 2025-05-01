

K8S configuration:

```yml
kubectl create secret docker-registry ghcr-creds \
  --docker-server=ghcr.io \
  --docker-username=santosjordi \
  --docker-password=github_pat \
  --docker-email=santos@email.com
```