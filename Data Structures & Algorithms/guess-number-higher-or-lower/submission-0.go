/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	
	s,e := 1, n

	for (s <= e) {
		mid := s + (e-s)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			// left 
			e = mid -1
		} else {
			s = mid +1
		}
		mid = s + (e-s)/2

	}
	return -1
}
