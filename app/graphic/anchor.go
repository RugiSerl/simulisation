package graphic

// Par Raphaël

// Fichier dans lequel on créer des ancres, elle permettent de fixer un rectangle , souvent associé à une texture, à un point et de calculer
// les coordonnées de ces derniers avec cette information

//énumération des différentes ancres disponibles
const ANCHOR_LEFT int8 = 0
const ANCHOR_RIGHT int8 = 1
const ANCHOR_HORIZONTAL_MiDDLE = 2

const ANCHOR_TOP int8 = 3
const ANCHOR_BOTTOM int8 = 4
const ANCHOR_VERTICAL_MiDDLE = 5

func GetRectCoordinatesWithAnchor(position Vector2, anchorX int8, anchorY int8, size Vector2, surfaceRect Rect) Vector2 {
	var DestVector Vector2

	if anchorX == ANCHOR_LEFT {
		DestVector.X = position.X + surfaceRect.X
	} else if anchorX == ANCHOR_RIGHT {
		DestVector.X = surfaceRect.Width + surfaceRect.X - position.X - size.X
	} else if anchorX == ANCHOR_HORIZONTAL_MiDDLE {
		DestVector.X = surfaceRect.X + surfaceRect.Width/2 - position.X - size.X/2
	}

	if anchorY == ANCHOR_TOP {
		DestVector.Y = position.Y + surfaceRect.Y
	} else if anchorY == ANCHOR_BOTTOM {
		DestVector.Y = surfaceRect.Height + surfaceRect.Y - position.Y - size.Y
	} else if anchorY == ANCHOR_VERTICAL_MiDDLE {
		DestVector.Y = surfaceRect.Y + surfaceRect.Height/2 - position.Y - size.Y/2
	}
	return DestVector

}
