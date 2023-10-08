//Madison Stulir
//Boids HW
//Due October 10, 2022

//Testing
package main

import(
  "testing"
  "fmt"
  "math"
  "gifhelper"
)

func TestUpdatePosition(t *testing.T) {
  type test struct {
    boid Boid
    timeStep float64
    width float64
    answer OrderedPair
  }
  tests:=make([]test,2)

  //assign test values hard coded
  //this test case shows the sky is a torus, and properly relocates in both x and y
  tests[0].boid.position.x=1999
  tests[0].boid.position.y=1999
  tests[0].boid.velocity.x=1
  tests[0].boid.velocity.y=1
  tests[0].boid.acceleration.x=0.2
  tests[0].boid.acceleration.y=0.2
  tests[0].timeStep=1.0
  tests[0].width=2000
  tests[0].answer.x=0.1
  tests[0].answer.y=0.1

  //test a normal case in the middle of the board
  tests[1].boid.position.x=100
  tests[1].boid.position.y=1000
  tests[1].boid.velocity.x=1
  tests[1].boid.velocity.y=0.5
  tests[1].boid.acceleration.x=0.2
  tests[1].boid.acceleration.y=0.7
  tests[1].timeStep=1.0
  tests[1].width=2000
  tests[1].answer.x=101.1
  tests[1].answer.y=1000.85

  for i, test := range tests {
		outcome := UpdatePosition(test.boid,test.timeStep,test.width)
    var numDigits uint = 4
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct position is %v,%v", i, outcome.x, outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The position is", test.answer)
		}
	}
}

//test convert speed to max
func TestConvertSpeedToMax(t *testing.T) {
  type test struct {
    maxBoidSpeed float64
    currentSpeed float64
    velocity OrderedPair
    answer OrderedPair
  }
  tests:=make([]test,4)

  //assign test values hard coded
  tests[0].maxBoidSpeed=2
  tests[0].currentSpeed=20
  tests[0].velocity.x=16
  tests[0].velocity.y=12
  tests[0].answer.x=1.6
  tests[0].answer.y=1.2

  tests[1].maxBoidSpeed=5
  tests[1].currentSpeed=20
  tests[1].velocity.x=16
  tests[1].velocity.y=12
  tests[1].answer.x=4
  tests[1].answer.y=3

  tests[2].maxBoidSpeed=2
  tests[2].currentSpeed=10.19803903
  tests[2].velocity.x=-2
  tests[2].velocity.y=-10
  tests[2].answer.x=-0.3922
  tests[2].answer.y=-1.9612

  tests[3].maxBoidSpeed=2
  tests[3].currentSpeed=10.19803903
  tests[3].velocity.x=-2
  tests[3].velocity.y=10
  tests[3].answer.x=-0.3922
  tests[3].answer.y=1.9612

  for i, test := range tests {
		outcome := ConvertSpeedToMax(test.maxBoidSpeed,test.currentSpeed,test.velocity)
    var numDigits uint = 4
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct speed is %v,%v", i, outcome.x, outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The speed is", test.answer)
		}
	}
}

//test separation force
func TestSeparationForce(t *testing.T) {
  type test struct {
    sky Sky
    distance float64
    answer OrderedPair
  }
  tests:=make([]test,1)
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=2.0
  tests[0].sky.boids[0].position.y=6.0
  tests[0].sky.boids[1].position.x=23.0
  tests[0].sky.boids[1].position.y=13.0
  tests[0].sky.separationFactor=1.5
  tests[0].distance=math.Sqrt((tests[0].sky.boids[0].position.x-tests[0].sky.boids[1].position.x)*(tests[0].sky.boids[0].position.x-tests[0].sky.boids[1].position.x)+(tests[0].sky.boids[0].position.y-tests[0].sky.boids[1].position.y)*(tests[0].sky.boids[0].position.y-tests[0].sky.boids[1].position.y))
  tests[0].answer.x=-0.0643
  tests[0].answer.y=-0.0214


  for i, test := range tests {
		outcome := UpdateSeparation(test.sky.boids[0],test.sky.boids[1],test.sky.separationFactor,test.distance)
    var numDigits uint = 4
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct separation force is %v,%v", i, outcome.x, outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The separation force is", test.answer)
		}
	}
}

