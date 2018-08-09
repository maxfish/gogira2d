package graphics

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Camera2D a Camera based on an orthogonal projection
type Camera2D struct {
	x                float32
	y                float32
	width            float32
	height           float32
	zoom             float32
	centered         bool
	near             float32
	far              float32
	projectionMatrix mgl32.Mat4
	matrixDirty      bool
}

// NewCamera2D sets up an orthogonal projection camera
func NewCamera2D(width int, height int, zoom float32, centered bool) *Camera2D {
	c := &Camera2D{
		width:    float32(width),
		height:   float32(height),
		zoom:     zoom,
		centered: centered,
	}
	c.far = -2
	c.near = 2
	c.rebuildMatrix()

	return c
}

// ProjectionMatrix returns the projection matrix of the camera
func (c *Camera2D) ProjectionMatrix() mgl32.Mat4 {
	if c.matrixDirty {
		c.rebuildMatrix()
	}
	return c.projectionMatrix
}

// SetPosition sets the current position of the camera. If the camera is centered, the center will be moving
func (c *Camera2D) SetPosition(x float32, y float32) {
	c.x = x
	c.y = y
	c.matrixDirty = true
}

// SetZoom sets the zoom factor
func (c *Camera2D) SetZoom(zoom float32) {
	c.zoom = zoom
	c.matrixDirty = true
}

func (c *Camera2D) rebuildMatrix() {
	var left, right, top, bottom float32

	if c.centered {
		halfWidth := c.width / 2 / c.zoom
		halfHeight := c.height / 2 / c.zoom
		left = -halfWidth
		right = halfWidth
		top = halfHeight
		bottom = -halfHeight
	} else {
		right = c.width / c.zoom
		top = c.height / c.zoom
	}

	left += c.x
	right += c.x
	top += c.y
	bottom += c.y

	c.projectionMatrix = mgl32.Ortho(left, right, top, bottom, c.near, c.far)
	c.matrixDirty = false
}