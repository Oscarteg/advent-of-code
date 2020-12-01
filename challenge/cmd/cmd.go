// Code generated by 'go run ./gen'; DO NOT EDIT
package cmd
import (
	"fmt"
	"time"
	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/oscarteg/advent-of-code-2020/challenge/day1"
	"github.com/oscarteg/advent-of-code-2020/challenge/example"
)
func addDays(root *cobra.Command) {
	example.AddCommandsTo(root)
	day1.AddCommandsTo(root)
}
type prof interface {
	Stop()
}
func NewRootCommand() *cobra.Command {
	var (
		start    time.Time
		profiler prof
	)
	result := &cobra.Command{
		Use:     "advent-of-code",
		Short:   "Advent of Code 2020 Solutions",
		Long:    "Golang implementations for the 2020 Advent of Code problems",
		Example: "go run main.go 1 a -i ./challenge/day1/input.txt",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if viper.GetBool("profile") {
				profiler = profile.Start()
			}
			start = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			if profiler != nil {
				profiler.Stop()
			}
			fmt.Println("Took", time.Since(start))
		},
	}
	addDays(result)
	flags := result.PersistentFlags()
	flags.StringP("input", "i", "", "Input File to read")
	_ = result.MarkPersistentFlagRequired("input")
	flags.Bool("profile", false, "Profile implementation performance")
	_ = viper.BindPFlags(flags)
	return result
}
