
## tests

```
kubectl get secret postgres.pg-minimal-cluster.credentials -o 'jsonpath={.data.password}' | base64 -D
kubectl get secret bygui.pg-minimal-cluster.credentials -o 'jsonpath={.data.password}' | base64 -D
```

```
export PGHOST="pg-minimal-cluster"
export PGHOST="pg-minimal-cluster-repl"
export PGHOST="10.1.0.40"
export PGHOST="10.1.0.41"
export PGPORT="5432"
postgres
export PGPASSWORD="gmvA4XLq6xVcEidSWVcD0LjDQT4wYjt39XTWEcLWgwZSnj1UJoWeaqjhoxeAMu3a"
bygui
export PGPASSWORD="EHvWYv4KH2l3enAicJ3QVWsjyyhFBxihij4PHUPxxszWGgLJ15cA2F3xQuopoCrX"
```

```
psql -U postgres
psql -U bygui
```
