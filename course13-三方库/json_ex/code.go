package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func valid() {
	html, _ := ioutil.ReadFile("test.html")
	//println(string(html))
	// body > script:nth-child(3)
	dom, _ := goquery.NewDocumentFromReader(bytes.NewReader(html))
	//res := dom.Find(`script[type="application/ld+json"]`).First()
	//println(res.Text())
	dom.Find(`script[type="application/ld+json"]`).Each(func(i int, selection *goquery.Selection) {
		//println(strings.ReplaceAll(selection.Text(), " ", ""))
		//println("_----------")
		if gjson.Valid(selection.Text()) {
			println(selection.Text())

		}
		//selection.Find

	})

}

func json() {
	html, _ := ioutil.ReadFile("brumate.html")
	println(string(html))
	//mat := regexp.MustCompile(`"actionModule":(.*?);`)
	//res := mat.FindSubmatch(html)
	//println(len(res))
	dom, _ := goquery.NewDocumentFromReader(bytes.NewReader(html))
	//println(dom.Html())
	//res := dom.Find(`meta[property="og:title"]`)
	res := dom.Find(`script`)
	//res := dom.Find(`script[id="__NEXT_DATA__"]`)
	//println(res.First().Attr("content"))
	//println(res.Text())
	res.Each(func(i int, selection *goquery.Selection) {
		println(selection.Text())
	})
}

func main() {
	json()

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(bytes.NewReader(html))
	//jsonData, err := jsonld.ReadJSON(buf.Bytes())
	//options := &jsonld.Options{}
	//options.Base = ""
	//options.ProduceGeneralizedRdf = false
	//dataSet, err := jsonld.ToRDF(jsonData, options)
	//if err != nil {
	//
	//}
	////for t := range dataSet.IterTriples() {
	////	g.AddTriple(jterm2term(t.Subject), jterm2term(t.Predicate), jterm2term(t.Object))
	////}
	//
	//b, _ := json.Marshal(dataSet)
	//println(string(b))
}
