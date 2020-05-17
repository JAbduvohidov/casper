package casper

type frame struct {
	data string
}

type animation struct {
	frames   []*frame
}

func NewFrames(frames []string) []*frame {
	var f []*frame
	for _, fr:= range frames {
		f = append(f, &frame{data: fr})
	}
	return f
}

func NewAnimation(frames []*frame) *animation {
	return &animation{frames: frames}
}