//test alignment force
func TestAlignmentForce(t *testing.T) {
  type test struct {
    sky Sky
    distance float64
    answer OrderedPair
  }
  tests:=make([]test,1)
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=2.0
  tests[0].sky.boids[0].position.y=6.0
  tests[0].sky.boids[1].position.x=23.0
  tests[0].sky.boids[1].position.y=13.0
  tests[0].sky.boids[1].velocity.x=2.0
  tests[0].sky.boids[1].velocity.y=2.0
  tests[0].sky.alignmentFactor=1.0
  tests[0].distance=math.Sqrt((tests[0].sky.boids[0].position.x-tests[0].sky.boids[1].position.x)*(tests[0].sky.boids[0].position.x-tests[0].sky.boids[1].position.x)+(tests[0].sky.boids[0].position.y-tests[0].sky.boids[1].position.y)*(tests[0].sky.boids[0].position.y-tests[0].sky.boids[1].position.y))
  tests[0].answer.x=0.0904
  tests[0].answer.y=0.0904


  for i, test := range tests {
		outcome := UpdateAlignment(test.sky.boids[1],test.sky.alignmentFactor,test.distance)
    var numDigits uint = 4
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct alignment force is %v,%v", i, outcome.x, outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The alignment force is", test.answer)
		}
	}
}
//test cohesion force
func TestCohesionForce(t *testing.T) {
  type test struct {
    sky Sky
    distance float64
    answer OrderedPair
  }
  tests:=make([]test,1)
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=2.0
  tests[0].sky.boids[0].position.y=6.0
  tests[0].sky.boids[1].position.x=23.0
  tests[0].sky.boids[1].position.y=13.0
  tests[0].sky.boids[1].velocity.x=2.0
  tests[0].sky.boids[1].velocity.y=2.0
  tests[0].sky.cohesionFactor=0.015
  tests[0].distance=math.Sqrt((tests[0].sky.boids[0].position.x-tests[0].sky.boids[1].position.x)*(tests[0].sky.boids[0].position.x-tests[0].sky.boids[1].position.x)+(tests[0].sky.boids[0].position.y-tests[0].sky.boids[1].position.y)*(tests[0].sky.boids[0].position.y-tests[0].sky.boids[1].position.y))
  tests[0].answer.x=0.0142
  tests[0].answer.y=0.0047


  for i, test := range tests {
		outcome := UpdateCohesion(test.sky.boids[0],test.sky.boids[1],test.sky.cohesionFactor,test.distance)
    var numDigits uint = 4
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct cohesion force is %v,%v", i, outcome.x, outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The cohesion force is", test.answer)
		}
	}
}
//test calc distance
func TestCalcDistance(t *testing.T) {
  type test struct {
    sky Sky
    answer float64
  }
  tests:=make([]test,2)
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=3.0
  tests[0].sky.boids[0].position.y=4.0
  tests[0].sky.boids[1].position.x=32.0
  tests[0].sky.boids[1].position.y=57.0
  tests[0].answer=60.4152

  tests[1].sky.boids=make([]Boid,2)
  tests[1].sky.boids[0].position.x=-5.0
  tests[1].sky.boids[0].position.y=2.0
  tests[1].sky.boids[1].position.x=500.0
  tests[1].sky.boids[1].position.y=12.0
  tests[1].answer=505.0990

  for i, test := range tests {
		outcome := CalcDistance(test.sky.boids[0].position,test.sky.boids[1].position)
    var numDigits uint = 4
    outcome=roundFloat(outcome,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v, and the correct distance is %v", i, outcome, test.answer)
		} else {
			fmt.Println("Correct! The distance is", test.answer)
		}
	}
}
//test update velocity
func TestUpdateVelocity(t *testing.T) {
  type test struct {
    sky Sky
    timeStep float64
    answer OrderedPair
  }
  tests:=make([]test,2)
  tests[0].sky.boids=make([]Boid,1)
  tests[0].sky.boids[0].velocity.x=.3
  tests[0].sky.boids[0].velocity.y=.4
  tests[0].sky.boids[0].acceleration.x=.1
  tests[0].sky.boids[0].acceleration.y=.5
  tests[0].sky.maxBoidSpeed=2
  tests[0].timeStep=1
  tests[0].answer.x=0.4
  tests[0].answer.y=0.9

  tests[1].sky.boids=make([]Boid,1)
  tests[1].sky.boids[0].velocity.x=2.0
  tests[1].sky.boids[0].velocity.y=-0.1
  tests[1].sky.boids[0].acceleration.x=.1
  tests[1].sky.boids[0].acceleration.y=.5
  tests[1].sky.maxBoidSpeed=2
  tests[1].timeStep=1
  tests[1].answer.x=2.0
  tests[1].answer.y=0.4

  for i, test := range tests {
		outcome := UpdateVelocity(test.sky.maxBoidSpeed,test.sky.boids[0],test.timeStep)
    var numDigits uint = 1
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct velocity is %v,%v", i, outcome.x,outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The velocity is", test.answer)
		}
	}
}
//test sum forces
func TestSumForces(t *testing.T) {
  type test struct {
    forceFromSeparation OrderedPair
    forceFromAlignment OrderedPair
    forceFromCohesion OrderedPair
    answer OrderedPair
  }
  tests:=make([]test,1)
  tests[0].forceFromSeparation.x=-4.0
  tests[0].forceFromSeparation.y=30.0
  tests[0].forceFromAlignment.x=-7.0
  tests[0].forceFromAlignment.y=5.0
  tests[0].forceFromCohesion.x=3.0
  tests[0].forceFromCohesion.y=0.0
  tests[0].answer.x=-8.0
  tests[0].answer.y=35

  for i, test := range tests {
		outcome := SumForces(test.forceFromSeparation,test.forceFromAlignment,test.forceFromCohesion)
    var numDigits uint = 1
    outcome.x=roundFloat(outcome.x,numDigits)
    outcome.y=roundFloat(outcome.y,numDigits)
		if outcome != test.answer {
			t.Errorf("Error! For input test dataset %d, your code gives %v,%v and the correct sum is %v,%v", i, outcome.x,outcome.y, test.answer.x,test.answer.y)
		} else {
			fmt.Println("Correct! The sum is", test.answer)
		}
	}

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}


