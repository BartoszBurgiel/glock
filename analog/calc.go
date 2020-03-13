package watch

import "fmt"

// getDegree the arm of the clock has to lean to
// show the proper time
func getDegree(t int, scale int) float64 {
	return float64((t * 360) / scale)
}

// convert the degree to the representing
// byte and insert it into the watch face
// template
func angleToFace(deg float64) string {
	if deg < 22.5 {
		return fmt.Sprintf(watchFace, " ", "|", " ", " ", " ", " ", " ", " ")
	}

	if deg > 22.5 && deg < 67.5 {
		return fmt.Sprintf(watchFace, " ", " ", "/", " ", " ", " ", " ", " ")
	}

	if deg > 67.5 && deg < 112.5 {
		return fmt.Sprintf(watchFace, " ", " ", " ", " ", "-", " ", " ", " ")
	}

	if deg > 112.5 && deg < 157.5 {
		return fmt.Sprintf(watchFace, " ", " ", " ", " ", " ", " ", " ", "\\")
	}

	if deg > 157.5 && deg < 202.5 {
		return fmt.Sprintf(watchFace, " ", " ", " ", " ", " ", " ", "|", " ")
	}

	if deg > 202.5 && deg < 247.5 {
		return fmt.Sprintf(watchFace, " ", " ", " ", " ", " ", "/", " ", " ")
	}

	if deg > 202.5 && deg < 292.5 {
		return fmt.Sprintf(watchFace, " ", " ", " ", "-", " ", " ", " ", " ")
	}

	if deg > 292.5 && deg < 337.5 {
		return fmt.Sprintf(watchFace, "\\", " ", " ", " ", " ", " ", " ", " ")
	}

	return fmt.Sprintf(watchFace, " ", " ", " ", " ", " ", " ", " ", " ")
}
