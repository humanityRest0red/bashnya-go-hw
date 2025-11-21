package main

import (
	"fmt"
	"sync"
)

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
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		seen = make(map[string]bool)
	)

	for email := range in {
		wg.Add(1)
		go func(email any) {
			defer wg.Done()
			if emailStr, ok := email.(string); ok {
				user := GetUser(emailStr)

				mu.Lock()
				_, exists := seen[user.Email]
				if !exists {
					seen[user.Email] = true
					out <- user
				}
				mu.Unlock()
			}
		}(email)
	}

	wg.Wait()
}

func SelectMessages(in, out chan any) {
	// 	in - User
	// 	out - MsgID
	var wg sync.WaitGroup

	for user := range in {
		wg.Add(1)
		go func(user any) {
			defer wg.Done()
			if userUser, ok := user.(User); ok {
				// if msgID, err := GetMessages(userUser); err == nil {
				// 	for _, elem := range msgID {
				// 		out <- elem
				// 	}
				// }
				msgID, _ := GetMessages(userUser)
				for _, elem := range msgID {
					out <- elem
				}
			}
		}(user)
	}

	wg.Wait()
}

func CheckSpam(in, out chan any) {
	// in - MsgID
	// out - MsgData
	var wg sync.WaitGroup

	for msgID := range in {
		wg.Add(1)
		go func(msgID any) {
			defer wg.Done()
			if msgIDConv, ok := msgID.(MsgID); ok {
				// if hasSpam, err := HasSpam(msgIDConv); err == nil {
				// 	out <- MsgData{ID: msgIDConv, HasSpam: hasSpam}
				// }
				hasSpam, _ := HasSpam(msgIDConv)
				out <- MsgData{ID: msgIDConv, HasSpam: hasSpam}
			}
		}(msgID)
	}

	wg.Wait()
}

func CombineResults(in, out chan any) {
	// in - MsgData
	// out - string
	var wg sync.WaitGroup

	for msgData := range in {
		wg.Add(1)
		go func(msgData any) {
			defer wg.Done()
			if msgDataConv, ok := msgData.(MsgData); ok {
				out <- fmt.Sprintf("%v %v", msgDataConv.HasSpam, msgDataConv.ID)
			}
			// msgDataConv, _ := msgData.(MsgData)
			// out <- fmt.Sprintf("%v %v", msgDataConv.HasSpam, msgDataConv.ID)

		}(msgData)
	}

	wg.Wait()
}
