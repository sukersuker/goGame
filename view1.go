package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type view1 struct {
	player player
}

type player struct {
	rect        Rect
	speedOnAxis int
	speedOnDiag int
}

func NewView1() view1 {
	playerRect := Rect{
		X: 50,
		Y: 50,
		W: 100,
		H: 100,
	}
	p := player{
		rect:        playerRect,
		speedOnAxis: 3,
		speedOnDiag: 2,
	}

	return view1{
		player: p,
	}
}

func (v1 *view1) Render(renderer *sdl.Renderer, events *Events) {

	v1.player.calculateMovement(events)

	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	renderer.SetDrawColor(255, 0, 0, 255)
	playerSDLRect := v1.player.rect.GetSDLRect()
	renderer.DrawRect(playerSDLRect)
	renderer.FillRect(playerSDLRect)

	enemy := Rect{300, 300, 100, 100}
	enemySDLRect := enemy.GetSDLRect()
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawRect(enemySDLRect)
	renderer.FillRect(enemySDLRect)

	if v1.player.rect.Colision(enemy) {
		os.Exit(0)
	}

	renderer.Present()

	return
}

func (player *player) calculateMovement(events *Events) {
	diagonalMovement := (events.up != events.down) && (events.left != events.right)
	var speed int

	if diagonalMovement {
		speed = player.speedOnDiag
	} else {
		speed = player.speedOnAxis
	}

	if events.left {
		player.rect.X -= speed
	} else if events.right {
		player.rect.X += speed
	}

	if events.up {
		player.rect.Y -= speed
	} else if events.down {
		player.rect.Y += speed
	}
}
