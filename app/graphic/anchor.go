package graphic

const ANCHOR_LEFT int8 = 0
const ANCHOR_RIGHT int8 = 1
const ANCHOR_TOP int8 = 2
const ANCHOR_BOTTOM int8 = 3

func GetObjectCoordinatesWithAnchor(x int32, y int32, anchorX int8, anchorY int8, objectWidth int32, objectHeight int32, surfaceRect *Rect) (int32, int32) {
	var (
		newX int32
		newY int32
	)

	if anchorX == ANCHOR_LEFT {
		newX = x + int32(surfaceRect.X)
	} else if anchorX == ANCHOR_RIGHT {
		newX = int32(surfaceRect.Width) + int32(surfaceRect.X) - x - objectWidth
	}

	if anchorY == ANCHOR_TOP {
		newY = y + int32(surfaceRect.Y)
	} else if anchorY == ANCHOR_BOTTOM {
		newY = int32(surfaceRect.Height) + int32(surfaceRect.Y) - y - objectHeight

	}
	return newX, newY

}
