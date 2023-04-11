package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/godcong/go-iface"
	"github.com/godcong/go-iface/generator"
	"github.com/spf13/cobra"
)

const (
	programName = `iface`
)

type rootConfig struct {
	Target string
}

func main() {
	//var rc rootConfig
	fp := ""
	gen := generator.New()

	var rootCmd = &cobra.Command{
		Use:     programName,
		Short:   "generate interface from struct",
		Version: iface.Version,
		Run: func(cmd *cobra.Command, args []string) {
			ifaces, err := gen.GenerateFromPath(fp)
			if err != nil {
				panic(err)
			}
			for name, iface := range ifaces {
				mode := 0644
				outFilePath := filepath.Join(fp, name+".go")
				err = os.WriteFile(outFilePath, iface, os.FileMode(mode))
				if err != nil {
					panic(fmt.Errorf("failed writing to file %s: %s", outFilePath, err))
				}
			}
			fmt.Println("all finished!")
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if fp == "." {
				fmt.Println("generating to default setting path")
			}
		},
		DisableSuggestions: false,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true,
			HiddenDefaultCmd:    true,
		},
		SuggestionsMinimumDistance: 1,
	}
	rootCmd.PersistentFlags().StringVarP(&fp, "filepath", "p", ".", "set generate path")

	e := rootCmd.Execute()
	if e != nil {
		panic(e)
	}
}
