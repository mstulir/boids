//Madison Stulir
//Boids HW
//Due October 10, 2022

package main

import(
  "math"
  "math/rand"
)
//place your non-drawing functions here.


//first 2 functions serve to generate the random initial sky

//GenerateInitialSky creates a sky with all the inputs indicated from the command line.
//Input: sky parameters from the command line
//Output: an initial sky with all parameters set based on command line arguments
func GenerateInitialSky(skyWidth float64, initialSpeed float64, numBoids int, maxBoidSpeed float64, proximity float64, separationFactor, alignmentFactor,cohesionFactor float64) Sky {
  //initialize the sky object
  var Sky Sky
  //make a sky object and set the inputs which we dont have to change
  Sky.width=skyWidth
  Sky.maxBoidSpeed=maxBoidSpeed
  Sky.proximity=proximity
  Sky.separationFactor=separationFactor
  Sky.alignmentFactor=alignmentFactor
  Sky.cohesionFactor=cohesionFactor

  //initialize boids for the sky
  boids:=make([]Boid,numBoids) //of length numBoids
  for i:=range boids{
    boids[i]=GenerateBoid(initialSpeed,skyWidth)
  }
  Sky.boids=boids
  return Sky
}

//GenerateBoid initializes a single boid within the area of the skyWidth with a random position and a random direction using the speed given for every input boid (uniform)
//Input: an initial speed and a skywidth to plsce the position within
//Output: a bod with a random direction and position given the input speed are skyWidth area
func GenerateBoid(initialSpeed float64, skyWidth float64) Boid {
  //define a new boid
  var b Boid

  //assign its features
  b.acceleration.x=0
  b.acceleration.y=0
  //generate position on board that is within the width of the board as a float
  b.position.x=rand.Float64()*skyWidth
  b.position.y=rand.Float64()*skyWidth
  //generate velocity vectors that add up to the initial speed to give the boid a direction
  b.velocity.x=rand.Float64()*initialSpeed
  b.velocity.y=math.Sqrt(initialSpeed*initialSpeed - b.velocity.x*b.velocity.x)

  //make sure boids travel in all 4 directions
  //generate a random variable such that it is in the range 0-4
  x:=rand.Float64()*4
  if x<1{
    //if this case, x will be negative
    b.velocity.x*=-1
  } else if x<2{
    //if this case, y will be negative
    b.velocity.y*=-1
  } else if x<3{
    //if this case, x and y will be negative
    b.velocity.x*=-1
    b.velocity.y*=-1
  } // if we get here, leave them alone
  //all of the boids parameters have been assigned to the initial value
  return b

}

//SimulateBoids takes in an initialSky, numGens which is a number of generations to simulate and a timeStep to update by each generation
//It returns a slice of sky objects corresponding to each generation
//Input: an initialSky , a number of generations numGens and a timeStep in seconds for the generations
//Output: a slice of Sky objects updated for each timepoint
func SimulateBoids(initialSky Sky, numGens int, timeStep float64) []Sky {
  //generate a slice of skies of length numGens+1
  skies:=make([]Sky,numGens+1)

  //set first sky to the initial sky
  skies[0]=initialSky
  //loop through numGens+1 and call update sky
  for i:=1;i<=numGens;i++{
    skies[i]=UpdateSky(skies[i-1], timeStep)
  }
  return skies
}

//UpdateSky takes in a sky for the previous time point and a timestep in seconds by which to update the sky according to the boid flying algorithm
//Input: a sky object and a time float in seconds
//Output: a new sky object that has had the boid algorithm applied to it for the specifed amount of time in seconds
func UpdateSky(sky Sky, time float64) Sky {
  //create a deep copy of the sky input to edit on and return
  newSky:=CopySky(sky)
  //loop through the bodies in the sky and call update boid
  for i:=range newSky.boids{
    newSky.boids[i]=UpdateBoid(sky,sky.boids[i], time)
  }
  return newSky
}

//CopySky creates a deep copy of every item in the current sky
//Input: a Sky object
//Output: a new sky object, all of whose fields are copied from the sky to the newSkys fields
func CopySky(currentSky Sky) Sky {
  var newSky Sky

  // define features besides boids
  newSky.width=currentSky.width
  newSky.maxBoidSpeed=currentSky.maxBoidSpeed
  newSky.proximity=currentSky.proximity
  newSky.separationFactor=currentSky.separationFactor
  newSky.alignmentFactor=currentSky.alignmentFactor
  newSky.cohesionFactor=currentSky.cohesionFactor

  //define boids
  numBoids:=len(currentSky.boids)
  newSky.boids=make([]Boid,numBoids)

  for i:=range currentSky.boids{
    newSky.boids[i].position.x=currentSky.boids[i].position.x
    newSky.boids[i].position.y=currentSky.boids[i].position.y
    newSky.boids[i].velocity.x=currentSky.boids[i].velocity.x
    newSky.boids[i].velocity.y=currentSky.boids[i].velocity.y
    newSky.boids[i].acceleration.x=currentSky.boids[i].acceleration.x
    newSky.boids[i].acceleration.y=currentSky.boids[i].acceleration.y
  }
  return newSky
}

