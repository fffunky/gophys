package particle

import (
	"gophys/vector"
	"testing"
)

func TestMakeParticle(t *testing.T) {
	posX := float64(10)
	posY := float64(-20)
	velX := float64(-30)
	velY := float64(40)
	mass := float64(50)

	actual := MakeParticle(
		*vector.MakeVector2(posX, posY),
		*vector.MakeVector2(velX, velY),
		mass,
	)

	expected := struct{
		px float64
		py float64
		vx float64
		vy float64
		m float64
	}{px: posX, py: posY, vx: velX, vy: velY, m: mass}

	if actual.X() != expected.px {
		t.Errorf("Expected particle.x to be %.2f, got=%.2f",
			expected.px, actual.X())
	}

	if actual.Y() != expected.py {
		t.Errorf("Expected particle.x to be %.2f, got=%.2f",
			expected.py, actual.Y())
	}

	if actual.VelocityX() != expected.vx {
		t.Errorf("Expected particle.VelocityX to be %.2f, got=%.2f",
			expected.vx, actual.VelocityX())
	}

	if actual.VelocityY() != expected.vy {
		t.Errorf("Expected particle.VelocityY to be %.2f, got=%.2f",
			expected.vy, actual.VelocityY())
	}

	if actual.Mass() != expected.m {
		t.Errorf("Expected particle.Mass to be %.2f, got=%.2f",
			expected.m, actual.Mass())
	}
}


func TestIncrementFuncs(t *testing.T) {
	posX := float64(10)
	posY := float64(-20)
	velX := float64(-30)
	velY := float64(40)
	mass := float64(50)

	actual := MakeParticle(
		*vector.MakeVector2(posX, posY),
		*vector.MakeVector2(velX, velY),
		mass,
	)

	iPosX := 1.0
	iPosY := -2.0
	iVelX := -3.0
	iVelY := 4.0

	expected := struct{
		px float64
		py float64
		vx float64
		vy float64
	}{
		px: posX + iPosX, 
		py: posY + iPosY, 
		vx: velX + iVelX, 
		vy: velY + iVelY, 
	}

	actual.IncrementPositionX(iPosX)
	if actual.X() != expected.px {
		t.Errorf("Expected particle.x to be %.2f, got=%.2f",
			expected.px, actual.X())
	}

	actual.IncrementPositionY(iPosY)
	if actual.Y() != expected.py {
		t.Errorf("Expected particle.x to be %.2f, got=%.2f",
			expected.py, actual.Y())
	}

	actual.IncrementVelocityX(iVelX)
	if actual.VelocityX() != expected.vx {
		t.Errorf("Expected particle.VelocityX to be %.2f, got=%.2f",
			expected.vx, actual.VelocityX())
	}

	actual.IncrementVelocityY(iVelY)
	if actual.VelocityY() != expected.vy {
		t.Errorf("Expected particle.VelocityY to be %.2f, got=%.2f",
			expected.vy, actual.VelocityY())
	}
}