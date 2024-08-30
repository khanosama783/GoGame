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
	musicPaused bool
	music rl.Music
	cam rl.Camera2D
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
	case rl.IsKeyPressed(rl.KeyQ):
		musicPaused = !musicPaused
	}
}

func update() {
	running = !rl.WindowShouldClose()
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
	cam.Target = rl.NewVector2(float32(playerDest.X - (playerDest.Width/2)), float32(playerDest.Y - (playerDest.Height/2)))
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)
	drawScene()
	rl.EndMode2D()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "GoGame")
	rl.SetTargetFPS(60)
	grassSprite = rl.LoadTexture("Sprout-Lands -Sprites/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("Sprout-Lands -Sprites/Characters/Basic Charakter Spritesheet.png")
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/Neon Doom - Steven O'Brien (Must Credit, CC-BY, www.steven-obrien.net).mp3")
	musicPaused = false
	rl.PlayMusicStream(music)
	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X - (playerDest.Width/2)), float32(playerDest.Y - (playerDest.Height/2))), 0.0, 1.0)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
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