func TestAlignment (t *testing.T) {
  type test struct {
    sky Sky
    numGens int
    timeStep float64
  }
  fmt.Println("Testing Alignment")
  //run simulateBoids with only an appropriate alignmentFactor (no cohesion or separation), and input only 2 boids with specific positions
  tests:=make([]test,1)
  tests[0].sky.alignmentFactor=1.0
  tests[0].sky.cohesionFactor=0
  tests[0].sky.separationFactor=0
  tests[0].sky.proximity=200
  tests[0].sky.maxBoidSpeed=2.0
  tests[0].sky.width=2000
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=1100
  tests[0].sky.boids[0].position.y=1000
  tests[0].sky.boids[0].velocity.x=1
  tests[0].sky.boids[0].velocity.y=1
  tests[0].sky.boids[0].acceleration.x=0.0
  tests[0].sky.boids[0].acceleration.y=0.0

  tests[0].sky.boids[1].position.x=1000
  tests[0].sky.boids[1].position.y=1000
  tests[0].sky.boids[1].velocity.x=1
  tests[0].sky.boids[1].velocity.y=-1
  tests[0].sky.boids[1].acceleration.x=0.0
  tests[0].sky.boids[1].acceleration.y=0.0

  tests[0].numGens=8000
  tests[0].timeStep=1.0
  canvasWidth:=2000
  imageFrequency:=20
  for _, test := range tests {
    skies:=SimulateBoids(test.sky, test.numGens, test.timeStep)
    fmt.Println("Simulation run successfully!")
    images:= AnimateSystem(skies,canvasWidth,imageFrequency)
    fmt.Println("Images drawn!")
    filename:="TestAlignment"
    gifhelper.ImagesToGIF(images,filename)
    fmt.Println("Animated GIF produced.")
  }

}

func TestCohesion (t *testing.T) {
  type test struct {
    sky Sky
    numGens int
    timeStep float64
  }
  fmt.Println("Testing Cohesion")
  //run simulateBoids with only an appropriate cohesionFactor (no alignment or separation), and input only 2 boids with specific positions
  tests:=make([]test,1)
  tests[0].sky.alignmentFactor=0.0
  tests[0].sky.cohesionFactor=0.015
  tests[0].sky.separationFactor=0
  tests[0].sky.proximity=200
  tests[0].sky.maxBoidSpeed=2.0
  tests[0].sky.width=2000
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=1100
  tests[0].sky.boids[0].position.y=1000
  tests[0].sky.boids[0].velocity.x=1
  tests[0].sky.boids[0].velocity.y=1
  tests[0].sky.boids[0].acceleration.x=0.0
  tests[0].sky.boids[0].acceleration.y=0.0

  tests[0].sky.boids[1].position.x=1000
  tests[0].sky.boids[1].position.y=1000
  tests[0].sky.boids[1].velocity.x=1
  tests[0].sky.boids[1].velocity.y=-1
  tests[0].sky.boids[1].acceleration.x=0.0
  tests[0].sky.boids[1].acceleration.y=0.0

  tests[0].numGens=8000
  tests[0].timeStep=1.0
  canvasWidth:=2000
  imageFrequency:=20
  for _, test := range tests {
    skies:=SimulateBoids(test.sky, test.numGens, test.timeStep)
    fmt.Println("Simulation run successfully!")
    images:= AnimateSystem(skies,canvasWidth,imageFrequency)
    fmt.Println("Images drawn!")
    filename:="TestCohesion"
    gifhelper.ImagesToGIF(images,filename)
    fmt.Println("Animated GIF produced.")
  }
}

