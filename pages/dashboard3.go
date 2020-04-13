package pages

import (
	"github.com/GoAdminGroup/components/echarts"
	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/charts"
	"math"
	"math/rand"
	"time"
)

func GetDashBoard3Content(ctx *gin.Context) (types.Panel, error) {

	components := template2.Get(config.GetTheme())
	colComp := components.Col()

	echart := echarts.NewChart()

	line := charts.NewLine()
	line.AddXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
		AddYAxis("map", []float32{19.9, 16.8, 19.9, 29.4, 61.3, 77.3, 93.0},
			charts.LabelTextOpts{Show: true, Position: "bottom"}).
		AddYAxis("slice", []float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131},
			charts.LabelTextOpts{Show: true, Position: "top"})
	line.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "平均值", Type: "average"},
		charts.LineOpts{Smooth: true},
		charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{a}: {b}"}},
	)
	line.SetGlobalOptions(
		charts.YAxisOpts{Name: "搜索时间(ns)", SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "元素数量"})
	line.Width = "250px"
	line.Height = "250px"

	ecbox := components.Box().WithHeadBorder().SetHeader("查询时间对比 哈希表 vs 二分查找").SetBody(echart.SetContent(line).GetContent()).GetContent()

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
	)
	bar.AddXAxis([]string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}).
		AddYAxis("商家A", randInt(), charts.BarOpts{YAxisIndex: 0}).
		AddYAxis("商家B", randInt(), charts.BarOpts{YAxisIndex: 1})
	bar.ExtendYAxis(charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/月"}})
	bar.Width = "250px"
	bar.Height = "250px"

	ecbox2 := components.Box().WithHeadBorder().SetHeader("Bar-多 Y 轴").SetBody(echart.SetContent(bar).GetContent()).GetContent()

	ecboxCol1 := colComp.SetSize(types.SizeMD(3)).SetContent(ecbox).GetContent()
	ecboxCol2 := colComp.SetSize(types.SizeMD(3)).SetContent(ecbox2).GetContent()

	bar3d := charts.NewBar3D()
	bar3d.SetGlobalOptions(
		charts.VisualMapOpts{
			Range:      []float32{0, 30},
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        30,
		},
		charts.Grid3DOpts{BoxDepth: 80, BoxWidth: 200},
	)
	bar3d.AddXYAxis(hours, days).AddZAxis("bar3d", genBar3dData())
	bar3d.Width = "250px"
	bar3d.Height = "250px"

	ecbox3 := components.Box().WithHeadBorder().SetHeader("Bar3D-示例图").SetBody(echart.SetContent(bar3d).GetContent()).GetContent()

	bp := charts.NewBoxPlot()
	bp.AddXAxis(bpX).AddYAxis("boxplot", bpY)
	bp.Width = "250px"
	bp.Height = "250px"

	ecbox4 := components.Box().WithHeadBorder().SetHeader("BoxPlot-示例图").SetBody(echart.SetContent(bp).GetContent()).GetContent()

	ecboxCol3 := colComp.SetSize(types.SizeMD(3)).SetContent(ecbox3).GetContent()
	ecboxCol4 := colComp.SetSize(types.SizeMD(3)).SetContent(ecbox4).GetContent()

	row1 := components.Row().SetContent(ecboxCol1 + ecboxCol2 + ecboxCol3 + ecboxCol4).GetContent()

	es := charts.NewEffectScatter()
	es.AddXAxis(nameItems).AddYAxis("es1", randInt())
	es.Width = "250px"
	es.Height = "250px"

	ecboxCol5 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("EffectScatter-示例图").SetBody(echart.SetContent(es).GetContent()).GetContent()).GetContent()

	funnel := charts.NewFunnel()
	funnel.Add("funnel", genKvData())
	funnel.Width = "250px"
	funnel.Height = "250px"

	ecboxCol6 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Funnel-示例图").SetBody(echart.SetContent(funnel).GetContent()).GetContent()).GetContent()

	gauge := charts.NewGauge()
	m := make(map[string]interface{})
	m["工作进度"] = rand.Intn(50)
	gauge.Add("gauge", m)
	gauge.Width = "250px"
	gauge.Height = "250px"

	ecboxCol7 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Gauge-示例图").SetBody(echart.SetContent(gauge).GetContent()).GetContent()).GetContent()

	geo := charts.NewGeo("china")
	geo.Add("geo", charts.ChartType.EffectScatter, mapData,
		charts.RippleEffectOpts{Period: 4, Scale: 6, BrushType: "stroke"})
	geo.Width = "250px"
	geo.Height = "250px"

	ecboxCol8 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Geo-示例图").SetBody(echart.SetContent(geo).GetContent()).GetContent()).GetContent()

	row2 := components.Row().SetContent(ecboxCol5 + ecboxCol6 + ecboxCol7 + ecboxCol8).GetContent()

	graph := charts.NewGraph()
	graph.Add("graph", graphNodes, genLinks(),
		charts.GraphOpts{Force: charts.GraphForce{Repulsion: 8000}},
	)
	graph.Width = "250px"
	graph.Height = "250px"

	ecboxCol9 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Graph-示例图").SetBody(echart.SetContent(graph).GetContent()).GetContent()).GetContent()

	hm := charts.NewHeatMap()
	hm.AddXAxis(hours).AddYAxis("heatmap", genHeatMapData())
	hm.SetGlobalOptions(
		charts.YAxisOpts{Data: days, Type: "category", SplitArea: charts.SplitAreaOpts{Show: true}},
		charts.XAxisOpts{Type: "category", SplitArea: charts.SplitAreaOpts{Show: true}},
		charts.VisualMapOpts{Calculable: true, Max: 10, Min: 0,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	hm.Width = "250px"
	hm.Height = "250px"

	ecboxCol10 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("HeatMap-示例图").SetBody(echart.SetContent(hm).GetContent()).GetContent()).GetContent()

	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}

	kline.AddXAxis(x).AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	kline.Width = "250px"
	kline.Height = "250px"

	ecboxCol11 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Kline-示例图").SetBody(echart.SetContent(kline).GetContent()).GetContent()).GetContent()

	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        30,
		},
	)
	line3d.AddZAxis("line3D", genLine3dData())
	line3d.Width = "250px"
	line3d.Height = "250px"

	ecboxCol12 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Line3D-示例图").SetBody(echart.SetContent(line3d).GetContent()).GetContent()).GetContent()

	row3 := components.Row().SetContent(ecboxCol9 + ecboxCol10 + ecboxCol11 + ecboxCol12).GetContent()

	liquid := charts.NewLiquid()
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LiquidOpts{IsWaveAnimation: true},
	)
	liquid.Width = "250px"
	liquid.Height = "250px"

	ecboxCol13 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Liquid-示例图").SetBody(echart.SetContent(liquid).GetContent()).GetContent()).GetContent()

	mc := charts.NewMap("china")
	mc.Add("map", mapData)
	mc.Width = "250px"
	mc.Height = "250px"

	ecboxCol14 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Map-示例图").SetBody(echart.SetContent(mc).GetContent()).GetContent()).GetContent()

	pie := charts.NewPie()
	pie.Add("pie", genKvData())
	pie.Width = "250px"
	pie.Height = "250px"

	ecboxCol15 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Pie-示例图").SetBody(echart.SetContent(pie).GetContent()).GetContent()).GetContent()

	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.RadarComponentOpts{
			Indicator: indicators,
			SplitLine: charts.SplitLineOpts{Show: true},
			SplitArea: charts.SplitAreaOpts{Show: true},
		},
	)
	radar.Add("北京", radarDataBJ)
	radar.Width = "250px"
	radar.Height = "250px"

	ecboxCol16 := colComp.SetSize(types.SizeMD(3)).SetContent(components.Box().WithHeadBorder().
		SetHeader("Radar-示例图").SetBody(echart.SetContent(radar).GetContent()).GetContent()).GetContent()

	row4 := components.Row().SetContent(ecboxCol13 + ecboxCol14 + ecboxCol15 + ecboxCol16).GetContent()

	return types.Panel{
		Content:     row1 + row2 + row3 + row4,
		Title:       "Echarts Dashboard",
		Description: "echarts dashboard example",
	}, nil
}

