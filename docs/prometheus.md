# Prometheus Metrics

By default, Jesters Prometheus server listens on `localhost:2112` for incoming scrape requets.
This can be changed in the config OR by using the `--metrics.address` flag.

Default metrics are available at `/metrics`.

Custom metrics can be accessed at `/jester/metrics`.

## Custom Metrics

| **Exported Metric**          | **Description**                                                                                   | **Type** |
|------------------------------|---------------------------------------------------------------------------------------------------|----------|
| eth_sub_interruption_counter | Total number of Ethereum subscription interruptions causing Jester to redial the websocket client | Counter  |