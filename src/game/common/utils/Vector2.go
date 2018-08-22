package utils

import (
	"math"
)

type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) *Vector2 {
	p := new(Vector2)
	p.X = x
	p.Y = y
	return p
}

func (p *Vector2) Copy(p1 *Vector2) {
	p.X = p1.X
	p.Y = p1.Y
}

func (p *Vector2) Add(p1 *Vector2) *Vector2 {
	p.X += p1.X
	p.Y += p1.Y
	return p
}

//func (*Vector2) Normalize(p *Vector2) *Vector2 {
//	val := 1.0 / float32(math.Sqrt(float64(p.X*p.X)+float64(p.Y*p.Y)))
//	return &Vector2{
//		X: p.X * val,
//		Y: p.Y * val,
//	}
//}

func (p *Vector2) Magnitude() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

func (p *Vector2) Normalize() *Vector2 {
	l := p.Magnitude()
	if l < 1E-5 && l > -1E-5 {
		p.X = 0
		p.Y = 0
		return p
	}
	val := 1.0 / l
	p.X *= val
	p.Y *= val
	return p
}

func (*Vector2) Lerp(a *Vector2, b *Vector2, v float32) *Vector2 {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	return NewVector2(a.X+(b.X-a.X)*v, a.Y+(b.Y-a.Y)*v)
}

func (p *Vector2) Distance(p1 *Vector2) float32 {
	return float32(math.Sqrt(math.Pow(float64(p.X-p1.X), 2) + math.Pow(float64(p.Y-p1.Y), 2)))
}
func (p *Vector2) DistanceSqr(p1 *Vector2) float32 {
	return (p.X-p1.X)*(p.X-p1.X) + (p.Y-p1.Y)*(p.Y-p1.Y)
}

func (p *Vector2) GetOffsetPoint(dir *Vector2, dis float32) *Vector2 {
	l := dir.Magnitude()
	if l < 1E-5 && l > -1E-5 {
		return p
	}
	tar := new(Vector2)
	tar.X = p.X + dis/l*(dir.X)
	tar.Y = p.Y + dis/l*(dir.Y)
	return tar
}

//暂时不支持负数dis
func (p0 *Vector2) MoveTowards(p1 *Vector2, dis float32) *Vector2 {
	if dis < 1E-5 {
		return &Vector2{
			X: p0.X,
			Y: p0.Y,
		}
	}
	p := NewVector2(0, 0)
	nowDis := p0.Distance(p1)
	if nowDis < dis {
		p.Copy(p1)
	} else {
		p.X = p0.X + (p1.X-p0.X)*dis/nowDis
		p.Y = p0.Y + (p1.Y-p0.Y)*dis/nowDis
	}
	return p
}

//返回与x轴正向夹角 [0,2*pi)
func (p *Vector2) CalAngle() (angle float64) {
	if p.X > 0 {
		angle = math.Atan(float64(p.Y) / float64(p.X))
	} else if p.X == 0 {
		if p.Y > 0 {
			angle = math.Pi / 2
		} else {
			angle = -math.Pi / 2
		}
	} else {
		angle = math.Pi + math.Atan(float64(p.Y)/float64(p.X))
	}
	if angle < 0 {
		angle += math.Pi * 2
	}
	return
}

//返回2个向量之间的夹角 (-pi,pi]
func (p *Vector2) CalAngleBetween(p1 *Vector2) (angle float64) {
	angleP := p.CalAngle()
	angleP1 := p1.CalAngle()
	angle = angleP1 - angleP
	if angle > math.Pi {
		angle -= math.Pi * 2
	}
	if angle <= -math.Pi {
		angle += math.Pi * 2
	}
	return
}

//返回vector绕原点旋转一定角度后的vector
func (p *Vector2) RotateAngle(angle float64) (p1 *Vector2) {
	a := p.CalAngle()
	p1 = NewVector2(float32(math.Cos(a+angle)), float32(math.Sin(a+angle)))
	return
}
