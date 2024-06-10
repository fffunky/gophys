# gophys
A simple physics engine written in Go

# Features
- ability to create and manipulate physics data structures of different types and Set or increment their attributes as you please
- create randomized RigidBody and particles
- Run a simulation of random particles or random rigid bodies

# Types
Vector2:
```go
type Vector2 struct {
	x, y float64
}
```

Particle:
```go
type Particle struct {
	position vector.Vector2
	velocity vector.Vector2
	mass     float64
}
```

RigidBody (+ BoxShape):
```go
type BoxShape struct {
	width           float64
	height          float64
	mass            float64
	momentOfInertia float64
}

type RigidBody struct {
	shape           *BoxShape
	position        *vector.Vector2
	linearVelocity  *vector.Vector2
	force           *vector.Vector2
	angle           float64
	angularVelocity float64
	torque          float64
}
```
