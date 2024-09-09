/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"slices"
	"strconv"
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
		start, _ := cmd.Flags().GetInt64("start")
		topic, _ := cmd.Flags().GetString("topic")
		dryrun, _ := cmd.Flags().GetBool("dry-run")

		pattern_re := regexp.MustCompile(pattern)
		curdir, _ := os.Getwd()

		var already_prefixed []string
		var no_prefix []string

		candidates, _ := os.ReadDir(curdir)
		fmt.Println(pattern_re)
		for _, file := range candidates {
			fname := file.Name()
			if file.Type().IsDir() {
				continue
			}
			if !pattern_re.MatchString(fname) {
				continue
			}
			if slices.Contains(IGNORES, fname) {
				continue
			}
			if strings.HasPrefix(fname, topic) {
				already_prefixed = append(already_prefixed, fname)
			} else {
				no_prefix = append(no_prefix, fname)
			}
		}

		slices.Sort(already_prefixed)
		if len(already_prefixed) > 0 {
			last := already_prefixed[len(already_prefixed)-1]
			ext := path.Ext(last)
			no_ext := strings.TrimSuffix(last, ext)
			last_no_topic := no_ext[len(topic)+2:]
			parsed, err := strconv.ParseInt(last_no_topic, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			start = parsed
		}

		var failed []string
		for _, file := range no_prefix {
			ext := path.Ext(file)
			newName := fmt.Sprintf("%v--%03d%v", topic, start, ext)
			parent := path.Dir(file)
			newPath := path.Join(parent, newName)
			if dryrun {
				fmt.Printf("%v -> %v\n", path.Base(file), newName)
			} else {
				err := os.Rename(file, newPath)
				if err != nil {
					failed = append(failed, path.Base(file), err.Error())
				}
			}
			start += 1
		}

		if len(failed) > 0 {
			fmt.Println("Failed: ")
			for _, f := range failed {
				fmt.Println(f)
			}
		}

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
	seqnameCmd.Flags().Int64P("start", "s", 1, "First index")
	seqnameCmd.Flags().StringP("topic", "t", curdir, "Prefix for files")
}
