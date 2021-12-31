# mggers-reports-api

### Usage

```
make run
```

### Overwrite config

Create a ``local/env.sh`` <br>

Overwrite variables by exporting ``CONF_<VAR>``

<i>e.g.</i>

```
#!/bin/bash

export CONF_PORT=3005
```

Then source them before running app

````source local/env````