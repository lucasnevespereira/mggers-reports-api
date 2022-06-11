# mggers-reports-api

### Usage

```
source local/env.sh
```

```
make run
```

### Set Envs

Create a `/local` directory with a `env.sh` and export an env variable
```
export ENV=dev
export MONGO_SERVER="mongodb://localhost:27017"
```

Source file
```
source local/env.sh
```

Use in code
```
os.GetEnv("ENV")
os.GetEnv("MONGO_SERVER")
```
