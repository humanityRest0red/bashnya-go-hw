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
	var numChannels = HasSpamMaxAsyncRequests
	channels := make([]chan User, numChannels)

	for i := range numChannels {
		channels[i] = make(chan User)
	}

	var wg sync.WaitGroup

	go func() {
		defer func() {
			for _, ch := range channels {
				close(ch)
			}
		}()

		i := 0
		for user := range in {
			chIdx := i % numChannels
			channels[chIdx] <- user.(User)
			i++
		}
	}()

	for _, ch := range channels {
		wg.Add(1)
		go func(c chan User) {
			defer wg.Done()
			var users []User
			for user := range c {
				users = append(users, user)
			}
			if msgIDs, err := GetMessages(users...); err == nil {
				for _, msg := range msgIDs {
					out <- msg
				}
			}
		}(ch)
	}

	wg.Wait()
}

func CheckSpam(in, out chan any) {
	// in - MsgID
	// out - MsgData
	var wg sync.WaitGroup

	sem := make(chan struct{}, HasSpamMaxAsyncRequests)

	for msgID := range in {
		wg.Add(1)
		sem <- struct{}{}
		go func(msgID any) {
			defer wg.Done()
			defer func() { <-sem }()
			if msgIDConv, ok := msgID.(MsgID); ok {
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
		}(msgData)
	}

	wg.Wait()
}
