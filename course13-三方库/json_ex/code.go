package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func main() {

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
