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

### Export Timeboard Configuration
```
datadog-terraformer timeboard [dashboard id]
```

e.g.
- command
    ```
    $ datadog-terraformer timeboard 123456789
    ```

- output
    ``` output
    resource "datadog_timeboard" "timeboard_123456789" {
        title       = "Sample Dashboard"
        description = "sample dashboard"
        read_only   = true


        graph {
            title = "Sample Graph"
            viz   = "timeseries"
            request {
                q    = "query1"
                type = "line"

            }
            request {
                q    = "query2"
                type = "line"

            }
            request {
                q    = "query3"
                type = "line"

            }

        }

        ...

    }
    ```

## License

Datadog Terraformer is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/kterada0509/datadog-terraformer/blob/master/LICENSE)