var (
	nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}

	seed       = rand.NewSource(time.Now().UnixNano())
	rangeColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}

	hours = [...]string{
		"12a", "1a", "2a", "3a", "4a", "5a", "6a", "7a", "8a", "9a", "10a", "11a",
		"12p", "1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p", "10p", "11p",
	}

	days = [...]string{"Saturday", "Friday", "Thursday", "Wednesday", "Tuesday", "Monday", "Sunday"}

	bpX = [...]string{"expr1", "expr2", "expr3", "expr4", "expr5"}
	bpY = [][]int{
		{850, 740, 900, 1070, 930, 850, 950, 980, 980, 880,
			1000, 980, 930, 650, 760, 810, 1000, 1000, 960, 960},
		{960, 940, 960, 940, 880, 800, 850, 880, 900, 840,
			830, 790, 810, 880, 880, 830, 800, 790, 760, 800},
		{880, 880, 880, 860, 720, 720, 620, 860, 970, 950,
			880, 910, 850, 870, 840, 840, 850, 840, 840, 840},
		{890, 810, 810, 820, 800, 770, 760, 740, 750, 760,
			910, 920, 890, 860, 880, 720, 840, 850, 850, 780},
		{890, 840, 780, 810, 760, 810, 790, 810, 820, 850,
			870, 870, 810, 740, 810, 940, 950, 800, 810, 870},
	}

	mapData = map[string]float32{
		"北京":   float32(rand.Intn(150)),
		"上海":   float32(rand.Intn(150)),
		"深圳":   float32(rand.Intn(150)),
		"辽宁":   float32(rand.Intn(150)),
		"青岛":   float32(rand.Intn(150)),
		"山西":   float32(rand.Intn(150)),
		"陕西":   float32(rand.Intn(150)),
		"乌鲁木齐": float32(rand.Intn(150)),
		"齐齐哈尔": float32(rand.Intn(150)),
	}

	maxNum = 50

	graphNodes = []charts.GraphNode{
		{Name: "节点1"},
		{Name: "节点2"},
		{Name: "节点3"},
		{Name: "节点4"},
		{Name: "节点5"},
		{Name: "节点6"},
		{Name: "节点7"},
		{Name: "节点8"},
	}

	hmData = [][3]int{
		{0, 0, 5}, {0, 1, 1}, {0, 2, 0}, {0, 3, 0}, {0, 4, 0}, {0, 5, 0},
		{0, 6, 0}, {0, 7, 0}, {0, 8, 0}, {0, 9, 0}, {0, 10, 0}, {0, 11, 2},
		{0, 12, 4}, {0, 13, 1}, {0, 14, 1}, {0, 15, 3}, {0, 16, 4}, {0, 17, 6},
		{0, 18, 4}, {0, 19, 4}, {0, 20, 3}, {0, 21, 3}, {0, 22, 2}, {0, 23, 5},
		{1, 0, 7}, {1, 1, 0}, {1, 2, 0}, {1, 3, 0}, {1, 4, 0}, {1, 5, 0},
		{1, 6, 0}, {1, 7, 0}, {1, 8, 0}, {1, 9, 0}, {1, 10, 5}, {1, 11, 2},
		{1, 12, 2}, {1, 13, 6}, {1, 14, 9}, {1, 15, 11}, {1, 16, 6}, {1, 17, 7},
		{1, 18, 8}, {1, 19, 12}, {1, 20, 5}, {1, 21, 5}, {1, 22, 7}, {1, 23, 2},
		{2, 0, 1}, {2, 1, 1}, {2, 2, 0}, {2, 3, 0}, {2, 4, 0}, {2, 5, 0}, {2, 6, 0},
		{2, 7, 0}, {2, 8, 0}, {2, 9, 0}, {2, 10, 3}, {2, 11, 2}, {2, 12, 1}, {2, 13, 9},
		{2, 14, 8}, {2, 15, 10}, {2, 16, 6}, {2, 17, 5}, {2, 18, 5}, {2, 19, 5},
		{2, 20, 7}, {2, 21, 4}, {2, 22, 2}, {2, 23, 4}, {3, 0, 7}, {3, 1, 3},
		{3, 2, 0}, {3, 3, 0}, {3, 4, 0}, {3, 5, 0}, {3, 6, 0}, {3, 7, 0},
		{3, 8, 1}, {3, 9, 0}, {3, 10, 5}, {3, 11, 4}, {3, 12, 7}, {3, 13, 14},
		{3, 14, 13}, {3, 15, 12}, {3, 16, 9}, {3, 17, 5}, {3, 18, 5}, {3, 19, 10},
		{3, 20, 6}, {3, 21, 4}, {3, 22, 4}, {3, 23, 1}, {4, 0, 1}, {4, 1, 3},
		{4, 2, 0}, {4, 3, 0}, {4, 4, 0}, {4, 5, 1}, {4, 6, 0}, {4, 7, 0},
		{4, 8, 0}, {4, 9, 2}, {4, 10, 4}, {4, 11, 4}, {4, 12, 2}, {4, 13, 4},
		{4, 14, 4}, {4, 15, 14}, {4, 16, 12}, {4, 17, 1}, {4, 18, 8}, {4, 19, 5},
		{4, 20, 3}, {4, 21, 7}, {4, 22, 3}, {4, 23, 0}, {5, 0, 2}, {5, 1, 1},
		{5, 2, 0}, {5, 3, 3}, {5, 4, 0}, {5, 5, 0}, {5, 6, 0}, {5, 7, 0}, {5, 8, 2},
		{5, 9, 0}, {5, 10, 4}, {5, 11, 1}, {5, 12, 5}, {5, 13, 10}, {5, 14, 5},
		{5, 15, 7}, {5, 16, 11}, {5, 17, 6}, {5, 18, 0}, {5, 19, 5}, {5, 20, 3},
		{5, 21, 4}, {5, 22, 2}, {5, 23, 0}, {6, 0, 1}, {6, 1, 0}, {6, 2, 0},
		{6, 3, 0}, {6, 4, 0}, {6, 5, 0}, {6, 6, 0}, {6, 7, 0}, {6, 8, 0},
		{6, 9, 0}, {6, 10, 1}, {6, 11, 0}, {6, 12, 2}, {6, 13, 1}, {6, 14, 3},
		{6, 15, 4}, {6, 16, 0}, {6, 17, 0}, {6, 18, 0}, {6, 19, 0}, {6, 20, 1},
		{6, 21, 2}, {6, 22, 2}, {6, 23, 6},
	}

	kd = [...]klineData{
		{date: "2018/1/24", data: [4]float32{2320.26, 2320.26, 2287.3, 2362.94}},
		{date: "2018/1/25", data: [4]float32{2300, 2291.3, 2288.26, 2308.38}},
		{date: "2018/1/28", data: [4]float32{2295.35, 2346.5, 2295.35, 2346.92}},
		{date: "2018/1/29", data: [4]float32{2347.22, 2358.98, 2337.35, 2363.8}},
		{date: "2018/1/30", data: [4]float32{2360.75, 2382.48, 2347.89, 2383.76}},
		{date: "2018/1/31", data: [4]float32{2383.43, 2385.42, 2371.23, 2391.82}},
		{date: "2018/2/1", data: [4]float32{2377.41, 2419.02, 2369.57, 2421.15}},
		{date: "2018/2/4", data: [4]float32{2425.92, 2428.15, 2417.58, 2440.38}},
		{date: "2018/2/5", data: [4]float32{2411, 2433.13, 2403.3, 2437.42}},
		{date: "2018/2/6", data: [4]float32{2432.68, 2434.48, 2427.7, 2441.73}},
		{date: "2018/2/7", data: [4]float32{2430.69, 2418.53, 2394.22, 2433.89}},
		{date: "2018/2/8", data: [4]float32{2416.62, 2432.4, 2414.4, 2443.03}},
		{date: "2018/2/18", data: [4]float32{2441.91, 2421.56, 2415.43, 2444.8}},
		{date: "2018/2/19", data: [4]float32{2420.26, 2382.91, 2373.53, 2427.07}},
		{date: "2018/2/20", data: [4]float32{2383.49, 2397.18, 2370.61, 2397.94}},
		{date: "2018/2/21", data: [4]float32{2378.82, 2325.95, 2309.17, 2378.82}},
		{date: "2018/2/22", data: [4]float32{2322.94, 2314.16, 2308.76, 2330.88}},
		{date: "2018/2/25", data: [4]float32{2320.62, 2325.82, 2315.01, 2338.78}},
		{date: "2018/2/26", data: [4]float32{2313.74, 2293.34, 2289.89, 2340.71}},
		{date: "2018/2/27", data: [4]float32{2297.77, 2313.22, 2292.03, 2324.63}},
		{date: "2018/2/28", data: [4]float32{2322.32, 2365.59, 2308.92, 2366.16}},
		{date: "2018/3/1", data: [4]float32{2364.54, 2359.51, 2330.86, 2369.65}},
		{date: "2018/3/4", data: [4]float32{2332.08, 2273.4, 2259.25, 2333.54}},
		{date: "2018/3/5", data: [4]float32{2274.81, 2326.31, 2270.1, 2328.14}},
		{date: "2018/3/6", data: [4]float32{2333.61, 2347.18, 2321.6, 2351.44}},
		{date: "2018/3/7", data: [4]float32{2340.44, 2324.29, 2304.27, 2352.02}},
		{date: "2018/3/8", data: [4]float32{2326.42, 2318.61, 2314.59, 2333.67}},
		{date: "2018/3/11", data: [4]float32{2314.68, 2310.59, 2296.58, 2320.96}},
		{date: "2018/3/12", data: [4]float32{2309.16, 2286.6, 2264.83, 2333.29}},
		{date: "2018/3/13", data: [4]float32{2282.17, 2263.97, 2253.25, 2286.33}},
		{date: "2018/3/14", data: [4]float32{2255.77, 2270.28, 2253.31, 2276.22}},
		{date: "2018/3/15", data: [4]float32{2269.31, 2278.4, 2250, 2312.08}},
		{date: "2018/3/18", data: [4]float32{2267.29, 2240.02, 2239.21, 2276.05}},
		{date: "2018/3/19", data: [4]float32{2244.26, 2257.43, 2232.02, 2261.31}},
		{date: "2018/3/20", data: [4]float32{2257.74, 2317.37, 2257.42, 2317.86}},
		{date: "2018/3/21", data: [4]float32{2318.21, 2324.24, 2311.6, 2330.81}},
		{date: "2018/3/22", data: [4]float32{2321.4, 2328.28, 2314.97, 2332}},
		{date: "2018/3/25", data: [4]float32{2334.74, 2326.72, 2319.91, 2344.89}},
		{date: "2018/3/26", data: [4]float32{2318.58, 2297.67, 2281.12, 2319.99}},
		{date: "2018/3/27", data: [4]float32{2299.38, 2301.26, 2289, 2323.48}},
		{date: "2018/3/28", data: [4]float32{2273.55, 2236.3, 2232.91, 2273.55}},
		{date: "2018/3/29", data: [4]float32{2238.49, 2236.62, 2228.81, 2246.87}},
		{date: "2018/4/1", data: [4]float32{2229.46, 2234.4, 2227.31, 2243.95}},
		{date: "2018/4/2", data: [4]float32{2234.9, 2227.74, 2220.44, 2253.42}},
		{date: "2018/4/3", data: [4]float32{2232.69, 2225.29, 2217.25, 2241.34}},
		{date: "2018/4/8", data: [4]float32{2196.24, 2211.59, 2180.67, 2212.59}},
		{date: "2018/4/9", data: [4]float32{2215.47, 2225.77, 2215.47, 2234.73}},
		{date: "2018/4/10", data: [4]float32{2224.93, 2226.13, 2212.56, 2233.04}},
		{date: "2018/4/11", data: [4]float32{2236.98, 2219.55, 2217.26, 2242.48}},
		{date: "2018/4/12", data: [4]float32{2218.09, 2206.78, 2204.44, 2226.26}},
		{date: "2018/4/15", data: [4]float32{2199.91, 2181.94, 2177.39, 2204.99}},
		{date: "2018/4/16", data: [4]float32{2169.63, 2194.85, 2165.78, 2196.43}},
		{date: "2018/4/17", data: [4]float32{2195.03, 2193.8, 2178.47, 2197.51}},
		{date: "2018/4/18", data: [4]float32{2181.82, 2197.6, 2175.44, 2206.03}},
		{date: "2018/4/19", data: [4]float32{2201.12, 2244.64, 2200.58, 2250.11}},
		{date: "2018/4/22", data: [4]float32{2236.4, 2242.17, 2232.26, 2245.12}},
		{date: "2018/4/23", data: [4]float32{2242.62, 2184.54, 2182.81, 2242.62}},
		{date: "2018/4/24", data: [4]float32{2187.35, 2218.32, 2184.11, 2226.12}},
		{date: "2018/4/25", data: [4]float32{2213.19, 2199.31, 2191.85, 2224.63}},
		{date: "2018/4/26", data: [4]float32{2203.89, 2177.91, 2173.86, 2210.58}},
		{date: "2018/5/2", data: [4]float32{2170.78, 2174.12, 2161.14, 2179.65}},
		{date: "2018/5/3", data: [4]float32{2179.05, 2205.5, 2179.05, 2222.81}},
		{date: "2018/5/6", data: [4]float32{2212.5, 2231.17, 2212.5, 2236.07}},
		{date: "2018/5/7", data: [4]float32{2227.86, 2235.57, 2219.44, 2240.26}},
		{date: "2018/5/8", data: [4]float32{2242.39, 2246.3, 2235.42, 2255.21}},
		{date: "2018/5/9", data: [4]float32{2246.96, 2232.97, 2221.38, 2247.86}},
		{date: "2018/5/10", data: [4]float32{2228.82, 2246.83, 2225.81, 2247.67}},
		{date: "2018/5/13", data: [4]float32{2247.68, 2241.92, 2231.36, 2250.85}},
		{date: "2018/5/14", data: [4]float32{2238.9, 2217.01, 2205.87, 2239.93}},
		{date: "2018/5/15", data: [4]float32{2217.09, 2224.8, 2213.58, 2225.19}},
		{date: "2018/5/16", data: [4]float32{2221.34, 2251.81, 2210.77, 2252.87}},
		{date: "2018/5/17", data: [4]float32{2249.81, 2282.87, 2248.41, 2288.09}},
		{date: "2018/5/20", data: [4]float32{2286.33, 2299.99, 2281.9, 2309.39}},
		{date: "2018/5/21", data: [4]float32{2297.11, 2305.11, 2290.12, 2305.3}},
		{date: "2018/5/22", data: [4]float32{2303.75, 2302.4, 2292.43, 2314.18}},
		{date: "2018/5/23", data: [4]float32{2293.81, 2275.67, 2274.1, 2304.95}},
		{date: "2018/5/24", data: [4]float32{2281.45, 2288.53, 2270.25, 2292.59}},
		{date: "2018/5/27", data: [4]float32{2286.66, 2293.08, 2283.94, 2301.7}},
		{date: "2018/5/28", data: [4]float32{2293.4, 2321.32, 2281.47, 2322.1}},
		{date: "2018/5/29", data: [4]float32{2323.54, 2324.02, 2321.17, 2334.33}},
		{date: "2018/5/30", data: [4]float32{2316.25, 2317.75, 2310.49, 2325.72}},
		{date: "2018/5/31", data: [4]float32{2320.74, 2300.59, 2299.37, 2325.53}},
		{date: "2018/6/3", data: [4]float32{2300.21, 2299.25, 2294.11, 2313.43}},
		{date: "2018/6/4", data: [4]float32{2297.1, 2272.42, 2264.76, 2297.1}},
		{date: "2018/6/5", data: [4]float32{2270.71, 2270.93, 2260.87, 2276.86}},
		{date: "2018/6/6", data: [4]float32{2264.43, 2242.11, 2240.07, 2266.69}},
		{date: "2018/6/7", data: [4]float32{2242.26, 2210.9, 2205.07, 2250.63}},
		{date: "2018/6/13", data: [4]float32{2190.1, 2148.35, 2126.22, 2190.1}},
	}

	indicators = []charts.IndicatorOpts{
		{Name: "AQI", Max: 300},
		{Name: "PM2.5", Max: 250},
		{Name: "PM10", Max: 300},
		{Name: "CO", Max: 5},
		{Name: "NO2", Max: 200},
		{Name: "SO2", Max: 100},
	}

	radarDataBJ = [][]float32{
		{55, 9, 56, 0.46, 18, 6, 1},
		{25, 11, 21, 0.65, 34, 9, 2},
		{56, 7, 63, 0.3, 14, 5, 3},
		{33, 7, 29, 0.33, 16, 6, 4},
		{42, 24, 44, 0.76, 40, 16, 5},
		{82, 58, 90, 1.77, 68, 33, 6},
		{74, 49, 77, 1.46, 48, 27, 7},
		{78, 55, 80, 1.29, 59, 29, 8},
		{267, 216, 280, 4.8, 108, 64, 9},
		{185, 127, 216, 2.52, 61, 27, 10},
		{39, 19, 38, 0.57, 31, 15, 11},
		{41, 11, 40, 0.43, 21, 7, 12},
		{64, 38, 74, 1.04, 46, 22, 13},
		{108, 79, 120, 1.7, 75, 41, 14},
		{108, 63, 116, 1.48, 44, 26, 15},
		{33, 6, 29, 0.34, 13, 5, 16},
		{94, 66, 110, 1.54, 62, 31, 17},
		{186, 142, 192, 3.88, 93, 79, 18},
		{57, 31, 54, 0.96, 32, 14, 19},
		{22, 8, 17, 0.48, 23, 10, 20},
		{39, 15, 36, 0.61, 29, 13, 21},
	}
)

