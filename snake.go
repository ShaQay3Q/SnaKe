package main

type position struct {
	x, y int
}

type direction string

type snake struct {
	body      []position
	direction direction
}

// Homework: re-right this function
func bananaMove(s snake) snake {
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

func move(s snake) (ret_s snake) {
	if s.direction == "up" {
		ret_s = snake{
			body:      []position{{x: s.body[0].x, y: s.body[0].y - 1}},
			direction: direction("up"),
		}
		return
	} else if s.direction == "down" {
		ret_s = snake{
			body:      []position{{x: s.body[0].x, y: s.body[0].y + 1}},
			direction: direction("down"),
		}
		return
	} else if s.direction == "right" {
		ret_s = snake{
			body:      []position{{x: s.body[0].x + 1, y: s.body[0].y}},
			direction: direction("right"),
		}
		return
	} else {
		ret_s = snake{
			body:      []position{{x: s.body[0].x - 1, y: s.body[0].y}},
			direction: direction("left"),
		}
		return
	}
}

func changeDirection(s snake, d direction) (ret_s snake) {
	//snake_d := snakeDirection(s, d)
	snake_d := banana(s, d)
	ret_s = snake{
		body:      []position{{x: s.body[0].x, y: s.body[0].y}},
		direction: snake_d,
	}
	return
}

func snakeDirection(s snake, d direction) (ret_d direction) {
	if s.direction == "up" {
		if d == "down" {
			ret_d = "up"
		} else {
			ret_d = d
		}
	} else if s.direction == "down" {
		if d == "up" {
			ret_d = "down"
		} else {
			ret_d = d
		}
	} else if s.direction == "right" {
		if d == "left" {
			ret_d = "right"
		} else {
			ret_d = d
		}
	} else {
		if d == "righ" {
			ret_d = "left"
		} else {
			ret_d = d
		}
	}
	return
}

func banana(s snake, d direction) direction {
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
