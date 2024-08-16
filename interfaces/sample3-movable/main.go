type Movable interface {
	Move(direction string)
}

func someFunc(obj any) {
	obj.Move("up")?
}

