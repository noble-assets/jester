# Prometheus Metrics

By default, Jester's Prometheus server listens on `localhost:2112` for incoming scrape requests.
This can be changed in the config or by using the `--metrics.address` flag.

> [!NOTE]
> If you need to scrape metrics from an external system, set the metrics address to `0.0.0.0:2112`.

All metrics are served at `/metrics`.

## Custom Metrics

| **Exported Metric**          | **Description**                                                                                   | **Type** |
|------------------------------|---------------------------------------------------------------------------------------------------|----------|
| eth_sub_interruption_counter | Total number of Ethereum subscription interruptions causing Jester to redial the websocket client | Counter  |
