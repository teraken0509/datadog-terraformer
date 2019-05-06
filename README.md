# datadog-terraformer
Export existing Datadog resources to HCL

[![Build Status](https://travis-ci.org/kterada0509/datadog-terraformer.svg?branch=master)](https://travis-ci.org/kterada0509/datadog-terraformer)

## Installing the CLI

### MacOS X with Homebrew

```
brew install kterada0509/tap/datadog-terraformer
```

### Download Packages

Download from [releases page](https://github.com/kterada0509/datadog-terraformer/releases).

## Supported Resources

* [x] [datadog_monitor](https://www.terraform.io/docs/providers/datadog/r/monitor.html)
* [ ] [datadog_integration_aws](https://www.terraform.io/docs/providers/datadog/r/integration_aws.html)
* [ ] [datadog_integration_gcp](https://www.terraform.io/docs/providers/datadog/r/integration_gcp.html)
* [ ] [datadog_metric_metadata](https://www.terraform.io/docs/providers/datadog/r/metric_metadata.html)
* [ ] [datadog_screenboard](https://www.terraform.io/docs/providers/datadog/r/screenboard.html)
* [ ] [datadog_synthetics_test](https://www.terraform.io/docs/providers/datadog/r/synthetics.html)
* [x] [datadog_timeboard](https://www.terraform.io/docs/providers/datadog/r/timeboard.html)
* [x] [datadog_user](https://www.terraform.io/docs/providers/datadog/r/user.html)
* [x] [datadog_downtime](https://www.terraform.io/docs/providers/datadog/r/downtime.html)

## Usage

<details>
<summary>Export Monitor Configuration</summary>

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
</details>

<details>
<summary>Export Timeboard Configuration</summary>

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
</details>

<details>
<summary>Export User Configuration</summary>

```
datadog-terraformer user [dashboard user handler]
```

e.g.
- command
    ```
    $ datadog-terraformer user test@example.co.jp
    ```

- output
    ``` output
    resource "datadog_user" "user_handle" {
        disabled = false
        email  = "test@example.co.jp"
        handle = "test@example.co.jp"
        name   = "Test user"
        //Deprecated
        is_admin = false
    }
    ```
</details>

## License

Datadog Terraformer is released under the Apache 2.0 license. See [LICENSE.txt](./LICENSE)
