
# Cron schedule scripts



## Compile

To compile the project is necessary golang >= 1.15

```bash
cd path/to/cron-job
```
```
make build
```


## Run

To start the service is only necessary to indicate the directory that contains the configuration files with parameter `--path`. The default value is `/usr/local/etc/cron-job`.

```
mkdir -p /path/to/configs/
```

```bash
bin/cron-service --path=/path/to/configs/
```

## Layout configuration file

```json
//cat /path/to/configs/routine-1.json
{
    "log_file": "/path/to/logs/routine-1.log",
    "scripts": [
        {
            "spec": "*/1 * * * * *",
            "command": ["/usr/bin/python3" , "/path/to/script.py"],
            "timeout": 100,
            "multi_processing_limit": 4
        }
    ]
}
```

# For mode details

- [Spec parameter in script](doc/spec-expression.md)

# Todo

 - [ ] (Documentation)
 - [ ] (CLI for create configuration file )
 - [ ] (Workflow to build binary from other architectures)