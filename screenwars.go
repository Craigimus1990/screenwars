package main

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/user/screenwars/systems"
)

type DefaultScene struct{}

type System interface {
	//run every frame, passed time in seconds since last frame
	Update(dt float32)

	//Removes an entity from the system
	Remove(ecs.BasicEntity)
}

// Type uniquely defines your game type
func (*DefaultScene) Type() string { return "screenwars" }

func (*DefaultScene) Preload() {}
func (*DefaultScene) Setup(w *ecs.World) {
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	
	player := &systems.PlayerSystem{}

	//ally := MyShape{BasicEntity: ecs.NewBasic()}
	//ally.SpaceComponent = common.SpaceComponent{Width: 30, Height: 30, Position: engo.Point{0, float32(40 + 30)}}
	//ally.RenderComponent = common.RenderComponent{Drawable: common.Circle{BorderWidth: 3, BorderColor: color.RGBA{255, 0, 0, 255}},
	//		Color: color.RGBA{0,0,0,255}}

	//enemy := MyShape{BasicEntity: ecs.NewBasic()}
	//enemy.SpaceComponent = common.SpaceComponent{Width: 30, Height: 30, Position: engo.Point{50, float32(40)}}
	//enemy.RenderComponent = common.RenderComponent{Drawable: common.Circle{BorderWidth: 3, BorderColor: color.RGBA{0, 255, 0, 255}}, 
	//	Color: color.RGBA{0,0,0,255}}


	w.AddSystem(player)
}

func (*DefaultScene) Exit() {
	log.Println("Exit event called; we can do whatever we want now")
	// Here if you want you can prompt the user if they're sure they want to close
	log.Println("Manually closing")
	engo.Exit()
}

func main() {
	opts := engo.RunOptions{
		Title:               "Exit Demo",
		OverrideCloseAction: true,
		Fullscreen:			 true,
		FPSLimit:			 60,
	}
	engo.Run(opts, &DefaultScene{})
}

