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

func changeDirection(s snake, d direction) snake {
	//snake_d := snakeDirection(s, d)
	snake_d := snakeDirection(s, d)
	ret_s := snake{
		body:      []position{{x: s.body[0].x, y: s.body[0].y}},
		direction: snake_d,
	}
	return ret_s
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

func newSnakeHead(s snake) position {
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

	var headPosition position
	headPosition.x = s.body[0].x
	headPosition.y = s.body[0].y

	// headP := position {x: snake.body[0].x, y: snake.body[0].y}

	return headPosition
}

func snakeGrow(s snake) (ret_s snake) {

	//snakeHeadPosition.body := move(s).body
	//s.body = append(s.body, snakeHeadPosition)

	return s
}

/*func snakeGrowth(s snake, ) (length int, capacity int, arr []position) {

	s = snake{
		body:      []position{},
		direction: s.direction,
	}

	length = len(s.body)
	capacity = cap(s.body)
	arr = s.body
	p := printSliceSlice(s.body)
	return

}

func printSliceSlice(s []position) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}*/