//UpdateBoid takes in a current Sky, a boid object and a time step and updates the position, velocity and acceleration of the boid by the timeStep amount in seconds
//Input: a sky, a boid and an amount of time
//Output: a new boid, updated by the time step
func UpdateBoid(sky Sky, b Boid, time float64) Boid {
  //create a copy of the boid
  newBoid:=CopyBoid(b)
  //create vectors for each function to add to
  var forceFromSep OrderedPair
  forceFromSep.x=0
  forceFromSep.y=0
  var forceFromAlign OrderedPair
  forceFromAlign.x=0
  forceFromAlign.y=0
  var forceFromCoh OrderedPair
  forceFromCoh.x=0
  forceFromCoh.y=0
  //n is the number of boids within the threshold distance. if a boid within distance, this will be incremented
  n:=0
  //call update functions for forces from 3 factors
    for i:=range(sky.boids){
      //update boid against from all forces from other boids besides itself
      if sky.boids[i]!=b{
        distance:=CalcDistance(b.position, sky.boids[i].position)
        //if the distance is less than the proximity, then we update the force values (otherwise do not update force from this boid)
        if distance<sky.proximity{
          n+=1
          sepForce:=UpdateSeparation(b,sky.boids[i], sky.separationFactor,distance)
          forceFromSep.x+=sepForce.x
          forceFromSep.y+=sepForce.y
          alignForce:=UpdateAlignment(sky.boids[i], sky.alignmentFactor,distance)
          forceFromAlign.x+=alignForce.x
          forceFromAlign.y+=alignForce.y
          cohForce:=UpdateCohesion(b,sky.boids[i], sky.cohesionFactor,distance)
          forceFromCoh.x+=cohForce.x
          forceFromCoh.y+=cohForce.y
        }
      }
    }
    //if there was at least 1 boid within proximity of the deisred boid, we need to divide by n
    //this step would be problematic if n was zero (cant divide by zero)
    if n>0{
      //forces need to also be divided componentwise
      forceFromSep.x=forceFromSep.x/float64(n)
      forceFromSep.y=forceFromSep.y/float64(n)
      forceFromAlign.x=forceFromAlign.x/float64(n)
      forceFromAlign.y=forceFromAlign.y/float64(n)
      forceFromCoh.x=forceFromCoh.x/float64(n)
      forceFromCoh.y=forceFromCoh.y/float64(n)
    } //if we get here, vectors are all zero

  //force equals accleration due to unit mass (1 and F=ma)
  newBoid.acceleration=SumForces(forceFromSep,forceFromAlign,forceFromCoh)
  newBoid.velocity=UpdateVelocity(sky.maxBoidSpeed,newBoid, time)

  //update position vectors
  newBoid.position=UpdatePosition(newBoid,time,sky.width)

  return newBoid
}

//CopyBoid creates a deep copy of a Boid object by copying all of its attributes
//Input: a Boid object
//Output: a copy of the Boid object
func CopyBoid(b Boid) Boid {
  var newBoid Boid
  newBoid.position.x=b.position.x
  newBoid.position.y=b.position.y
  newBoid.velocity.x=b.velocity.x
  newBoid.velocity.y=b.velocity.y
  newBoid.acceleration.x=b.acceleration.x
  newBoid.acceleration.y=b.acceleration.y

  return newBoid
}

//CalcDistance calculates the distance between 2 boids
//Input: 2 distance points p1 and p2 with x and y coordinates
//Output: a float64 of the distance between the points
func CalcDistance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

//UpdateSeparation accounts for the repulsion force between 2 boids based on their positions
//Input: 2 boids, a separationFactor float64 value and the distance between the 2 Boids
//Output: an OrderedPair object corresponding to the force from separation
func UpdateSeparation(b Boid, otherBoid Boid, separationFactor float64, distance float64) OrderedPair {
  var sep OrderedPair
  sep.x=separationFactor*(b.position.x-otherBoid.position.x)/(distance*distance)
  sep.y=separationFactor*(b.position.y-otherBoid.position.y)/(distance*distance)

  return sep
}

