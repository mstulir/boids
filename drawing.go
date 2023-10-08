//Madison Stulir
//Boids HW
//Due October 10, 2022

package main

import (
	"canvas"
	"image"
	//"fmt"
)

//place your drawing code here.

//AnimateSystem takes a slice of Sky objects along with a canvas width
//parameter and generates a slice of images corresponding to drawing each Sky
//on a canvasWidth x canvasWidth canvas
func AnimateSystem(skies []Sky, canvasWidth, drawingFrequency int) []image.Image {
	images:=make([]image.Image,0)

	for i:=range skies {
		//fmt.Println("printing sky", i)
		if i%drawingFrequency==0{//only draw if current index of universe is divisible by some parameter frequency
			images=append(images,DrawToCanvas(skies[i],canvasWidth))
		}
	}
	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Sky
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(sky Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the skies and draw them.
	for _, b := range sky.boids {
		c.SetFillColor(canvas.MakeColor(223,227,202))
		cx := (b.position.x / sky.width) * float64(canvasWidth)
		cy := (b.position.y / sky.width) * float64(canvasWidth)
		r := 5.0
		c.Circle(cx, cy, r)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
