package robot

import "github.com/nineora/nineora/nine/nineora"

var gSuperiorMap = make(map[nineora.Link]nineora.Link)

func superiorPut(cur nineora.Link, superior nineora.Link) {
	gSuperiorMap[cur] = superior
}

func GetSuperior(cur nineora.Link) nineora.Link {
	s, ok := gSuperiorMap[cur]
	if !ok {
		return ""
	}
	return s
}
