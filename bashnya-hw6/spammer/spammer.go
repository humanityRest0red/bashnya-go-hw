package main

import "sync"

func RunPipeline(cmds ...cmd) {
	var wg sync.WaitGroup

	var in, out chan any
	for _, c := range cmds {
		in = out
		out = make(chan any)
		wg.Add(1)
		go func(c cmd, in, out chan any) {
			defer wg.Done()
			c(in, out)
			close(out)
		}(c, in, out)
	}
	wg.Wait()
}

func SelectUsers(in, out chan any) {
	// 	in - string
	// 	out - User
}

func SelectMessages(in, out chan any) {
	// 	in - User
	// 	out - MsgID
}

func CheckSpam(in, out chan any) {
	// in - MsgID
	// out - MsgData
}

func CombineResults(in, out chan any) {
	// in - MsgData
	// out - string
}
