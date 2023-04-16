package graphic

const ANCHOR_LEFT int8 = 0
const ANCHOR_RIGHT int8 = 1
const ANCHOR_TOP int8 = 2
const ANCHOR_BOTTOM int8 = 3

func GetRectCoordinatesWithAnchor(position Vector2, anchorX int8, anchorY int8, size Vector2, surfaceRect Rect) Vector2 {
	var DestVector Vector2

	if anchorX == ANCHOR_LEFT {
		DestVector.X = position.X + surfaceRect.X
	} else if anchorX == ANCHOR_RIGHT {
		DestVector.X = surfaceRect.Width + surfaceRect.X - position.X - size.X
	}

	if anchorY == ANCHOR_TOP {
		DestVector.Y = position.Y + surfaceRect.Y
	} else if anchorY == ANCHOR_BOTTOM {
		DestVector.Y = surfaceRect.Height + surfaceRect.Y - position.Y - size.Y

	}
	return DestVector

}
