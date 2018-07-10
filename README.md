# check-jolokia
Icinga Checks based on Jolokia results

## Default Flags



| Flag | Short | Description |
| --- | --- | --- |
| COMMAND |  | The command to execute followed by the Jolokia URL incl authentication credentials |
| --help | -h | Help for check-jolokia |
| --verbose | -v | Enable verbose output |

## Commands

### queueAttribute

Query any Attribute of a given queue.

- Search for `AddressMemoryUsage` in queue broker="0.0.0.0" with thresholds
```
check-jolokia queueAttribute http://<username>:<password>@<url>:<port>/console/jolokia/read
-q 'broker="0.0.0.0"' -a AddressMemoryUsage -c 3000000 -w 2500000
```

### Flags

| Flag | Short | Description |
|--- |--- |--- |
| --attribute | -a | the attributes to query from the queue (default "*") |
| --critical | -c | critical threshold for minimum amount of result (default "10:") |
| --domain | -d | the domain of the queue to query (default "org.apache.activemq.artemis") |
| --help | -h | Help for queueAttribute |
| --ok_if_queue_is_missing | -o | The queue to search first and return OK is missing (default "*") |
| --queue | -q | The queue to query (default "*") |
| --warning | -w | Warning threshold for minimum amount of result (default "5:") |


## Thresholds

<https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT>

| Range definition |	Generate an alert if x... |
|--- |--- |
| 10 | < 0 or > 10, (outside the range of {0 .. 10}) |
| 10: | < 10, (outside {10 .. ∞}) |
| ~:10 | > 10, (outside the range of {-∞ .. 10}) |
| 10:20 | < 10 or > 20, (outside the range of {10 .. 20}) |
| @10:20 | ≥ 10 and ≤ 20, (inside the range of {10 .. 20} |