package main

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func main() {

	var pingCmd = &cobra.Command{
		Use:   "ping [host]",
		Short: "Ping a host",
		Long:  `Send ICMP echo requests to a host to check its availability.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			host := args[0]
			fmt.Printf("Pinging %s...\n", host)
			count, _ := cmd.Flags().GetInt("count")
			fmt.Printf("Sending %d packets to %s...\n", count, args[0])
		},
	}

	pingCmd.Flags().IntP("count", "c", 4, "Number of echo requests to send")

	var tracerouteCmd = &cobra.Command{
		Use:   "traceroute [host]",
		Short: "Trace the route to a host",
		Long:  `Show the route packets take to a specified host.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			host := args[0]
			fmt.Printf("Tracing route to %s...\n", host)
			// Add your traceroute implementation here
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "MyCLI is a tool for network diagnostics",
		Long:  `An interactive command-line tool for diagnosing network problems.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to MyCLI!")
		},
	}

	var interactiveCmd = &cobra.Command{
		Use:   "interactive",
		Short: "Start an interactive session",
		Run: func(cmd *cobra.Command, args []string) {
			interactiveMode()
		},
	}

	var progressBarCmd = &cobra.Command{
		Use:   "progress",
		Short: "Progress bar demonstration",
		Run: func(cmd *cobra.Command, args []string) {
			progressBar()
		},
	}

	rootCmd.AddCommand(progressBarCmd)
	rootCmd.AddCommand(interactiveCmd)
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(tracerouteCmd)

	rootCmd.Execute()
}

func interactiveMode() {
	prompt := promptui.Select{
		Label: "Select a tool",
		Items: []string{"Ping", "Traceroute", "Quit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Ping":
		fmt.Println("Ping selected")
		// Add interactive ping logic here
	case "Traceroute":
		fmt.Println("Traceroute selected")
		// Add interactive traceroute logic here
	case "Quit":
		fmt.Println("Exiting...")
	}

	prompt1 := promptui.Prompt{Label: "Enter new password", Mask: '*'}
	password1, _ := prompt1.Run()

	prompt2 := promptui.Prompt{
		Label: "Confirm password",
		Mask:  '*',
		Validate: func(input string) error {
			if input != password1 {
				return fmt.Errorf("Passwords do not match")
			}
			return nil
		},
	}
	password2, _ := prompt2.Run()
	fmt.Println(fmt.Sprintf("Password successfully set! Password: %v", password2))

	prompt3 := promptui.Prompt{
		Label: "Enter your name",
	}
	name, err := prompt3.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
	} else {
		fmt.Printf("Hello, %s\n", name)
	}

	prompt4 := promptui.Prompt{
		Label: "Enter a positive number",
		Validate: func(input string) error {
			if input == "" || input[0] == '-' {
				return fmt.Errorf("Input must be a positive number")
			}
			return nil
		},
	}
	result1, _ := prompt4.Run()
	fmt.Printf("You entered: %s\n", result1)

	items := []string{"Option 1", "Option 2", "Option 3"}
	prompt5 := promptui.Select{
		Label: "Select an option",
		Items: items,
	}
	_, result2, err := prompt5.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
	} else {
		fmt.Printf("You chose: %s\n", result2)
	}

	type Item struct {
		Name  string
		Value int
	}
	items1 := []Item{
		{"First Option", 1},
		{"Second Option", 2},
		{"Third Option", 3},
	}
	prompt6 := promptui.Select{
		Label: "Choose an item",
		Items: items1,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "> {{ .Name | cyan }}", // Highlight active item
			Inactive: "  {{ .Name }}",
			Selected: "You chose {{ .Name | red }}",
		},
	}
	_, result3, err := prompt6.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
	} else {
		fmt.Printf("You chose: %v\n", result3)
	}

	items3 := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	prompt7 := promptui.Select{
		Label:             "Choose a fruit",
		Items:             items3,
		StartInSearchMode: true,
	}
	_, result4, _ := prompt7.Run()
	fmt.Printf("You chose: %s\n", result4)

	prompt8 := promptui.Prompt{
		Label: "Enter your password",
		Mask:  '*',
	}
	password3, err := prompt8.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
	} else {
		fmt.Println(fmt.Sprintf("Password successfully entered: %v", password3))
	}

	prompt9 := promptui.Prompt{Label: "Enter your username"}
	username, _ := prompt9.Run()

	prompt10 := promptui.Prompt{Label: "Enter your password", Mask: '*'}
	password, _ := prompt10.Run()

	fmt.Printf("Username: %s, Password: %s\n", username, password)

	prompt11 := promptui.Select{
		Label: "Custom Key Bindings",
		Items: []string{"Option A", "Option B", "Option C"},
		Keys: &promptui.SelectKeys{
			Prev: promptui.Key{Code: 'w'}, // Use 'w' for previous
			Next: promptui.Key{Code: 's'}, // Use 's' for next
		},
	}
	_, result5, _ := prompt11.Run()
	fmt.Printf("You chose: %s\n", result5)
}

func progressBar() {
	p := mpb.New() // Create a new progress pool

	total := 100
	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name("Processing: "),
			decor.CountersNoUnit("%d/%d"), // Show counts
		),
		mpb.AppendDecorators(
			decor.Percentage(),
			decor.Elapsed(decor.ET_STYLE_GO), // Show elapsed time
		),
	)

	for i := 0; i < total; i++ {
		time.Sleep(50 * time.Millisecond) // Simulate work
		bar.Increment()
	}

	p.Wait() // Wait for all bars to complete
	fmt.Println("Task complete!")

	p1 := mpb.New()
	total1 := 100
	bar1 := p1.AddBar(int64(total1))

	for i := 0; i < total1; i++ {
		time.Sleep(50 * time.Millisecond) // Simulate work
		bar1.Increment()
	}

	p1.Wait()

	p2 := mpb.New()
	total2 := 100
	bar2 := p2.AddBar(int64(total),
		mpb.BarWidth(60),                       // Set bar width
		mpb.BarStyle("=>-"),                    // Customize fill, head, and empty
		mpb.PrependDecorators(
			decor.Name("Task: "),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
		),
	)

	for i := 0; i < total2; i++ {
		time.Sleep(50 * time.Millisecond) // Simulate work
		bar2.Increment()
	}

	p2.Wait()

	p3 := mpb.New()
	total3 := 100
	total4 := 200

	bar3 := p3.AddBar(int64(total3),
		mpb.PrependDecorators(
			decor.Name("Task 1: "),
		),
		mpb.AppendDecorators(
			decor.Percentage(),
		),
	)

	bar4 := p3.AddBar(int64(total4),
		mpb.PrependDecorators(
			decor.Name("Task 2: "),
		),
		mpb.AppendDecorators(
			decor.CountersNoUnit("%d/%d"),
		),
	)

	for i := 0; i < total3 || i < total4; i++ {
		if i < total3 {
			bar3.Increment()
			time.Sleep(20 * time.Millisecond)
		}
		if i < total4 {
			bar4.Increment()
			time.Sleep(10 * time.Millisecond)
		}
	}

	p3.Wait()

	p4 := mpb.New()
	total5 := 100
	bar5 := p4.AddBar(int64(total),
		mpb.AppendDecorators(
			decor.OnComplete(
				decor.Name("Completed! "),
				"✔️",
			),
		),
	)

	for i := 0; i < total5; i++ {
		time.Sleep(50 * time.Millisecond) // Simulate work
		bar5.Increment()
	}

	p4.Wait()

	p5 := mpb.New()
	total6 := 100
	bar6 := p5.AddBar(int64(total6),
		mpb.PrependDecorators(
			decor.Name("Downloading: "),
			decor.CountersKibiByte("% .2f / % .2f"), // Show download in KiB/bytes
		),
		mpb.AppendDecorators(
			decor.Percentage(),
			decor.EwmaETA(decor.ET_STYLE_GO, 60),   // Estimated time remaining
			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60), // Speed in KiB/s
		),
	)

	for i := 0; i < total6; i++ {
		time.Sleep(50 * time.Millisecond) // Simulate work
		bar6.Increment()
	}

	p5.Wait()

	p8 := mpb.New()
	bar7 := p8.AddBar(0, mpb.BarRemoveOnComplete()) // Initialize with no total
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		bar7.IncrBy(1)
		bar7.SetTotal(int64(i+1), false) // Dynamically increase total
	}
	p8.Wait()
}
