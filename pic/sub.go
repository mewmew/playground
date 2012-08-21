package pic

// HSubs returns the horizontal sub images in an image. Each sub image is
// separated by one or more transparent vertical line.
func HSubs(orig SubImager) (subs []SubImager) {
	rect := orig.Bounds()
	var sub SubImager
	for x := rect.Min.X; x < rect.Max.X; x++ {
		sub, x = GetHSub(orig, x)
		if sub != nil {
			subs = append(subs, sub)
		}
	}
	return subs
}

// GetHSub returns the first horizontal sub image located in orig, as well as
// the end x offset.
func GetHSub(orig SubImager, xStart int) (sub SubImager, xEnd int) {
	rect := orig.Bounds()
	subRect := rect
	var subFound bool
	var x int
	for x = xStart; x < rect.Max.X; x++ {
		if IsVLineAlpha(orig, x) {
			if subFound {
				break
			}
		} else {
			if !subFound {
				subRect.Min.X = x
				subFound = true
			}
		}
	}
	if !subFound {
		return nil, x
	}
	subRect.Max.X = x + 1
	sub = orig.SubImage(subRect).(SubImager)
	return Crop(sub), x
}

// VSubs returns the vertical sub images in an image. Each sub image is
// separated by one or more transparent horizontal line.
func VSubs(orig SubImager) (subs []SubImager) {
	/// ### todo ###
	///   - not yet implemented.
	/// ############
	return subs
}
