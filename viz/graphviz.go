package viz

import (
	"bytes"
	"log"

	"github.com/goccy/go-graphviz"
)

func Dot2SVG(dotString string) string {
	g := graphviz.New()
	graphFromString, err := graphviz.ParseBytes([]byte(dotString))
	defer func() {
		if err := graphFromString.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	if err != nil {
		log.Fatal(err)
	}

	//graph, err := g.Graph()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//defer func() {
	//	if err := graph.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//	g.Close()
	//}()

	var buf bytes.Buffer
	if err := g.Render(graphFromString, "svg", &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
