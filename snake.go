package main

func main() {

}

type position struct {
	x, y int
}

type direction string

type snake struct {
	body      []position
	direction direction
}

func move(s snake) snake {
	if s.direction == "up" {
		var ret_s = snake{
			body:      []position{{x: s.body[0].x, y: s.body[0].y - 1}},
			direction: direction("up"),
		}
		return ret_s
	} else if s.direction == "down" {
		var ret_s = snake{
			body:      []position{{x: s.body[0].x, y: s.body[0].y + 1}},
			direction: direction("down"),
		}
		return ret_s
	} else if s.direction == "right" {
		var ret_s = snake{
			body:      []position{{x: s.body[0].x + 1, y: s.body[0].y}},
			direction: direction("right"),
		}
		return ret_s
	} else {
		var ret_s = snake{
			body:      []position{{x: s.body[0].x - 1, y: s.body[0].y}},
			direction: direction("left"),
		}
		return ret_s
	}
}
