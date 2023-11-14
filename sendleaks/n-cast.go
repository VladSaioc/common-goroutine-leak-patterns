package sendleaks

import "github.com/VladSaioc/common-goroutine-leak-patterns/utils"

func operateOnItem(item any) any {
	return utils.RandomValue(0, 1, 2, 3, 4)
}

// Process a number of items. First result to pass the post is retrieved from the channel queue.
func NCastLeak(items []any) {
	// Channel is synchronous.
	ch := make(chan any)

	// Iterate over every item
	for _, item := range items {
		go func(item any) {
			// Process item and send result to channel
			ch <- operateOnItem(item)
			// Channel is synchronous: only one sender will synchronise
		}(item)
	}
	// Retrieve first result. All other senders block.
	// Receiver blocks if there are no senders.
	<-ch
}

func FixedNCastLeak(items []any) {
	// Do not communicate if the list is empty. Receiver does not block
	if len(items) == 0 {
		return
	}
	// The maximum payload of the channel is len(items). All senders unblock
	ch := make(chan any, len(items))

	for _, item := range items {
		go func(item any) {
			ch <- operateOnItem(item)
		}(item)
	}
	// Retrieve first result. Senders do not unblock
	// Receiver is not executed if there are no senders.
	<-ch
}
