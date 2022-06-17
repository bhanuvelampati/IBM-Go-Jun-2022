/*
	Span 100 goroutines with ids (1 - 100) which does the following
		print "goroutine [id] started"
		wait for random (1 - 10) seconds
		print "goroutine [id] completed"

	Once all the goroutines finished their task the app should print "Done" and exit

*/