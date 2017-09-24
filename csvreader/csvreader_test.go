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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadConfig(t *testing.T) {
	cfg := make(plugin.Config)
	cfg["source"] = "/var/cache/snap"
	cfg["indexes"] = "0,1"
	cfg["units"] = "max,min"

	Convey("should not panic", t, func() {
		So(func() {
			l := CSVReader{}
			l.getConfig(cfg)
		}, ShouldNotPanic)
	})

	Convey("should load config properly", t, func() {
		l := CSVReader{}
		l.getConfig(cfg)

		So(len(l.configInt), ShouldEqual, 0)
		So(len(l.configStr), ShouldEqual, 3)
	})
}

func TestGetMetricTypes(t *testing.T) {
	cfg := make(plugin.Config)
	cfg["source"] = "/opt/snap/files/metrics.csv"
	cfg["indexes"] = "0,1,2"
	cfg["units"] = "u1,u2,u3"

	Convey("should not panic", t, func() {
		So(func() {
			l := CSVReader{}
			l.GetMetricTypes(cfg)
		}, ShouldNotPanic)
	})

	Convey("should return valid metric type", t, func() {
		l := CSVReader{}
		mt, err := l.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(mt, ShouldNotBeEmpty)
		So(len(mt), ShouldEqual, 1)
		So(mt[0].Namespace.Strings(), ShouldResemble, []string{"intel", "csvreader", "*", "*"})
	})
}

func makeMetric(metricName string, cfg plugin.Config) []plugin.Metric {
	mts := []plugin.Metric{
		plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "csvreader").AddDynamicElement("Index", "Index name defined by indexes").AddDynamicElement("Source", "Source of metrics"),
			Config:    cfg,
		},
	}
	mts[0].Namespace[2].Value = metricName

	return mts
}

func joinMetricData(m []plugin.Metric) string {
	allData := ""
	for _, v := range m {
		allData += v.Data.(string)
	}
	return allData
}

func TestCollectMetrics(t *testing.T) {
	Convey("should not panic and return valid namespace", t, func() {
		cfgIndex01 := make(plugin.Config)
		cfgIndex01["source"] = "/opt/snap/files/metrics.csv"
		cfgIndex01["indexes"] = "0,1"
		cfgIndex01["units"] = "date,psi"
		mtsIndex0 := makeMetric("0", cfgIndex01)
		l := CSVReader{}
		mts, err := l.CollectMetrics(mtsIndex0)
		So(err, ShouldBeNil)
		for _, m := range mts {
			ok, dyn := m.Namespace.IsDynamic()
			So(ok, ShouldBeTrue)
			So(dyn, ShouldResemble, []int{2, 3})
		}
	})
}

func TestGetConfigPolicy(t *testing.T) {
	Convey("should not panic", t, func() {
		So(func() {
			l := CSVReader{}
			l.GetConfigPolicy()
		}, ShouldNotPanic)
	})

	Convey("should resemble default config policy", t, func() {
		defaultPolicy := plugin.NewConfigPolicy()
		defaultPolicy.AddNewStringRule([]string{"intel", Name}, "source", false, plugin.SetDefaultString("/opt/snap/files/metrics.csv"))
		defaultPolicy.AddNewStringRule([]string{"intel", Name}, "indexes", false, plugin.SetDefaultString("0,1"))
		defaultPolicy.AddNewStringRule([]string{"intel", Name}, "units", false, plugin.SetDefaultString("unit,unit"))

		l := CSVReader{}
		policy, err := l.GetConfigPolicy()
		So(err, ShouldBeNil)
		So(policy, ShouldResemble, *defaultPolicy)
	})
}
