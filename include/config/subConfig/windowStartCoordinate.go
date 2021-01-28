package subConfig

const (
	defaultStartX = 300
	defaultStartY = 300
	defaultEndX   = 900
	defaultEndY   = 900
)

//Проверяет корректны ли координаты взятые из конфига
func CheckCorrectPosition(coordinate []int)[]int{
	if !isCorrect(coordinate) {
		return []int{defaultStartX,defaultStartY,defaultEndX,defaultEndY}
	}
	return coordinate
}

func isCorrect(coordinate []int)bool{
	if incorrectCoordinateCount(coordinate){
		return false
	}

	if lessThanZero(coordinate){
		return false
	}

	if incorrectPosition(coordinate){
		return false
	}

	return true
}

func lessThanZero(c []int)bool{
	for _,v := range c {
		if v < 0 {
			return true
		}
	}
	return false
}

//c[0] - startX
//c[1] - startY
//c[2] - endX
//c[3] - endY
func incorrectPosition(c []int)bool{

	if c[0]>c[2]{
		return true
	}
	if c[1]>c[3]{
		return true
	}
	return false
}

func incorrectCoordinateCount(c []int)bool{
	if len(c) != 4 {
		return true
	}
	return false
}