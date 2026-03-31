/*
   Creme de la creme of system design, apparantly.
   Where to use? simple, where there is object initiation just replace it with factory call.
   But the question is, would you create a Factory to create a Factory? Its a joke, dont take it literally, unless.
   Use it when you call fresh uncontaminated object on multiple spaces so wrapping it would save you future refactoring
*/

package main

import "fmt"

func main() {
	m1 := CreateMesh("CUBE")
	m1.printVolume()
	m2 := CreateMesh("PYRAMID")
	m2.printVolume()
	m3 := CreateMesh("POINT")
	m3.printVolume()
	m4 := CreateMesh("")
	m4.printVolume()
}

func CreateMesh(meshType string) Volume {
	if meshType == "" {
		return &Mesh{}
	}
	switch meshType {
	case "CUBE":
		c := NewCube()
		return &c
	case "PYRAMID":
		p := NewPyramid()
		return &p
	case "POINT":
		p := NewPoint()
		return &p
	default:
		return &Mesh{}
	}
}

type Volume interface {
    printVolume()
}

type Mesh struct {
	vertices []Vertex
	edges    []Edge
	faces    []Face
}

func (m *Mesh) addVertex(x float32, y float32, z float32) {
	v := Vertex{
		x: x,
		y: y,
		z: z,
	}
	m.vertices = append(m.vertices, v)
}

func (m *Mesh) addEdge(vi1 int, vi2 int) {
	e := Edge{
		v1: m.vertices[vi1],
		v2: m.vertices[vi2],
	}
	m.edges = append(m.edges, e)
}

func (m *Mesh) addFace(vertexIndices ...int) {
	faceVertices := []Vertex{}
	for _, vi := range vertexIndices {
		faceVertices = append(faceVertices, m.vertices[vi])
	}
	m.faces = append(m.faces, Face{
		faceVertices: faceVertices,
	})
}

func (m *Mesh) printVolume() {
	fmt.Println("Volume of custom mesh is ∞")
}

type Vertex struct {
	x float32
	y float32
	z float32
}

type Edge struct {
	v1 Vertex
	v2 Vertex
}

type Face struct {
	faceVertices []Vertex
}

type Cube struct {
	Mesh
}

func NewCube() Cube {
	c := Cube{}
	c.addVertex(0,0,0) 
	c.addVertex(1,0,0) 
	c.addVertex(1,1,0) 
	c.addVertex(0,1,0)
	c.addVertex(0,0,1) 
	c.addVertex(1,0,1) 
	c.addVertex(1,1,1) 
	c.addVertex(0,1,1)
	c.addEdge(0, 1) 
	c.addEdge(1, 2) 
	c.addEdge(2, 3) 
	c.addEdge(3, 0)
	c.addEdge(4, 5) 
	c.addEdge(5, 6) 
	c.addEdge(6, 7) 
	c.addEdge(7, 4)
	c.addEdge(0, 4) 
	c.addEdge(1, 5) 
	c.addEdge(2, 6) 
	c.addEdge(3, 7)
	c.addFace(0, 1, 2, 3)
	c.addFace(4, 5, 6, 7)
	c.addFace(0, 1, 5, 4)
	c.addFace(1, 2, 6, 5)
	c.addFace(2, 3, 7, 6)
	c.addFace(3, 0, 4, 7)
	return c;
}

func (c *Cube) printVolume() {
	fmt.Println("Volume of cube is 1")
}

type Pyramid struct {
	Mesh
}

func NewPyramid() Pyramid {
	p := Pyramid{}
	p.addVertex(0, 0, 0)
	p.addVertex(1, 0, 0)
	p.addVertex(1, 1, 0)
	p.addVertex(0, 1, 0)
	p.addVertex(0.5, 0.5, 1)
	p.addEdge(0, 1) 
	p.addEdge(1, 2) 
	p.addEdge(2, 3) 
	p.addEdge(3, 0)
	p.addEdge(0, 4) 
	p.addEdge(1, 4) 
	p.addEdge(2, 4) 
	p.addEdge(3, 4)
	p.addFace(0, 1, 2, 3) 
	p.addFace(0, 1, 4) 
	p.addFace(1, 2, 4)
	p.addFace(2, 3, 4)
	p.addFace(3, 0, 4)
	return p;
}

func (p *Pyramid) printVolume() {
	fmt.Println("Volume of pyramid is 0.33")
}

type Point struct {
	Mesh
}

func NewPoint() Point {
	p := Point{}
	p.addVertex(0, 0, 0)
	return p
}

func (p *Point) printVolume() {
	fmt.Println("Volume of point is 0")
}