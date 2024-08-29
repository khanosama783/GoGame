package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running      = true
	bkgColor     = rl.NewColor(147, 211, 196, 255)
	whiteColor   = rl.NewColor(255, 255, 255, 255)
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerSpeed  float32 = 3.0
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, whiteColor)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width/2, playerDest.Height/2), 0, whiteColor)
}

func input() {
	switch {
	case rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp):
		playerDest.Y -= playerSpeed
	case rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown):
		playerDest.Y += playerSpeed
	case rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft):
		playerDest.X -= playerSpeed
	case rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight):
		playerDest.X += playerSpeed
	}
}

func update() {
	running = !rl.WindowShouldClose()
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	drawScene()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "GoGame")
	rl.SetTargetFPS(60)
	grassSprite = rl.LoadTexture("Sprout-Lands -Sprites/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("Sprout-Lands -Sprites/Characters/Basic Charakter Spritesheet.png")
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}
