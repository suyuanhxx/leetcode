package main

func ClimbStairs(n int) int {
	if (n <= 1) {
		return 1;
	}

	var nCur, nPrev, nPrevP int
	nPrevP = 1;
	nPrev = 2;
	for i := 3; i <= n; i++ {
		nCur = nPrev + nPrevP;
		nPrevP = nPrev;
		nPrev = nCur;
	}
	return nCur;
}

func pr() {
	//i,j :=0,0
}
