package cmd

import (
	"fmt"
	"github.com/bhbosman/goFit/cmd/Idl/parsers/fitGenerator"
	"github.com/bhbosman/goFit/cmd/Idl/parsers/idlGenerator"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path"
)

type nopCloserWriter struct {
	io.Writer
}

func (nopCloserWriter) Close() error { return nil }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "IdlReader",
	Short: "IdlReader",
	Long:  "IdlReader",
	Run: func(cmd *cobra.Command, args []string) {

		defer func() {
			r := recover()
			switch x := r.(type) {
			case error:
				_, _ = os.Stderr.Write([]byte(x.Error()))
				os.Exit(1)
			}
		}()

		var err error

		var generate string
		if generate, err = cmd.Flags().GetString("generate"); err != nil {
			_, _ = os.Stderr.Write([]byte(err.Error()))
			os.Exit(1)
			return
		}

		var packageName string
		if packageName, err = cmd.Flags().GetString("packageName"); err != nil {
			_, _ = os.Stderr.Write([]byte(err.Error()))
			os.Exit(1)
			return
		}

		var outputFile string
		if outputFile, err = cmd.Flags().GetString("outputFile"); err != nil {
			_, _ = os.Stderr.Write([]byte(err.Error()))
			os.Exit(1)
			return
		}

		var verbose bool
		if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
			_, _ = os.Stderr.Write([]byte(err.Error()))
			os.Exit(1)
			return
		}
		var output io.WriteCloser = nopCloserWriter{os.Stdout}
		ff := path.Join(os.TempDir(), "tempFileNameForOperation.go")
		if outputFile != "" {
			output, err = os.Create(ff)
			if err != nil {
				_, _ = os.Stderr.Write([]byte(err.Error()))
				os.Exit(1)
			}
		}
		for _, inputFile := range args {
			readCloser, err := os.Open(inputFile)
			if err != nil {
				_, _ = os.Stderr.Write([]byte(err.Error()))
				os.Exit(1)
			}

			generateIdl := idlGenerator.GenerateIdl
			setDebugLevel := idlGenerator.SetDebugLevel

			switch generate {
			case "fit":
				generateIdl = fitGenerator.GenerateFit
				setDebugLevel = fitGenerator.SetDebugLevel
			case "idl":
				generateIdl = idlGenerator.GenerateIdl
				setDebugLevel = idlGenerator.SetDebugLevel
				break
			default:
				generateIdl = idlGenerator.GenerateIdl
				setDebugLevel = idlGenerator.SetDebugLevel
				break
			}
			var debugLevel int
			if debugLevel, err = cmd.Flags().GetInt("debugLevel"); err != nil {
				_, _ = os.Stderr.Write([]byte(err.Error()))
				os.Exit(1)
				return
			}

			setDebugLevel(debugLevel)
			if err = generateIdl(
				inputFile,
				packageName,
				readCloser,
				output,
				func(
					fileName string,
				) (io.ReadCloser, string, error) {
					dir, _ := path.Split(inputFile)
					loadAssetName := path.Join(dir, fileName)
					if verbose {
						_, _ = io.WriteString(os.Stdout, fmt.Sprintf("trying to open %v\n", loadAssetName))
					}
					open, err := os.Open(loadAssetName)
					if err != nil {
						_, _ = os.Stderr.Write([]byte(err.Error()))
						os.Exit(1)
					}
					return open, loadAssetName, nil
				},
			); err != nil {
				_, _ = os.Stderr.Write([]byte(err.Error()))
				os.Exit(1)
			}
		}
		_ = output.Close()
		if outputFile != "" {
			if err = MoveFile(ff, outputFile); err != nil {
				_, _ = os.Stderr.Write([]byte(err.Error()))
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func MoveFile(sourcePath, destPath string) error {

	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer inputFile.Close()

	dir, _ := path.Split(destPath)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("couldn't open dest file: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("couldn't copy to dest from source: %v", err)
	}

	_ = inputFile.Close()

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't remove source file: %v", err)
	}
	return nil
}

func init() {
	rootCmd.Flags().IntP("debugLevel", "d", 0, "Debug Level")
	rootCmd.Flags().StringP("outputFile", "o", "", "output File")
	rootCmd.Flags().StringP("generate", "g", "idl", "Which Generator")
	rootCmd.Flags().BoolP("verbose", "v", false, "Which Generator")
	rootCmd.Flags().StringP("packageName", "p", "", "Which Generator")
}
