package goweb

import (
	"fmt"
	"net/http"
	"strings"
)

// *只能出现于最后

type Router struct {
	handlers map[string]HandlerFunc // 路由的处理函数
	roots    map[string]*node // 路由的trie节点
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc),
		roots: make(map[string]*node),
	}
}

func parsePatten(pattern string) []string {
	v := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range v {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePatten(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	key := method + "-" + pattern
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	fmt.Printf("[Context]: req.method:%s, req.path:%s\n", c.Method, c.Path)
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Req.Method + "-" + c.Req.URL.Path
		c.handlers = append(c.handlers, r.handlers[key])
		//r.handlers[key](c)
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}

func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	//searchParts := strings.Split(path, "/")
	searchParts := parsePatten(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePatten(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
