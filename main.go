package main

import (
	"embed"
	"fmt"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	camera "github.com/tducasse/ebiten-camera"
)

//go:embed assets
var EmbeddedAssets embed.FS

type cameraDemoGame struct {
	background     *ebiten.Image
	displayedLevel *ebiten.Image
	cameraView     *camera.Camera
	player         player
	drawOps        ebiten.DrawImageOptions
}

type player struct {
	pict *ebiten.Image
	x, y int
}

func (demo cameraDemoGame) Update() error {
	return nil
}

func (demo cameraDemoGame) Draw(screen *ebiten.Image) {
	//draw to the world at first
	//first draw background
	demo.drawOps.GeoM.Reset()
	demo.displayedLevel.DrawImage(demo.background, &demo.drawOps)
	//next draw player
	demo.drawOps.GeoM.Reset()
	demo.drawOps.GeoM.Translate(float64(demo.player.x), float64(demo.player.y))
	demo.displayedLevel.DrawImage(demo.player.pict, &demo.drawOps)

	//finally draw to the screen
	demo.cameraView.Draw(demo.displayedLevel, screen)
}

func (demo cameraDemoGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(800, 720)
	ourCamera := camera.Init(800, 720)
	backgroundImage := LoadEmbeddedImage("background", "BACKGROUND4.png")
	displayWorld := ebiten.NewImage(backgroundImage.Bounds().Dx(), backgroundImage.Bounds().Dy())
	theplayer := player{
		pict: LoadEmbeddedImage("sprites", "player.png"),
		x:    100,
		y:    100,
	}
	game := &cameraDemoGame{
		background:     backgroundImage,
		displayedLevel: displayWorld,
		cameraView:     ourCamera,
		player:         theplayer,
	}
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadEmbeddedImage(folderName string, imageName string) *ebiten.Image {
	embeddedFile, err := EmbeddedAssets.Open(path.Join("assets", folderName, imageName))
	if err != nil {
		log.Fatal("failed to load embedded image ", imageName, err)
	}
	ebitenImage, _, err := ebitenutil.NewImageFromReader(embeddedFile)
	if err != nil {
		fmt.Println("Error loading tile image:", imageName, err)
	}
	return ebitenImage
}
