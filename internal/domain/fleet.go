package domain

import "time"

// domain/fleet.go

type Robot struct {
	ID       string
	LastSeen time.Time
	IsActive bool
}

type Fleet struct {
	Robots map[string]*Robot
}

func (f *Fleet) UpdateRobotStatus(id string, active bool) {
	if robot, ok := f.Robots[id]; ok {
		robot.LastSeen = time.Now()
		robot.IsActive = active
	}
}
