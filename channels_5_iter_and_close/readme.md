- submitWords() goes through the words and writes them to the channel "pending"
- countWords() reads from the channel "pending", counts the digits and writes to the channel "counted"
- fillStats() reads from the channel "counted" and fills in the final counter "stats"
In goroutines-writers use close() to signal to readers that there are no more values in the channel. In goroutines-readers use range to read from channels.