package main

import (
	"fmt"
	"gophys/particle"
	"gophys/rigid_body"
	"gophys/utils"
	"gophys/vector"
)

// Initializes n number of particles with random positions,
// zero velocities and 1kg mass
func initRandomParticles(n int) []*particle.Particle {
	var particles []*particle.Particle
	for range n {
		p := particle.RandomParticle(50, 50, 1, 1)
		particles = append(particles, p)
	}
	return particles
}

func printParticles(particles []*particle.Particle) {
	for i, p := range particles {
		fmt.Printf("particle[%d] (%.2f, %.2f, %.2f, %.2f)\n",
			i, p.X(), p.Y(), p.VelocityX(), p.VelocityY())
	}
}

func initRandomRigidBodies(n int) []*rigid_body.RigidBody {
	var rigidBodies []*rigid_body.RigidBody
	for range n {
		r := rigid_body.MakeRandomRigidBody(50, 50, 2, 2)
		rigidBodies = append(rigidBodies, r)
	}
	return rigidBodies
}

func printRigidBodies(rigidBodies []*rigid_body.RigidBody) {
	for i, r := range rigidBodies {
		fmt.Printf("rigid body[%d] p = (%.2f, %.2f), a = %.2f\n",
			i, r.X(), r.Y(), r.Angle())
	}
}

func RunParticleSimulation(duration int) {
	currentTime := 0
	dt := 1 // each update takes 1 second

	particles := initRandomParticles(1)
	printParticles(particles)

	for currentTime < duration {
		utils.Sleep(dt)

		for _, p := range particles {
			force := p.ComputeForce()
			acceleration := vector.MakeVector2(force.X()/p.Mass(), force.Y()/p.Mass())
			p.IncrementVelocityX(acceleration.X() * float64(dt))
			p.IncrementVelocityY(acceleration.Y() * float64(dt))
			p.IncrementPositionX(p.VelocityX() * float64(dt))
			p.IncrementPositionY(p.VelocityY() * float64(dt))
		}

		printParticles(particles)
		currentTime += dt
	}
}

func RunRigidBodySimulation(duration int) {
	currentTime := 0
	dt := 1

	rigidBodies := initRandomRigidBodies(1)
	printRigidBodies(rigidBodies)

	for currentTime < duration {
		utils.Sleep(dt)

		for _, r := range rigidBodies {
			r.ComputeForceAndTorque()
			linearAcceleration := vector.MakeVector2(r.ForceX() / r.Mass(), r.ForceY() / r.Mass())
			r.IncrementLinVelX(linearAcceleration.X() * float64(dt))
			r.IncrementLinVelY(linearAcceleration.Y() * float64(dt))
			r.IncrementPositionX(r.LinVelX() * float64(dt))
			r.IncrementPositionY(r.LinVelY() * float64(dt))

			angularAcceleration := r.Torque() / r.MomentOfInertia()
			r.IncrementAngularVelocity(angularAcceleration * float64(dt))
			r.IncrementAngle(r.AngularVelocity() * float64(dt))
		}

		printRigidBodies(rigidBodies)
		currentTime += dt
	}
}

func main() {
	RunRigidBodySimulation(10)
}