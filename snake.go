package main

type position struct {
	x, y int
}

type direction string

type snake struct {
	body      []position
	direction direction
}

func move(s snake) snake {
	s = snake{
		body:      []position{{x: s.body[0].x, y: s.body[0].y}},
		direction: s.direction,
	}
	switch s.direction {
	case "up":
		s.body[0].y = s.body[0].y - 1
	case "down":
		s.body[0].y = s.body[0].y + 1
	case "right":
		s.body[0].x = s.body[0].x + 1
	case "left":
		s.body[0].x = s.body[0].x - 1
	}
	return s
}

func changeDirection(s snake, d direction) (ret_s snake) {
	//snake_d := snakeDirection(s, d)
	snake_d := snakeDirection(s, d)
	ret_s = snake{
		body:      []position{{x: s.body[0].x, y: s.body[0].y}},
		direction: snake_d,
	}
	return
}

func snakeDirection(s snake, d direction) direction {
	switch s.direction {
	case "up":
		if d == "down" {
			return "up"
		}
	case "down":
		if d == "up" {
			return "down"
		}
	case "right":
		if d == "left" {
			return "right"
		}
	case "left":
		if d == "right" {
			return "left"
		}
	}
	return d
}