type klineData struct {
	date string
	data [4]float32
}

func genKvData() map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(nameItems); i++ {
		m[nameItems[i]] = rand.Intn(maxNum)
	}
	return m
}

func randInt() []int {
	r := make([]int, 0)
	for i := 0; i < 6; i++ {
		r = append(r, int(seed.Int63())%50)
	}
	return r
}

func genBar3dData() [][3]int {

	data := [][3]int{
		{0, 0, 5}, {0, 1, 1}, {0, 2, 0}, {0, 3, 0}, {0, 4, 0}, {0, 5, 0},
		{0, 6, 0}, {0, 7, 0}, {0, 8, 0}, {0, 9, 0}, {0, 10, 0}, {0, 11, 2},
		{0, 12, 4}, {0, 13, 1}, {0, 14, 1}, {0, 15, 3}, {0, 16, 4}, {0, 17, 6},
		{0, 18, 4}, {0, 19, 4}, {0, 20, 3}, {0, 21, 3}, {0, 22, 2}, {0, 23, 5},
		{1, 0, 7}, {1, 1, 0}, {1, 2, 0}, {1, 3, 0}, {1, 4, 0}, {1, 5, 0},
		{1, 6, 0}, {1, 7, 0}, {1, 8, 0}, {1, 9, 0}, {1, 10, 5}, {1, 11, 2},
		{1, 12, 2}, {1, 13, 6}, {1, 14, 9}, {1, 15, 11}, {1, 16, 6}, {1, 17, 7},
		{1, 18, 8}, {1, 19, 12}, {1, 20, 5}, {1, 21, 5}, {1, 22, 7}, {1, 23, 2},
		{2, 0, 1}, {2, 1, 1}, {2, 2, 0}, {2, 3, 0}, {2, 4, 0}, {2, 5, 0},
		{2, 6, 0}, {2, 7, 0}, {2, 8, 0}, {2, 9, 0}, {2, 10, 3}, {2, 11, 2},
		{2, 12, 1}, {2, 13, 9}, {2, 14, 8}, {2, 15, 10}, {2, 16, 6}, {2, 17, 5},
		{2, 18, 5}, {2, 19, 5}, {2, 20, 7}, {2, 21, 4}, {2, 22, 2}, {2, 23, 4},
		{3, 0, 7}, {3, 1, 3}, {3, 2, 0}, {3, 3, 0}, {3, 4, 0}, {3, 5, 0},
		{3, 6, 0}, {3, 7, 0}, {3, 8, 1}, {3, 9, 0}, {3, 10, 5}, {3, 11, 4},
		{3, 12, 7}, {3, 13, 14}, {3, 14, 13}, {3, 15, 12}, {3, 16, 9}, {3, 17, 5},
		{3, 18, 5}, {3, 19, 10}, {3, 20, 6}, {3, 21, 4}, {3, 22, 4}, {3, 23, 1},
		{4, 0, 1}, {4, 1, 3}, {4, 2, 0}, {4, 3, 0}, {4, 4, 0}, {4, 5, 1},
		{4, 6, 0}, {4, 7, 0}, {4, 8, 0}, {4, 9, 2}, {4, 10, 4}, {4, 11, 4},
		{4, 12, 2}, {4, 13, 4}, {4, 14, 4}, {4, 15, 14}, {4, 16, 12}, {4, 17, 1},
		{4, 18, 8}, {4, 19, 5}, {4, 20, 3}, {4, 21, 7}, {4, 22, 3}, {4, 23, 0},
		{5, 0, 2}, {5, 1, 1}, {5, 2, 0}, {5, 3, 3}, {5, 4, 0}, {5, 5, 0},
		{5, 6, 0}, {5, 7, 0}, {5, 8, 2}, {5, 9, 0}, {5, 10, 4}, {5, 11, 1},
		{5, 12, 5}, {5, 13, 10}, {5, 14, 5}, {5, 15, 7}, {5, 16, 11}, {5, 17, 6},
		{5, 18, 0}, {5, 19, 5}, {5, 20, 3}, {5, 21, 4}, {5, 22, 2}, {5, 23, 0},
		{6, 0, 1}, {6, 1, 0}, {6, 2, 0}, {6, 3, 0}, {6, 4, 0}, {6, 5, 0},
		{6, 6, 0}, {6, 7, 0}, {6, 8, 0}, {6, 9, 0}, {6, 10, 1}, {6, 11, 0},
		{6, 12, 2}, {6, 13, 1}, {6, 14, 3}, {6, 15, 4}, {6, 16, 0}, {6, 17, 0},
		{6, 18, 0}, {6, 19, 0}, {6, 20, 1}, {6, 21, 2}, {6, 22, 2}, {6, 23, 6},
	}

	for i := 0; i < len(data); i++ {
		data[i][0], data[i][1] = data[i][1], data[i][0]
	}
	return data
}

func genLinks() []charts.GraphLink {
	links := make([]charts.GraphLink, 0)
	for i := 0; i < len(graphNodes); i++ {
		for j := 0; j < len(graphNodes); j++ {
			links = append(links,
				charts.GraphLink{Source: graphNodes[i].Name, Target: graphNodes[j].Name})
		}
	}
	return links
}

func genHeatMapData() [][3]interface{} {
	res := make([][3]interface{}, 0)
	for i := 0; i < len(hmData); i++ {
		if int(hmData[i][2]) == 0 {
			res = append(res, [3]interface{}{hmData[i][1], hmData[i][0], "-"})
		} else {
			res = append(res, [3]interface{}{hmData[i][1], hmData[i][0], hmData[i][2]})
		}
	}
	return res
}

func genLine3dData() [][3]float64 {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{
				(1 + 0.25*math.Cos(75*float64(t))) * math.Cos(float64(t)),
				(1 + 0.25*math.Cos(75*float64(t))) * math.Sin(float64(t)),
				float64(t) + 2.0*math.Sin(75.0*float64(t)),
			},
		)
	}
	return data
}
