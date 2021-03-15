package day23

import "strconv"

func StrongestReachable(bots Bots) int {
	var largestRadius, count int
	var largestPos Coordinate

	for c, rs := range bots {
		for _, r := range rs {
			if r > largestRadius {
				largestPos = c
				largestRadius = r
			}
		}
	}

	for c, rs := range bots {
		if largestPos.Distance(c) <= largestRadius {
			count += len(rs)
		}
	}

	return count
}

func ClosestSuccess(bots Bots) int {
	var cur, topLeft, bottomRight Coordinate
	zoom := 1 << (strconv.IntSize - 2)

	for {
		zoomedBots := make(Bots)
		best := struct {
			pos   Coordinate
			count int
		}{}

		for c, rs := range bots {
			for _, r := range rs {
				zc := Coordinate{c.X / zoom, c.Y / zoom, c.Z / zoom}
				zoomedBots[zc] = append(zoomedBots[zc], r/zoom)
			}
		}

		for cur.X = topLeft.X; cur.X <= bottomRight.X; cur.X++ {
			for cur.Y = topLeft.Y; cur.Y <= bottomRight.Y; cur.Y++ {
				for cur.Z = topLeft.Z; cur.Z <= bottomRight.Z; cur.Z++ {
					c := zoomedBots.HaveInRange(cur)

					// skip less bots
					if c < best.count {
						continue
					}
					// skip same amount of bots but Distance from Zero is the same or more
					if c == best.count && Zero.Distance(cur) >= Zero.Distance(best.pos) {
						continue
					}
					// more bots or same and closer to Zero
					best.pos, best.count = cur, c
				}
			}
		}

		// zoom in
		topLeft.X, topLeft.Y, topLeft.Z = (best.pos.X-1)<<1, (best.pos.Y-1)<<1, (best.pos.Z-1)<<1
		bottomRight.X, bottomRight.Y, bottomRight.Z = (best.pos.X+1)<<1, (best.pos.Y+1)<<1, (best.pos.Z+1)<<1
		zoom >>= 1

		if zoom == 0 {
			return Zero.Distance(best.pos)
		}
	}
}
