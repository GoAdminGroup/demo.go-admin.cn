package pages

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/sword/components/card"
	"html/template"
)

func GetDashBoard2Content() (types.Panel, error) {

	components := template2.Get(config.Get().Theme)
	colComp := components.Col()

	/**************************
	 * Info Box
	/**************************/

	cardcard := card.New().
		SetTitle("总销售额").
		SetSubTitle("¥ 113,340").
		SetAction(template.HTML(`<i aria-label="图标: info-circle-o" class="anticon anticon-info-circle-o"><svg viewBox="64 64 896 896" focusable="false" class="" data-icon="info-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M464 336a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm72 112h-48c-4.4 0-8 3.6-8 8v272c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V456c0-4.4-3.6-8-8-8z"></path></svg></i>`)).
		SetContent(template.HTML(`<div><div title="" style="margin-right: 16px;"><span><span>周同比</span><span style="margin-left: 8px;">12%</span></span><span style="color: #f5222d;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-up" class="anticon anticon-caret-up"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-up" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M858.9 689L530.5 308.2c-9.4-10.9-27.5-10.9-37 0L165.1 689c-12.2 14.2-1.2 35 18.5 35h656.8c19.7 0 30.7-20.8 18.5-35z"></path></svg></i></span></div><div class="antd-pro-pages-dashboard-analysis-components-trend-index-trendItem" title=""><span><span>日同比</span><span style="margin-left: 8px;">11%</span></span><span style="color: #52c41a;margin-left: 4px;top: 1px;"><i style="font-size: 12px;" aria-label="图标: caret-down" class="anticon anticon-caret-down"><svg viewBox="0 0 1024 1024" focusable="false" class="" data-icon="caret-down" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M840.4 300H183.6c-19.7 0-30.7 20.8-18.5 35l328.4 380.8c9.4 10.9 27.5 10.9 37 0L858.9 335c12.2-14.2 1.2-35-18.5-35z"></path></svg></i></span></div></div>`)).
		SetFooter(template.HTML(`日销售额 <strong style="margin-left:8px;">￥11,325</strong>`))
	infobox := cardcard.GetContent()

	infobox2 := cardcard.GetContent()

	infobox3 := cardcard.GetContent()

	infobox4 := cardcard.GetContent()

	var size = map[string]string{"md": "3", "sm": "6", "xs": "12"}
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()
	row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

	/**************************
	 * Box
	/**************************/

	chartdata := `{
  "datasets": [
    {
      "data": [65,59,80,81,56,55,40],
      "fillColor": "#ececec",
      "label": "Electronics",
      "pointColor": "#ececec",
      "pointHighlightFill": "#ececec",
      "pointHighlightStroke": "#ececec",
      "pointStrokeColor": "#ececec",
      "strokeColor": "#ececec"
    },
    {
      "data": [28,48,40,19,86,27,90],
      "fillColor": "rgba(102, 176, 255, 0.9)",
      "label": "Digital Goods",
      "pointColor": "rgb(102, 176, 255)",
      "pointHighlightFill": "rgb(102, 176, 255)",
      "pointHighlightStroke": "rgb(102, 176, 255)",
      "pointStrokeColor": "rgb(102, 176, 255)",
      "strokeColor": "rgb(102, 176, 255)"
    }
  ],
  "labels": ["January","February","March","April","May","June","July"]
}`

	lineChart := components.AreaChart().SetID("salechart").
		SetData(chartdata).
		SetHeight(180).
		SetTitle("Sales: 1 Jan, 2019 - 30 Jul, 2019").GetContent()

	title := `<p class="text-center"><strong>Goal Completion</strong></p>`
	progressGroup := components.ProgressGroup().
		SetTitle("Add Products to Cart").
		SetColor("#76b2d4").
		SetDenominator(200).
		SetMolecular(160).
		SetPercent(80).
		GetContent()

	progressGroup1 := components.ProgressGroup().
		SetTitle("Complete Purchase").
		SetColor("#f17c6e").
		SetDenominator(400).
		SetMolecular(310).
		SetPercent(80).
		GetContent()

	progressGroup2 := components.ProgressGroup().
		SetTitle("Visit Premium Page").
		SetColor("#ace0ae").
		SetDenominator(800).
		SetMolecular(490).
		SetPercent(80).
		GetContent()

	progressGroup3 := components.ProgressGroup().
		SetTitle("Send Inquiries").
		SetColor("#fdd698").
		SetDenominator(500).
		SetMolecular(250).
		SetPercent(50).
		GetContent()

	boxInternalCol1 := colComp.SetContent(lineChart).SetSize(map[string]string{"md": "8"}).GetContent()
	boxInternalCol2 := colComp.
		SetContent(template.HTML(title) + progressGroup + progressGroup1 + progressGroup2 + progressGroup3).
		SetSize(map[string]string{"md": "4"}).
		GetContent()

	boxInternalRow := components.Row().SetContent(boxInternalCol1 + boxInternalCol2).GetContent()

	description1 := components.Description().SetPercent("17").
		SetNumber("¥140,100").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description2 := components.Description().
		SetPercent("2").
		SetNumber("440,560").
		SetTitle("TOTAL REVENUE").
		SetArrow("down").
		SetColor("red").
		SetBorder("right").
		GetContent()

	description3 := components.Description().
		SetPercent("12").
		SetNumber("¥140,050").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		SetBorder("right").
		GetContent()

	description4 := components.Description().
		SetPercent("1").
		SetNumber("30943").
		SetTitle("TOTAL REVENUE").
		SetArrow("up").
		SetColor("green").
		GetContent()

	size2 := map[string]string{"sm": "3", "xs": "6"}
	boxInternalCol3 := colComp.SetContent(description1).SetSize(size2).GetContent()
	boxInternalCol4 := colComp.SetContent(description2).SetSize(size2).GetContent()
	boxInternalCol5 := colComp.SetContent(description3).SetSize(size2).GetContent()
	boxInternalCol6 := colComp.SetContent(description4).SetSize(size2).GetContent()

	boxInternalRow2 := components.Row().SetContent(boxInternalCol3 + boxInternalCol4 + boxInternalCol5 + boxInternalCol6).GetContent()

	box := components.Box().WithHeadBorder(true).SetHeader("Monthly Recap Report").
		SetBody(boxInternalRow).
		SetFooter(boxInternalRow2).
		GetContent()

	boxcol := colComp.SetContent(box).SetSize(map[string]string{"md": "12"}).GetContent()
	row2 := components.Row().SetContent(boxcol).GetContent()

	/**************************
	 * Pie Chart
	/**************************/

	pieData := `[
  {
    "color": "#f56954",
    "highlight": "#f56954",
    "label": "Chrome",
    "value": 700
  },
  {
    "color": "#00a65a",
    "highlight": "#00a65a",
    "label": "IE",
    "value": 500
  },
  {
    "color": "#f39c12",
    "highlight": "#f39c12",
    "label": "FireFox",
    "value": 400
  },
  {
    "color": "#00c0ef",
    "highlight": "#00c0ef",
    "label": "Safari",
    "value": 600
  },
  {
    "color": "#3c8dbc",
    "highlight": "#3c8dbc",
    "label": "Opera",
    "value": 300
  },
  {
    "color": "#d2d6de",
    "highlight": "#d2d6de",
    "label": "Navigator",
    "value": 100
  }
]`
	pie := components.PieChart().SetHeight(170).SetData(pieData).SetID("pieChart").GetContent()
	legend := components.ChartLegend().SetData([]map[string]string{
		{
			"label": " Chrome",
			"color": "red",
		}, {
			"label": " IE",
			"color": "Green",
		}, {
			"label": " FireFox",
			"color": "yellow",
		}, {
			"label": " Sarafri",
			"color": "blue",
		}, {
			"label": " Opera",
			"color": "light-blue",
		}, {
			"label": " Navigator",
			"color": "gray",
		},
	}).GetContent()

	boxDanger := components.Box().SetTheme("danger").WithHeadBorder(true).SetHeader("Browser Usage").
		SetBody(components.Row().
			SetContent(colComp.SetSize(map[string]string{"md": "8"}).
				SetContent(pie).
				GetContent() + colComp.SetSize(map[string]string{"md": "4"}).
				SetContent(legend).
				GetContent()).GetContent()).
		SetFooter(`<p class="text-center"><a href="javascript:void(0)" class="uppercase">View All Users</a></p>`).
		GetContent()

	tabs := components.Tabs().SetData([]map[string]template.HTML{
		{
			"title": "tabs1",
			"content": template.HTML(`<b>How to use:</b>

                <p>Exactly like the original bootstrap tabs except you should use
                  the custom wrapper <code>.nav-tabs-custom</code> to achieve this style.</p>
                A wonderful serenity has taken possession of my entire soul,
                like these sweet mornings of spring which I enjoy with my whole heart.
                I am alone, and feel the charm of existence in this spot,
                which was created for the bliss of souls like mine. I am so happy,
                my dear friend, so absorbed in the exquisite sense of mere tranquil existence,
                that I neglect my talents. I should be incapable of drawing a single stroke
                at the present moment; and yet I feel that I never was a greater artist than now.`),
		}, {
			"title": "tabs2",
			"content": template.HTML(`
                The European languages are members of the same family. Their separate existence is a myth.
                For science, music, sport, etc, Europe uses the same vocabulary. The languages only differ
                in their grammar, their pronunciation and their most common words. Everyone realizes why a
                new common language would be desirable: one could refuse to pay expensive translators. To
                achieve this, it would be necessary to have uniform grammar, pronunciation and more common
                words. If several languages coalesce, the grammar of the resulting language is more simple
                and regular than that of the individual languages.
              `),
		}, {
			"title": "tabs3",
			"content": template.HTML(`
                Lorem Ipsum is simply dummy text of the printing and typesetting industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
                when an unknown printer took a galley of type and scrambled it to make a type specimen book.
                It has survived not only five centuries, but also the leap into electronic typesetting,
                remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
                sheets containing Lorem Ipsum passages, and more recently with desktop publishing software
                like Aldus PageMaker including versions of Lorem Ipsum.
              `),
		},
	}).GetContent()

	buttonTest := `<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#exampleModal" data-whatever="@mdo">Open modal for @mdo</button>`
	popupForm := `<form>
          <div class="form-group">
            <label for="recipient-name" class="col-form-label">Recipient:</label>
            <input type="text" class="form-control" id="recipient-name">
          </div>
          <div class="form-group">
            <label for="message-text" class="col-form-label">Message:</label>
            <textarea class="form-control" id="message-text"></textarea>
          </div>
        </form>`
	popup := components.Popup().SetID("exampleModal").
		SetFooter("Save Change").
		SetTitle("this is a popup").
		SetBody(template.HTML(popupForm)).
		GetContent()

	col5 := colComp.SetSize(map[string]string{"md": "8"}).SetContent(tabs + template.HTML(buttonTest)).GetContent()
	col6 := colComp.SetSize(map[string]string{"md": "4"}).SetContent(boxDanger + popup).GetContent()

	row4 := components.Row().SetContent(col5 + col6).GetContent()

	return types.Panel{
		Content:     row1 + row2 + row4,
		Title:       "Dashboard",
		Description: "dashboard example",
	}, nil
}
