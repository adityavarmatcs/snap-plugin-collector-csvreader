# Snap logs collector plugin example output
Example output from Snap logs collector plugin generated using Snap elasticsearch publisher plugin.

```json
[
         {
            "_index": "esp-telemetry-data-2017.08.24",
            "_type": "message",
            "_id": "AV4SaYkV6ykml-1WEbF8",
            "_score": 1,
            "_source": {
               "Data": 711,
               "Index": "3",
               "Namespace": "intel/csvreader/3/index",
               "Timestamp": "2017-08-24T04:04:19.853793351Z",
               "Unit": "min",
               "plugin_running_on": "db598704f49f"
            }
         },
         {
            "_index": "esp-telemetry-data-2017.08.24",
            "_type": "message",
            "_id": "AV4SaYFF6ykml-1WEbF1",
            "_score": 1,
            "_source": {
               "Data": 709,
               "Index": "2",
               "Namespace": "intel/csvreader/2/index",
               "Timestamp": "2017-08-24T04:04:17.852712191Z",
               "Unit": "max",
               "plugin_running_on": "db598704f49f"
            }
         },
         {
            "_index": "esp-telemetry-data-2017.08.24",
            "_type": "message",
            "_id": "AV4SaXoR6ykml-1WEbFw",
            "_score": 1,
            "_source": {
               "Data": 714,
               "Index": "3",
               "Namespace": "intel/csvreader/3/index",
               "Timestamp": "2017-08-24T04:04:15.849626881Z",
               "Unit": "min",
               "plugin_running_on": "db598704f49f"
            }
         }
]
```
