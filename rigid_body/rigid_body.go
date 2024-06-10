package rigid_body

import (
	"gophys/utils"
	"gophys/vector"
	"math"
)

// 2D box shape
type BoxShape struct {
	width           float64
	height          float64
	mass            float64
	momentOfInertia float64
}

// Calculates the inertia of the box and stores it in BoxShape.momentOfInertia
func (b *BoxShape) CalculateBoxInertia() {
	m := b.mass
	w := b.width
	h := b.height
	b.momentOfInertia = (m * ((w * w) + (h * h))) / 12
}

// 2D rigid body
type RigidBody struct {
	shape           *BoxShape
	position        *vector.Vector2
	linearVelocity  *vector.Vector2
	force           *vector.Vector2
	angle           float64
	angularVelocity float64
	torque          float64
}

// Initializes rigid bodies with random positions and angles and zero
// linear and angular velocities.
// They're all initializes with a box shape of random dimensions.
func MakeRandomRigidBody(maxX, maxY, maxWidth, maxHeight float64) *RigidBody {
	box := &BoxShape{
		mass:  10,
		width: 1 + utils.RandomFloat64Range(0, maxWidth),
		height: 1 + utils.RandomFloat64Range(0, maxHeight),
	}
	box.CalculateBoxInertia()

	rbAngle := utils.RandomFloat64Range(0, 360) / float64(360) * math.Pi * 2

	return &RigidBody{
		shape: box,
		position: vector.MakeVector2(
			utils.RandomFloat64Range(0, maxX),
			utils.RandomFloat64Range(0, maxY),
		),
		linearVelocity: vector.MakeVector2(
			0, 0,
		),
		angle: rbAngle,
		angularVelocity: 0,
	}
}

func (r *RigidBody) X() float64 {
	return r.position.X()
}

func (r *RigidBody) Y() float64 {
	return r.position.Y()
}

func (r *RigidBody) Width() float64 {
	return r.shape.width
}

func (r *RigidBody) Height() float64 {
	return r.shape.height
}

func (r *RigidBody) Mass() float64 {
	return r.shape.mass
}

func (r *RigidBody) Angle() float64 {
	return r.angle
}

func (r *RigidBody) Torque() float64 {
	return r.torque
}

func (r *RigidBody) MomentOfInertia() float64 {
	return r.shape.momentOfInertia
} 

func (r *RigidBody) ForceX() float64 {
	return r.force.X()
}

func (r *RigidBody) ForceY() float64 {
	return r.force.Y()
}

func (r *RigidBody) SetAngle(a float64) {
	r.angle = a
}

func (r *RigidBody) IncrementAngle(inc float64) {
	r.SetAngle(r.Angle() + inc)
}

func (r *RigidBody) SetPositionX(x float64) {
	r.position.SetX(x)
}

func (r *RigidBody) SetPositionY(y float64) {
	r.position.SetY(y)
}

func (r *RigidBody) IncrementPositionX(inc float64) {
	r.SetPositionX(r.X() + inc)
}

func (r *RigidBody) IncrementPositionY(inc float64) {
	r.SetPositionY(r.Y() + inc)
}

func (r *RigidBody) LinVelX() float64 {
	return r.linearVelocity.X()
}

func (r *RigidBody) LinVelY() float64 {
	return r.linearVelocity.Y()
}

func (r *RigidBody) SetLinVelX(vx float64) {
	r.linearVelocity.SetX(vx)
}

func (r *RigidBody) SetLinVelY(vy float64) {
	r.linearVelocity.SetY(vy)
}

func (r *RigidBody) IncrementLinVelX(inc float64) {
	r.SetLinVelX(r.LinVelX() + inc)
}

func (r *RigidBody) IncrementLinVelY(inc float64) {
	r.SetLinVelY(r.LinVelY() + inc)
}

func (r *RigidBody) AngularVelocity() float64 {
	return r.angularVelocity
}

func (r *RigidBody) SetAngularVelocity(av float64) {
	r.angularVelocity = av
}

func (r *RigidBody) IncrementAngularVelocity(inc float64) {
	r.SetAngularVelocity(r.AngularVelocity() + inc)
}

// Applies a force at a point in the body, inducing some torque.
func (r *RigidBody) ComputeForceAndTorque() {
	f := vector.MakeVector2(0, 100)
	r.force = f

	// arm is the 'arm vector' that  goes from the center of mass to the
	// point of force application. Torque is the cross product of the arm and force.
	arm := vector.MakeVector2(r.Width() / 2, r.Height() / 2)
	r.torque = arm.X() * f.Y() - arm.Y() * f.X()
}