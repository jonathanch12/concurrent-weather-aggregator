package main

func (s *stats) add(t float64) {
	if s.count == 0 {
		s.min = t
		s.max = t
	} else {
		if t < s.min {
			s.min = t
		}
		if t > s.max {
			s.max = t
		}
	}
	s.sum += t
	s.count++
}

func (s *stats) merge(o stats) {
	if o.count == 0 {
		return
	}
	if s.count == 0 {
		s.min = o.min
		s.max = o.max
	} else {
		if o.min < s.min {
			s.min = o.min
		}
		if o.max > s.max {
			s.max = o.max
		}
	}
	s.sum += o.sum
	s.count += o.count
}

func (s stats) total() float64 {
	return s.sum
}

func (s stats) mean() float64 {
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

func (s stats) minimum() float64 {
	return s.min
}

func (s stats) maximum() float64 {
	return s.max
}
