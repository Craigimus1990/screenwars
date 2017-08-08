package systems

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type PlayerSystem struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	world *ecs.World
	mouseTracker MouseTracker
}

func (ps *PlayerSystem) Remove(ecs.BasicEntity) {}

func (ps *PlayerSystem) Update(dt float32) {
	point := engo.Point{ps.mouseTracker.MouseX, ps.mouseTracker.MouseY}
	ps.SpaceComponent.Position = point
}

func (ps *PlayerSystem) New(w *ecs.World) {
	ps.world = w

	ps.mouseTracker.BasicEntity = ecs.NewBasic()
	ps.SpaceComponent = common.SpaceComponent{Width: 30, Height: 30, Position: engo.Point{10, 10}}
	ps.RenderComponent = common.RenderComponent{Drawable: common.Circle{BorderWidth: 7, BorderColor: color.RGBA{0, 0, 255, 255}}, 
		Color: color.RGBA{0,0,0,255}}

	ps.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
			case *common.MouseSystem:
				sys.Add(&ps.mouseTracker.BasicEntity, &ps.mouseTracker.MouseComponent, nil, nil)
			case *common.RenderSystem:
				sys.Add(&ps.BasicEntity, &ps.RenderComponent, &ps.SpaceComponent)
		}
	}
}