func TestSeparation (t *testing.T) {
  type test struct {
    sky Sky
    numGens int
    timeStep float64
  }
  fmt.Println("Testing Separation")
  //run simulateBoids with only an appropriate separationFactor (no cohesion or alignment), and input only 2 boids with specific positions
  tests:=make([]test,1)
  tests[0].sky.alignmentFactor=0.0
  tests[0].sky.cohesionFactor=0
  tests[0].sky.separationFactor=1.5
  tests[0].sky.proximity=200
  tests[0].sky.maxBoidSpeed=2.0
  tests[0].sky.width=2000
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=1100
  tests[0].sky.boids[0].position.y=1000
  tests[0].sky.boids[0].velocity.x=1
  tests[0].sky.boids[0].velocity.y=1
  tests[0].sky.boids[0].acceleration.x=0.0
  tests[0].sky.boids[0].acceleration.y=0.0

  tests[0].sky.boids[1].position.x=1000
  tests[0].sky.boids[1].position.y=1000
  tests[0].sky.boids[1].velocity.x=1
  tests[0].sky.boids[1].velocity.y=-1
  tests[0].sky.boids[1].acceleration.x=0.0
  tests[0].sky.boids[1].acceleration.y=0.0

  tests[0].numGens=8000
  tests[0].timeStep=1.0
  canvasWidth:=2000
  imageFrequency:=20
  for _, test := range tests {
    skies:=SimulateBoids(test.sky, test.numGens, test.timeStep)
    fmt.Println("Simulation run successfully!")
    images:= AnimateSystem(skies,canvasWidth,imageFrequency)
    fmt.Println("Images drawn!")
    filename:="TestSeparation"
    gifhelper.ImagesToGIF(images,filename)
    fmt.Println("Animated GIF produced.")
  }
}

func TestBoids (t *testing.T) {
  type test struct {
    sky Sky
    numGens int
    timeStep float64
  }
  fmt.Println("Testing Boid Simulation of 2 boids")
  //run simulateBoids with only an appropriate alignmentFactor (no cohesion or separation), and input only 2 boids with specific positions
  tests:=make([]test,1)
  tests[0].sky.alignmentFactor=1.0
  tests[0].sky.cohesionFactor=0.00015
  tests[0].sky.separationFactor=1.5
  tests[0].sky.proximity=200
  tests[0].sky.maxBoidSpeed=2.0
  tests[0].sky.width=2000
  tests[0].sky.boids=make([]Boid,2)
  tests[0].sky.boids[0].position.x=1100
  tests[0].sky.boids[0].position.y=1000
  tests[0].sky.boids[0].velocity.x=1
  tests[0].sky.boids[0].velocity.y=1
  tests[0].sky.boids[0].acceleration.x=0.0
  tests[0].sky.boids[0].acceleration.y=0.0

  tests[0].sky.boids[1].position.x=1000
  tests[0].sky.boids[1].position.y=1000
  tests[0].sky.boids[1].velocity.x=1
  tests[0].sky.boids[1].velocity.y=-1
  tests[0].sky.boids[1].acceleration.x=0.0
  tests[0].sky.boids[1].acceleration.y=0.0

  tests[0].numGens=8000
  tests[0].timeStep=1.0
  canvasWidth:=2000
  imageFrequency:=20
  for _, test := range tests {
    skies:=SimulateBoids(test.sky, test.numGens, test.timeStep)
    fmt.Println("Simulation run successfully!")
    images:= AnimateSystem(skies,canvasWidth,imageFrequency)
    fmt.Println("Images drawn!")
    filename:="Test2Boids"
    gifhelper.ImagesToGIF(images,filename)
    fmt.Println("Animated GIF produced.")
  }
}
