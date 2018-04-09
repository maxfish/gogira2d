package main

import (
	g "gojira2d/pkg/graphics"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	a "gojira2d/pkg/app"
)

func main() {
	app := a.InitApp(800, 600, false, "Quads")
	defer a.TerminateApp()

	quads := make([]*g.Primitive2D, 0, 12)
	texture := g.NewTextureFromFile("examples/assets/texture.png")

	// Creates 12 quads in a grid 4x3
	for y := 0; y < 3; y++ {
		for x := 0; x < 4; x++ {
			quad := g.NewQuadPrimitive(mgl32.Vec3{float32(100 + x*200), float32(100 + y*200), 0}, mgl32.Vec2{120, 120})
			quad.SetTexture(texture)
			quad.SetAnchorToCenter()
			quads = append(quads, quad)
		}
	}

	// Flipped horizontally
	quads[1].SetFlipX(true)
	// Flipped vertically
	quads[2].SetFlipY(true)
	// Flipped in both directions
	quads[3].SetFlipX(true)
	quads[3].SetFlipY(true)
	// Scaled smaller
	quads[4].SetScale(mgl32.Vec2{0.8, 0.8})
	// Scaled bigger
	quads[5].SetScale(mgl32.Vec2{1.2, 1.2})
	// Rotated 45 degrees
	quads[6].SetAngle(math.Pi / 4)
	// Rotated 22 degrees and scaled to 50%
	quads[7].SetAngle(math.Pi / 8)
	quads[7].SetScale(mgl32.Vec2{0.8, 0.8})
	// Stretched
	quads[8].SetScale(mgl32.Vec2{0.5, 1.2})

	var animationAngle float32 = 0

	app.MainLoop(func(speed float64) {
		animationAngle += float32(speed)
		animationScale := float32(math.Abs(math.Sin(float64(animationAngle * 2))))
		quads[9].SetScale(mgl32.Vec2{animationScale, animationScale})
		quads[10].SetAngle(animationAngle)
		quads[11].SetScale(mgl32.Vec2{animationScale, animationScale})
		quads[11].SetAngle(animationAngle)
	}, func() {
		for i:=0; i<2000; i++ {
			for _, q := range quads {
				q.EnqueueForDrawing(app.Context)
			}
		}
		println("FPS: ", app.FpsCounter.FPS())
	})
}
