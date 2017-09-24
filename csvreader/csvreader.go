/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 CuongQuay <cuong3ihut@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csvreader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	// Name of plugin
	Name = "csvreader"
	// Version of plugin
	Version = 1
)

// CSVReader collector implementation used for testing
type CSVReader struct {
	configStr       map[string]string
	configInt       map[string]int64
	currentRowIndex int
}

// New logs collector plugin instance
func New() *CSVReader {
	return &CSVReader{}
}

// CollectMetrics collects metrics for testing
func (c *CSVReader) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	if err := c.getConfig(mts[0].Config); err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{"ColumnIndex": c.configStr["Index"], "source": c.configStr["source"]}).Info("CollectMetrics")
	metrics := []plugin.Metric{}
	var strVal string
	var sliceStr []string
	var headers []string

	csvPath := c.configStr["source"]
	if csvFile, err := os.Open(csvPath); err == nil {
		defer csvFile.Close()
		csvReader := csv.NewReader(csvFile)
		headers, err = csvReader.Read()
		pos := 1
		for {
			sliceStr, err = csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			if c.currentRowIndex == pos {
				break
			}
			pos++
		}
		c.currentRowIndex++

		indexes := strings.Split(c.configStr["attrs"], ",")
		units := strings.Split(c.configStr["units"], ",")
		fileMetrics := []plugin.Metric{}
		for i, colIdxStr := range indexes {
			// Add new metric base on task configuration
			ns := make([]plugin.NamespaceElement, len(mts[0].Namespace))
			copy(ns, mts[0].Namespace)
			colIdxStr = strings.TrimSpace(colIdxStr)
			if colIdx, err := strconv.Atoi(colIdxStr); err == nil {
				ns[2].Value = colIdxStr
				ns[3].Value = csvPath
				mt := plugin.Metric{
					Data:        0.1,
					Unit:        units[i],
					Description: headers[colIdx],
					Timestamp:   time.Now(),
					Version:     Version,
					Namespace:   ns,
				}
				if strVal, err = c.getColumnFromCSV(sliceStr, colIdx); err == nil {
					if mt.Data, err = strconv.ParseFloat(strVal, 64); err == nil {
						fileMetrics = append(fileMetrics, mt)
					}
				}
			}
		}
		// Return file metrics only if cache file successfully saved
		metrics = append(metrics, fileMetrics...)
	}
	// Return file metrics only if cache file successfully saved
	return metrics, nil
}

//GetMetricTypes returns metric types for testing
func (c *CSVReader) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	if err := c.getConfig(cfg); err != nil {
		return nil, err
	}

	mts := []plugin.Metric{}
	mts = append(mts, plugin.Metric{
		Namespace:   plugin.NewNamespace("intel", Name).AddDynamicElement("Index", "Index of column").AddDynamicElement("Source", "Source of metrics").AddStaticElement("message"),
		Description: "Single Column Index",
		Unit:        "string",
		Version:     Version,
	})

	return mts, nil
}

//GetConfigPolicy returns a ConfigPolicy for testing
func (c *CSVReader) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{"intel", Name}, "source", false, plugin.SetDefaultString("/opt/snap/files/metrics.csv"))
	policy.AddNewStringRule([]string{"intel", Name}, "attrs", false, plugin.SetDefaultString("0,1"))
	policy.AddNewStringRule([]string{"intel", Name}, "units", false, plugin.SetDefaultString("unit,unit"))
	return *policy, nil
}

// Load config values
func (c *CSVReader) getConfig(cfg plugin.Config) error {
	c.configStr = make(map[string]string)
	c.configInt = make(map[string]int64)

	for key := range cfg {
		if val, err := cfg.GetInt(key); err == nil {
			c.configInt[key] = val
		}
		if val, err := cfg.GetString(key); err == nil {
			c.configStr[key] = val
		}
	}

	return nil
}

func getSliceValueAtIndex(sliceStr []string, idx int) (string, error) {
	for k, v := range sliceStr {
		if k == idx {
			return v, nil
		}
	}
	return "", fmt.Errorf("Value not found at %d", idx)
}

func (c *CSVReader) getColumnFromCSV(sliceStr []string, colIdx int) (string, error) {
	var strEmptyVal string = ""
	for k, v := range sliceStr {
		if k == colIdx {
			return trimQuotes(v), nil
		}
	}
	return trimQuotes(strEmptyVal), fmt.Errorf("Value not found at %d", colIdx)
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}