//UpdateAlignment accounts for the force of alignment (flying in the same direction) of 2 boids
//Input: 2 boids, an alignmentFactor float64 value and the distance between the 2 Boids
//Output: an OrderedPair object corresponding to the force from alignment
func UpdateAlignment(b Boid, alignmentFactor float64,distance float64) OrderedPair {
  var align OrderedPair
  align.x=alignmentFactor*b.velocity.x/distance
  align.y=alignmentFactor*b.velocity.y/distance

  return align
}

//UpdateAlignment accounts for the force of cohesion (staying close together in a flock) of 2 boids
//Input: 2 boids, a cohesionFactor float64 value and the distance between the 2 Boids
//Output: an OrderedPair object corresponding to the force from cohesion
func UpdateCohesion(b Boid,otherBoid Boid, cohesionFactor float64,distance float64) OrderedPair {
  var coh OrderedPair
  coh.x=cohesionFactor*(otherBoid.position.x-b.position.x)/distance
  coh.y=cohesionFactor*(otherBoid.position.y-b.position.y)/distance

  return coh
}

//SumForces takes in several Ordered pair forces and adds them componentwise
//Input: 3 force OrderedPair's
//Output: a single OrderedPair corresponding to the sum of the forces
func SumForces(forceFromSep,forceFromAlign,forceFromCoh OrderedPair) OrderedPair {
  var sum OrderedPair
  sum.x=forceFromSep.x+forceFromAlign.x+forceFromCoh.x
  sum.y=forceFromSep.y+forceFromAlign.y+forceFromCoh.y

  return sum
}

//Compute speed takes x and y vector components and converts it into the speed
//Input: a boid containing x and y velocity vectors
//Output: the speed of the boid as a float64
func ComputeSpeed(velocity OrderedPair) float64 {
  speed:=math.Sqrt(velocity.x*velocity.x+velocity.y*velocity.y)
  return speed
}

//ConvertSpeedToMax takes in a speed that is greated than the max speed, and its components and scales down the components to achieve the max speed in the same direction
//Input: a maxboidspeed float, the currentspeed float and the velocity components as an ordered pair
//Output: the scaled down velocity components in the same direction such that the speed is now the max speed
func ConvertSpeedToMax(maxBoidSpeed float64, currentSpeed float64, velocity OrderedPair) OrderedPair {
  //define a new ordered pair vel
  var vel OrderedPair
  scaleFactor:=maxBoidSpeed/currentSpeed
  vel.x=scaleFactor*velocity.x
  vel.y=scaleFactor*velocity.y

  return vel
}

//UpdateVelocity creates the new velocity based on a timepoint and an updated acceleration
//Input: a boid object and a timepoint
//Output: the new velocity of the object as component vectors
func UpdateVelocity(maxSpeed float64, b Boid,time float64) OrderedPair {
	var vel OrderedPair
	//new velocity is current velocity + new acceleration *time
	vel.x=b.velocity.x + b.acceleration.x*time
	vel.y=b.velocity.y + b.acceleration.y*time
  var velocity OrderedPair
  speed:=ComputeSpeed(vel)
  if speed>maxSpeed{
    velocity=ConvertSpeedToMax(maxSpeed, speed, vel)
  } else {
    velocity=vel
  }
	return velocity
}
//UpdatePosition
//Input: a body and a timestep float64
//Output: the orderedpair corresponding to the updated position of the body after a single time step using the bodies current acceleration and velocity
func UpdatePosition(b Boid, time float64, width float64) OrderedPair {
	var pos OrderedPair
  //uses new velocity and acceleration to update the old not updated position (we use the old position in this calculation)
	pos.x=0.5*b.acceleration.x*time*time +b.velocity.x*time +b.position.x
	pos.y=0.5*b.acceleration.y*time*time +b.velocity.y*time +b.position.y

	//need to account for if we are off the edge of the sky. if we are we need to loop back to other side.
  pos.x=CheckPosX(pos.x,width)
  pos.y=CheckPosY(pos.y,width)

  return pos
}

//CheckPosX determines if the boid position is off the sky in the x direction, and loops it back onto the sky on the other side if so
//Input: an x coordinate and the width of the sky as float64 values
//Output: the position of x such that it remains within the sky
func CheckPosX(x, width float64) float64 {
if x<0{
  x=width+x
} else if x>width {
  x=x-width
}
return x
}

//CheckPosY determines if the boid position is off the sky in the y direction, and loops it back onto the sky on the other side if so
//Input: a y coordinate and the width of the sky as float64 values
//Output: the position of y such that it remains within the sky
func CheckPosY(y, width float64) float64 {
  if y<0{
  y=width+y
  } else if y>width {
  y=y-width
  }
  return y
}
