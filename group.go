package chef 

type Group struct {
	prefix string 
	handlers []Handler 
}

func newGroup(prefix string) Group {
	return Group{}
}

func (g *Group) Use(middleware ...Handler) {
	g.handlers = append(g.handlers, middleware...)
} 