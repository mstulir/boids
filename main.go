//Madison Stulir
//Boids HW
//Due October 10, 2022

package main

import(
  "fmt"
  "gifhelper"
  "strconv"
  "os"
)

func main() {
	//Place your code here.

  /*
  //Hardcoding for ease of testing
  width:=2000.0
  initialSpeed:=1.0
  numBoids:=200
  maxBoidSpeed:=2.0
  proximity:=200.0
  separationFactor:=1.5
  alignmentFactor:=1.0
  cohesionFactor:=0.2
  numGens:=8000
  canvasWidth:=2000
  imageFrequency:=20
  timeStep:=1.0
  */

  //os.Args[1] is the number of boids
  numBoids,err1:=strconv.Atoi(os.Args[1])
  if err1!=nil{
		panic(err1)
	}
	if numBoids<0{
		panic("Negative number of boids given")
	}
  //os.Args[2] is the width of the skyWidth
  width,err2:=strconv.ParseFloat(os.Args[2],64)
	if err2!=nil{
		panic(err2)
	}
  //os.Args[3] is the initialSpeed of boids
  initialSpeed,err3:=strconv.ParseFloat(os.Args[3],64)
  if err3!=nil{
    panic(err3)
  }
  //os.Args[4] is the maxBoidSpeed
  maxBoidSpeed,err4:=strconv.ParseFloat(os.Args[4],64)
  if err4!=nil{
    panic(err4)
  }
  //os.Args[5] is the number of generations numGens
  numGens,err5:=strconv.Atoi(os.Args[5])
  if err5!=nil{
		panic(err5)
	}
  //os.Args[6] is the proximity of the boids to determine if forces act upon the boid
  proximity,err6:=strconv.ParseFloat(os.Args[6],64)
  if err6!=nil{
    panic(err6)
  }
  //os.Args[7] is the separationFactor
  separationFactor,err7:=strconv.ParseFloat(os.Args[7],64)
  if err7!=nil{
    panic(err7)
  }
  //os.Args[8] is the alignmentFactor
  alignmentFactor,err8:=strconv.ParseFloat(os.Args[8],64)
  if err8!=nil{
    panic(err8)
  }
  //os.Args[9] is the cohesionFactor
  cohesionFactor,err9:=strconv.ParseFloat(os.Args[9],64)
  if err9!=nil{
    panic(err9)
  }
  //os.Args[10] is the timeStep
  timeStep,err10:=strconv.ParseFloat(os.Args[10],64)
  if err10!=nil{
    panic(err10)
  }
  //os.Args[11] is the canvasWidth
  canvasWidth,err11:=strconv.Atoi(os.Args[11])
  if err11!=nil{
		panic(err11)
	}
  //os.Args[12] is the imageFrequency
  imageFrequency,err12:=strconv.Atoi(os.Args[12])
  if err12!=nil{
		panic(err12)
	}

  //Gnerate a random sky with the number of boids indicated and the characteristics input in command line
  initialSky:=GenerateInitialSky(width,initialSpeed,numBoids,maxBoidSpeed,proximity,separationFactor,alignmentFactor,cohesionFactor)
  //fmt.Println(initialSky)

  //Run Simulation
  skies:=SimulateBoids(initialSky, numGens, timeStep)
  //fmt.Println(skies[0])
  //fmt.Println(skies[10])
  fmt.Println("Simulation run successfully!")

  images:= AnimateSystem(skies,canvasWidth,imageFrequency)
  fmt.Println("Images drawn!")
  filename:="Boids"
  gifhelper.ImagesToGIF(images,filename)
  fmt.Println("Animated GIF produced. Exiting normally.")

}
