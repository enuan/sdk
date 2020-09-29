package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

// Row represents single row of Grafana dashboard.
type Row struct {
	Title     string        `json:"title"`
	ShowTitle bool          `json:"showTitle"`
	Collapse  bool          `json:"collapse"`
	Editable  bool          `json:"editable"`
	Height    FloatOrString `json:"height"`
	Panels    []Panel       `json:"panels"`
	Repeat    *string       `json:"repeat"`
}

var lastPanelID uint

func (r *Row) Add(panel *Panel) {
	lastPanelID++
	panel.ID = lastPanelID
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddCustom(data *CustomPanel) {
	lastPanelID++
	panel := NewCustom("")
	panel.ID = lastPanelID
	panel.CustomPanel = data
	r.Panels = append(r.Panels, *panel)
}
