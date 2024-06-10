package particle

import (
	"gophys/utils"
	"gophys/vector"
)

type Particle struct {
	position vector.Vector2
	velocity vector.Vector2
	mass     float64
}

func MakeParticle(position vector.Vector2, velocity vector.Vector2, mass float64) *Particle {
	return &Particle{
		position: position,
		velocity: velocity,
		mass:     mass,
	}
}

// Initializes all particles with random positions, zero veloticies and 1kg mass.
func RandomParticle(maxX, maxY, maxVelocity, maxMass float64) *Particle {
	return &Particle{
		position: *vector.MakeVector2(
			utils.RandomFloat64Range(0, maxX),
			utils.RandomFloat64Range(0, maxY),
		),
		velocity: *vector.MakeVector2(
			0, 0,
		),
		mass: 1,
	}
}

func (p *Particle) X() float64 {
	return p.position.X()
}

func (p *Particle) Y() float64 {
	return p.position.Y()
}

func (p *Particle) Mass() float64 {
	return p.mass
}

func (p *Particle) SetPositionX(x float64) {
	p.position.SetX(x)
}

func (p *Particle) SetPositionY(y float64) {
	p.position.SetY(y)
}

func (p *Particle) IncrementPositionX(inc float64) {
	p.SetPositionX(p.X() + inc)
}

func (p *Particle) IncrementPositionY(inc float64) {
	p.SetPositionY(p.Y() + inc)
}

func (p *Particle) VelocityX() float64 {
	return p.velocity.X()
}

func (p *Particle) VelocityY() float64 {
	return p.velocity.Y()
}

func (p *Particle) SetVelocityX(vx float64) {
	p.velocity.SetX(vx)
}

func (p *Particle) SetVelocityY(vy float64) {
	p.velocity.SetY(vy)
}

func (p *Particle) IncrementVelocityX(inc float64) {
	p.SetVelocityX(p.VelocityX() + inc)
}

func (p *Particle) IncrementVelocityY(inc float64) {
	p.SetVelocityY(p.VelocityY() + inc)
}

// applies Earth's gravity force (m * 9.81 m/s^2) to each 
// particle
func (p *Particle) ComputeForce() vector.Vector2 {
	return *vector.MakeVector2(0, p.mass * -9.81)
}

