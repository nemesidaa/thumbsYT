package logic

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nemesidaa/thumbsYT/CLI/internal/client"
	"github.com/nemesidaa/thumbsYT/CLI/internal/config"
)

func SafeExecution() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from:", r)
		}
	}()
}

func Start(cfg *config.ClientConfig) {

	stdin := make(chan Task)

	client, closeConn := client.NewClient(cfg)
	defer closeConn()

	log.Printf("Initialized connection to server at %s.\n", client.ServerAddr)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter commands (type 'q' to quit):")

		for scanner.Scan() {
			input := scanner.Text()
			task, args := parseCommand(input)
			if task == TermSig {
				close(stdin) // Закрываем канал, когда вводим "q"
				break
			}
			if task == NilSig {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
				continue
			}
			stdin <- task
			handleCommand(task, args, client)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		}
	}()

	// main loop to handle commands
	for command := range stdin {
		fmt.Printf("Received command: %s\n", command)
	}

}

// Handling func
func handleCommand(task Task, args []string, client *client.Client) {
	switch task {
	case LoadSig:
		fmt.Printf("Handling 'load' command with args: %s\n", args)
		err := client.HandleLoad(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
	case LogSig:
		fmt.Printf("Handling 'loglevel' command with args: %s\n", args)
		// LogLevel logic
	case HelpSig:
		fmt.Println("Available commands:")
		fmt.Println("  q - Quit")
		fmt.Println("  load <video_id> <resolution(opt)> - Load data")
		fmt.Println("  loglevel <debug|info|warn|error> - Set log level")
		fmt.Println("  help - Show this help message")
	default:
		fmt.Println("Unknown command!")
	}
}
