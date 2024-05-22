package gee

import "strings"

// 基于前缀树实现动态路由
type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	//注意child的数据类型，是node*类型的指针
	for _, child := range n.children {
		//精准匹配成功或者模糊匹配
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			//注意child的数据类型
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 路由的注册功能，递归实现
func (n *node) insert(pattern string, parts []string, height int) {
	//当递归的深度已经等于url的节点数量时，停止递归
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	//如果没有匹配到相同的节点
	if child == nil {
		//新建一个节点
		//isWild参数用力控制是否精准匹配
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 路由的匹配功能，递归实现
func (n *node) search(parts []string, height int) *node {

	//当匹配到结束了或者是模糊匹配
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		//匹配失败
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
