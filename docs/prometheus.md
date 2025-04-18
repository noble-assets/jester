# Prometheus Metrics

By default, Jester's Prometheus server listens on `localhost:2112` for incoming scrape requests.
This can be changed in the config or by using the `--metrics.address` flag.

> [!NOTE]
> If you need to scrape metrics from an external system, set the metrics address to `0.0.0.0:2112`.

All metrics are served at `/metrics`.

## Custom Metrics

| **Exported Metric**                  | **Description**                                                                                                                 | **Type**  |
|--------------------------------------|---------------------------------------------------------------------------------------------------------------------------------|-----------|
| eth_sub_interruption_counter         | Total number of Ethereum subscription interruptions causing Jester to redial the websocket client.                              | Counter   |
| getVoteExtension_counter             | Total number of times `getVoteExtension` is queried. If you are a validator, this should happen each time you are the proposer. | Counter   |
| logMessagePublished_counter          | Total number of times Wormhole's Ethereum event `LogMessagePublished` is observed.                                              | Counter   |
| mTokenSent_counter                   | Total number of times M0's Ethereum event `mTokenSent` is observed.                                                             | Counter   |
| mTokenIndexSent_counter              | Total number of times M0's Ethereum event `mTokenIndexSent` is observed.                                                        | Counter   |
| mailbox_dispatch_counter             | Total number of times Hyperlanes Ethereum event `Dispatch` is observed.                                                         | Counter   |
| vaa_receive_duration_minutes         | Summary of the time it takes wormhole to pick up the VAA in minutes.                                                            | Summary   |
| vaa_found_total                      | Total number of times a VAA was found.                                                                                          | Counter   |
| vaa_failed_total                     | Total number of times fetching a VAA failed.                                                                                    | Counter   |
| vaa_failed_max_attempts_reached      | Total number of times fetching a VAA failed after reaching the maximum number of attempts.                                      | Counter   |
| reorg_detected_counter               | The total number of times a relevant event was included in an Ethereum reorg.                                                   | Counter   |
| reorg_event_recovered_counter        | The total number of times a relevant event was recovered after an Ethereum reorg.                                               | Counter   |
| reorg_event_lost_counter             | The total number of times a relevant event was lost after an Ethereum reorg.                                                    | Counter   |
| finality_reached_seconds_summary     | Summary of the time it takes for a block to reach finality in seconds.                                                          | Summary   |
| average_block_interval_seconds_guage | The average time it takes for a block to be produced in seconds. Update every hour.                                             | Guage     |
