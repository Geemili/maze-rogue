package main

import (
	"github.com/geemili/maze-rogue/context"
	"github.com/geemili/maze-rogue/model"
	"github.com/geemili/maze-rogue/view"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

var (
	player *model.Player
	window *glfw.Window
	room   *model.Room
)

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.SetKeyCallback(onKey)

	player = &model.Player{3, 6, 6}
	room = &model.Room{5, 5, 6, 4}

	gl.Ortho(0, 40, 0, 30, -1, 3)

	for !window.ShouldClose() {
		render()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func render() {
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	view.RenderRoom(room)
	view.RenderPlayer(player)
}

func onKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch k {
	case glfw.KeyLeft:
		context.MovePlayer(player, []*model.Room{room}, -1, 0)
	case glfw.KeyRight:
		context.MovePlayer(player, []*model.Room{room}, 1, 0)
	case glfw.KeyUp:
		context.MovePlayer(player, []*model.Room{room}, 0, 1)
	case glfw.KeyDown:
		context.MovePlayer(player, []*model.Room{room}, 0, -1)
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}
