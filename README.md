<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


    Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Snap collector plugin - csvreader
This plugin collects log messages partially for each collection run. Log file reading is limited by time.

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [golang 1.7+](https://golang.org/dl/) (needed only for building)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-csvreader

Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-csvreader.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

#### Task manifest configuration options
Option|Description|Default value
------|-----------|-------------
"file"|Declaration of full path file name, in CSV format
"indexes"|Defines the list of column indexes that want to collect, separated by comma
"indexes"|The units corresponding to the indexes of each metric

## Documentation

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description
----------|-----------------------
/intel/csvreader/[index]/index|Single column metric


### Examples
This is an example running csvreader collector and writing data to a file. It is assumed that you are using the latest Snap binary and plugins.

The example is run from a directory which includes snaptel, snapteld, along with the plugins and task file.

In one terminal window, open the Snap daemon (in this case with logging set to 1 and trust disabled):
```
$ snapteld -l 1 -t 0
```

In another terminal window:

Download csvreader collector plugin:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-csvreader/latest/linux/x86_64/snap-plugin-collector-csvreader
```

Load csvreader plugin
```
$ snaptel plugin load snap-plugin-collector-csvreader
Plugin loaded
Name: csvreader
Version: 1
Type: collector
Signed: false
Loaded Time: Thu, 05 Jan 2017 11:58:11 CET
```
See available metrics for your system
```
$ snaptel metric list
```

Create a task manifest file (e.g. `task-csvreader.json`):    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "3s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/csvreader/*": {}
            },
            "config": {
                "/intel/csvreader": {
                    "file": "/opt/snap/files/test.csv",
                    "indexes": "0,1",
                    "unit": "u1,u2"
                }
            },
            "publish": [
                {
                    "plugin_name": "file",
                    "config": {
                        "file": "/tmp/published_csvreader"
                    }
                }
            ]
        }
    }
}
```

Download file publisher plugin:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
```
Load file plugin for publishing:
```
$ snaptel plugin load snap-plugin-publisher-file
Plugin loaded
Name: file
Version: 4
Type: publisher
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:41:39 PST
```

Create task:
```
$ snaptel task create -t task-csvreader.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See file output (this is just single collection output with default collection_time of 300ms): [EXAMPLE_OUTPUT.md](EXAMPLE_OUTPUT.md)

Stop task:
```
$ snaptel task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-csvreader/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-csvreader/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to bcsvreader to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@cuongquay](https://github.com/cuongquay)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
