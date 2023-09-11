package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnakeUp(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("up"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_moving := move(snake_before_moving)
	snake_after_2nd_moving := move(snake_after_moving)

	expected := snake{
		direction: direction("up"),
		body: []position{
			{x: 2, y: 3},
		},
	}

	expected_snake_before_moving := snake{
		direction: direction("up"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_2nd_moving := snake{
		direction: direction("up"),
		body:      []position{{x: 2, y: 2}},
	}

	require.Equal(t, expected, snake_after_moving)
	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_2nd_moving, snake_after_2nd_moving)

}

func TestSnakeDown(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("down"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_1stmove := move(snake_before_moving)

	expected_snake_before_moving := snake{
		direction: direction("down"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_1stmove := snake{
		direction: direction("down"),
		body:      []position{{x: 2, y: 5}},
	}

	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_1stmove, snake_after_1stmove)

}

func TestSnakeRight(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("right"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_1stmove := move(snake_before_moving)

	expected_snake_before_moving := snake{
		direction: direction("right"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_1stmove := snake{
		direction: direction("right"),
		body:      []position{{x: 3, y: 4}},
	}

	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_1stmove, snake_after_1stmove)
}

func TestSnakeLeft(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("left"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_1stmove := move(snake_before_moving)

	expected_snake_before_moving := snake{
		direction: direction("left"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_1stmove := snake{
		direction: direction("left"),
		body:      []position{{x: 1, y: 4}},
	}

	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_1stmove, snake_after_1stmove)

}

func TestChangeDirection(t *testing.T) {
	snake_before_moving := snake{
		body:      []position{{x: 5, y: 7}},
		direction: direction("up"),
	}
	snake_after_changing_direction := changeDirection(snake_before_moving, direction("right"))

	expected_snake_after_changing_direction := snake{
		body:      []position{{x: 5, y: 7}},
		direction: direction("right"),
	}

	snake_after_changing_direction_X2 := changeDirection(snake_after_changing_direction, direction("down"))

	expected_snake_after_changing_direction_X2 := snake{
		body:      []position{{x: 5, y: 7}},
		direction: direction("down"),
	}

	snake_after_changing_direction_X3 := changeDirection(snake_after_changing_direction_X2, "up")

	expected_snake_after_changing_direction_X3 := snake{
		body:      []position{{x: 5, y: 7}},
		direction: direction("down"),
	}

	require.Equal(t, expected_snake_after_changing_direction, snake_after_changing_direction)
	require.Equal(t, expected_snake_after_changing_direction_X2, snake_after_changing_direction_X2)
	require.Equal(t, expected_snake_after_changing_direction_X3, snake_after_changing_direction_X3)
}

func TestSnakeGrowth(t *testing.T) {

	snake_before_growing := snake{
		body:      []position{{6, 10}},
		direction: "up",
	}

	snake_after_growing_x1 := snakeGrow(snake_before_growing)

	expected_snake_after_growing_x1 := snake{
		body:      []position{{6, 9}, {6, 10}},
		direction: "up",
	}

	snake_after_growing_x2 := snakeGrow(snake_after_growing_x1)

	expected_snake_after_growing_x2 := snake{
		body:      []position{{6, 8}, {6, 9}, {6,10}},
		direction: "up",
	}

	require.Equal(t, expected_snake_after_growing_x1, snake_after_growing_x1)
	require.Equal(t, expected_snake_after_growing_x2, snake_after_growing_x2)
}

func TestNewSnakeHead (t *testing.T){

	snake := snake{
		body: []position{{6,9}, {6,10}},
		direction: "up",
	}

	snakeHead := newSnakeHead(snake)

	expected_snake_head := position {6, 8}



	require.Equal(t, expected_snake_head,snakeHead)

}
