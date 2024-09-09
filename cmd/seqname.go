/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var IGNORES []string = []string{"Thumbs.db", ".gitignore", ".stignore", ".DS_File"}

// seqnameCmd represents the seqname command
var seqnameCmd = &cobra.Command{
	Use:   "seqname PATTERN [FILES...]",
	Short: "Rename files in a sequence",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(cmd.Usage())
			return
		}

		pattern := args[0]
		files := args[1:]
		start, _ := cmd.Flags().GetInt("start")
		topic, _ := cmd.Flags().GetString("topic")
		dryrun, _ := cmd.Flags().GetBool("dry-run")

		pattern_re := regexp.MustCompile(pattern)
		curdir, _ := os.Getwd()
		candidates, err := os.ReadDir(curdir)
		if err != nil {
			log.Fatal(err)
		}
		var already_prefixed []os.DirEntry
		var no_prefix []os.DirEntry

		for _, file := range candidates {
			fname := file.Name()
			if file.Type().IsDir() {
				continue
			}
			if pattern_re.MatchString(fname) && !slices.Contains(IGNORES, fname) {
				if strings.HasPrefix(fname, topic) {
					already_prefixed = append(already_prefixed, fname)
				} else {
					no_prefix = append(no_prefix, fname)
				}
			}
		}

		fmt.Println(already_prefixed)
		fmt.Println("---")
		fmt.Println(no_prefix)

		// matches = sorted(
		//     [f for f in options if pattern.match(f.name) and f.name not in IGNORES]
		// )
		// matches_already_prefixed = [f for f in matches if f.name.startswith(topic)]
		// matches_no_prefix = [f for f in matches if not f.name.startswith(topic)]

		// i = start
		// if len(matches_already_prefixed) > 0:
		//     name_last_no_topic = matches_already_prefixed[-1].stem[len(topic) + 2:]
		//     i = int(name_last_no_topic)

		// failed = []
		// for m in matches_no_prefix:
		//     try:
		//         new = f"{topic}--{i:03d}{m.suffix}"
		//         newPath = Path(m.parent) / new
		//         if dry_run:
		//             print(f"{m.name} -> {new}")
		//         else:
		//             m.rename(newPath)
		//         i += 1
		//     except Exception as E:
		//         failed.append((m.stem, E))

		// if len(failed) > 0:
		//     print("Failed:")
		//     for f in failed:
		//         print(f)

	},
}

func init() {
	rootCmd.AddCommand(seqnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seqnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	curdir, _ := os.Getwd()
	seqnameCmd.Flags().BoolP("dry-run", "n", false, "Preview what would be renamed")
	seqnameCmd.Flags().IntP("start", "s", 1, "First index")
	seqnameCmd.Flags().StringP("topic", "t", curdir, "Prefix for files")
	seqnameCmd.Flags().
}
