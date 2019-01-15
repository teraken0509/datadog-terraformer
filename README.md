# datadog-terraformer
Export existing Datadog resources to HCL

## Usage

### Export Monitor Configuration
```
datadog-terraformer monitor [monitorID]
```

e.g.
- command
    ```
    $ datadog-terraformer monitor 123456789
    ```

- output
    ```output
    resource "datadog_monitor" "monitor_123456789" {
        name               = "sample monitor"
        type               = "query alert"
        message            = "sample message"

        query = "sample query"

        ...

        tags = [
            "sample",
        ]
    }
    ```

## License

Datadog Terraformer is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/kterada0509/datadog-terraformer/blob/master/LICENSE)
