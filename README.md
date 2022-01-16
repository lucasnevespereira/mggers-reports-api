# mggers-reports-api

### Usage

```
make start
```

```
make stop
```

##### Docker logs
```
make api-logs
```

```
make db-logs
```

### Use Envs

Create a `/local` directory with a `env.sh` and export an env variable
```
export test_env=postgres
```

Source file
```
source local/env.sh
```

Use in code
```
os.GetEnv("test_env")
```
