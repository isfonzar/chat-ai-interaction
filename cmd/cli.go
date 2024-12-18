package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"isfonzar/chat-ai-interaction/pkg/parser"

	"github.com/spf13/cobra"
)

// CLI variables
var (
	directory string
	dateInput string
)

const chatFileName = "_chat.txt"

// rootCmd defines the base command
var rootCmd = &cobra.Command{
	Use:   "chat-ai-interaction",
	Short: "From an exported chat, interact with the chat using AI",
	Long: `chat-ai-interaction processes exported chats and allows users to interact with the chat.
- Get a summary of the conversation in a date ("On January 1st, members of the project discussed the implementation of ...")
- Interact with the conversation by asking questions ("Who delivered the project at the city hall on march 1st?")`,
	Example: `  chat-ai-interaction -d /path/to/exported_chat -date 2024
  chat-ai-interaction -d /path/to/exported_chat -date 2024-03
  chat-ai-interaction -d /path/to/exported_chat -date 2024-03-21`,
	Args: cobra.NoArgs,
	Run:  runCommand,
}

// Initialize flags
func init() {
	rootCmd.Flags().StringVarP(&directory, "directory", "d", "", "Path to the WhatsApp exported directory (required)")
	rootCmd.Flags().StringVarP(&dateInput, "date", "D", "", "Date filter: 'YYYY', 'YYYY-MM', or 'YYYY-MM-DD' (required)")
	rootCmd.MarkFlagRequired("directory")
	rootCmd.MarkFlagRequired("date")
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// runCommand is the core logic for parsing inputs
func runCommand(_ *cobra.Command, _ []string) {
	// Validate directory
	chatFilePath := filepath.Join(directory, chatFileName)
	if _, err := os.Stat(chatFilePath); os.IsNotExist(err) {
		fmt.Printf("Error: '_chat.txt' not found in directory: %s\n", directory)
		os.Exit(1)
	}

	// Validate date input
	dateRegex := regexp.MustCompile(`^\d{4}(-\d{2})?(-\d{2})?$`)
	if !dateRegex.MatchString(dateInput) {
		fmt.Println("Error: Invalid date format. Use 'YYYY', 'YYYY-MM', or 'YYYY-MM-DD'.")
		os.Exit(1)
	}

	// Check for the required environment variable
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		fmt.Println("Error: Environment variable 'OPENAI_API_KEY' is not set.")
		os.Exit(1)
	}

	// Print parsed inputs
	fmt.Println("Starting to process the WhatsApp chat...")
	fmt.Printf("Directory: %s\n", directory)
	fmt.Printf("Date Input: %s\n", dateInput)

	// Process chat file
	processChatFile(chatFilePath, dateInput, openAIAPIKey)
}

// processChatFile processes the chat file (placeholder logic)
func processChatFile(chatFilePath, dateInput, openAIAPIKey string) {
	fmt.Printf("Processing file: %s for date: %s\n", chatFilePath, dateInput)
	fmt.Println("Output:")

	// Open the file (ensure this file exists or create it beforehand)
	file, err := os.Open(chatFilePath)
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Use the parser package to read from the file descriptor
	output, err := parser.Parse(file, dateInput, openAIAPIKey)
	if err != nil {
		fmt.Printf("error reading file content: %v\n", err)
		return
	}

	fmt.Println(output)
}
