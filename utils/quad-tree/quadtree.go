package quadtree

type quadtreeInterface interface {
	RangeQuery(r1,c1,r2,c2 int) int
	PointUpdate(r,c,val int)
	BuildQuadTree() []int
}


type quadTree struct {
	Grid [][]int
	row  int
	col  int
	Tree []int
}

func InitQuadTree(grid [][]int, row, col int) *quadTree{
	tree := quadTree{Grid: grid,row:row-1,col:col-1}
	qTree := tree.BuildQuadTree()
	tree.Tree = qTree
	return &tree
}

func (tree quadTree) BuildQuadTree() []int{
	arr := make([]int,4*(tree.row+1)*(tree.col+1))
	buildQuadTreeUtil(tree.Grid,&arr,0,0,0,tree.row,tree.col)
	return arr
}

func buildQuadTreeUtil(grid [][]int ,tree *[]int , index, r1 , c1 , r2 , c2 int) int {

	if isSingleCell(r1,c1,r2,c2 )  {
		//base case
		(*tree)[index] = grid[r1][c1]
		return (*tree)[index]
	}
	if c1 == c2 {
		midRow := (r1+r2)/2
		val1 := buildQuadTreeUtil(grid,tree,4*index+1,r1,c1,midRow,c1)
		val2 := buildQuadTreeUtil(grid,tree,4*index+2,r1,c1,midRow+1,c1)
		(*tree)[index] = val1 + val2
		return (*tree)[index]
	}
	if r1 == r2 {
		midCol := (c1+c2)/2
		val3 := buildQuadTreeUtil(grid,tree,4*index+1,r1,c1,r1,midCol)
		val4 := buildQuadTreeUtil(grid,tree,4*index+2,r1,midCol+1,r1,c2)
		(*tree)[index] = val3 + val4
		return (*tree)[index]
	}
	midRow := (r1+r2)/2
	midCol := (c1+c2)/2

	val1 := buildQuadTreeUtil(grid,tree,4*index+1,r1,c1,midRow,midCol)
	val2 := buildQuadTreeUtil(grid,tree,4*index+2,r1,midCol+1,midRow,c2)
	val3 := buildQuadTreeUtil(grid,tree,4*index+3,midRow+1,c1,r2,midCol)
	val4 := buildQuadTreeUtil(grid,tree,4*index+4,midRow+1,midCol+1,r2,c2)

	(*tree)[index] = val1 + val2 + val3 + val4

	return (*tree)[index]
}

func isSingleCell(r1 int, c1 int, r2 int, c2 int) bool{
	return r1==r2 && c1 ==c2
}

func (tree quadTree) PointUpdate(r,c,val int){
	diff := val - tree.Grid[r][c]
	tree.Grid[r][c] = val
	pointUpdateUtil(&tree.Tree,0,0,0,tree.row,tree.col,r,c,diff)
}

func pointUpdateUtil(tree *[]int,index int, x1 int, y1 int, x2 int, y2 int, r int, c int, diff int) {
	if isPointOutsideRectangle(r,c , x1,y1,x2,y2){
		return
	}
	if isSingleCell(x1,y1,x2,y2){
		(*tree)[index] += diff
		return
	}
	midRow := (x1+x2)/2
	midCol := (y1+y2)/2
	pointUpdateUtil(tree, index*4+1, x1,y1,midRow,midCol, r,c, diff)
	pointUpdateUtil(tree, index*4+2, x1,midCol+1,midRow,y2, r,c, diff)
	pointUpdateUtil(tree, index*4+3, midRow+1,y1,x2,midCol, r,c, diff)
	pointUpdateUtil(tree, index*4+4, midRow+1,midCol+1,x2,y2, r,c, diff)
	(*tree)[index]  += diff

}

func (tree quadTree) RangeQuery(r1,c1,r2,c2 int) int{
	return rangeQueryUtil(tree.Tree,0,0,0,tree.row,tree.col,r1,c1,r2,c2)

}

func rangeQueryUtil(tree []int,index int, x1 int, y1 int, x2 int, y2 int, r1 int, c1 int, r2 int, c2 int) int {
	if isRangeOutside(x1,y1,x2,y2,r1,c1,r2,c2) && !isRangeContained(r1,c1,r2,c2,x1,y1,x2,y2){
		return 0
	}
	if isRangeContained(x1,y1,x2,y2,r1,c1,r2,c2) {
		return tree[index]
	}
	midRow := (x1+x2)/2
	midCol := (y1+y2)/2
	if x1 == x2 {
		val1 := rangeQueryUtil(tree, 4*index+1,x1,y1,midRow,midCol,r1,c1,r2,c2)
		val2 := rangeQueryUtil(tree, 4*index+2,x1,midCol+1,midRow,y2,r1,c1,r2,c2)
		return val1+val2
	}
	if y1==y2 {
		val1 := rangeQueryUtil(tree, 4*index+1,midRow+1,y1,x2,midCol,r1,c1,r2,c2)
		val2 := rangeQueryUtil(tree, 4*index+2,midRow+1,midCol+1,x2,y2,r1,c1,r2,c2)
		return val1+val2
	}

	val1 := rangeQueryUtil(tree,4*index+1,x1,y1,midRow,midCol,r1,c1,r2,c2)
	val2 := rangeQueryUtil(tree,4*index+2,x1,midCol+1,midRow,y2,r1,c1,r2,c2)
	val3 := rangeQueryUtil(tree,4*index+3,midRow+1,y1,x2,midCol,r1,c1,r2,c2)
	val4 := rangeQueryUtil(tree,4*index+4,midRow+1,midCol+1,x2,y2,r1,c1,r2,c2)

	return val1 + val2 + val3 + val4
}

func isRangeContained(x1 int , y1 int, x2 int, y2 int, r1 int, c1 int, r2 int, c2 int) bool {
	return liesWithin(x1, r1, r2) && liesWithin(x2,r1,r2) && liesWithin(y1,c1,c2) && liesWithin(y2,c1,c2)
}

func isRangeOutside(x1 int , y1 int, x2 int, y2 int, r1 int, c1 int, r2 int, c2 int) bool {
	res1 := isPointOutsideRectangle(x1, y1, r1, c1, r2, c2)
	res2 := isPointOutsideRectangle(x2, y2, r1, c1, r2, c2)
	return res1 && res2
}
func isPointOutsideRectangle(x,y, x1,y1,x2,y2 int) bool {
	within1 := liesWithin(x, x1, x2)
	within2 := liesWithin(y, y1, y2)
	return !(within1 && within2)
}

// Returns true when a lies in [b,c] else false
func liesWithin(a int, b int, c int) bool {
	return a >= b && a <= c
}






