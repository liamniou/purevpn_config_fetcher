package flags

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type PromptType int

const (
	String PromptType = iota
	// Int
	Password
)

type FlagData struct {
	Name      string
	Shorthand string
	Usage     string
	Default   string
	Type      PromptType
}

func AddFlags(cmd *cobra.Command, flags *[]FlagData) {
	cmdFlags := cmd.Flags()
	for _, f := range *flags {
		if f.Shorthand != "" {
			cmdFlags.StringP(f.Name, f.Shorthand, f.Default, f.Usage)
		} else {
			cmdFlags.String(f.Name, f.Default, f.Usage)
		}
	}
}

func validateEmptyInput(input string) error {
	if len(strings.TrimSpace(input)) < 1 {
		return errors.New("this input must not be empty")
	}
	return nil
}

// func validateIntegerNumberInput(input string) error {
// 	if _, err := strconv.ParseInt(input, 0, 64); err != nil {
// 		return errors.New("invalid number")
// 	}
// 	return nil
// }

func GetArgs(cmd *cobra.Command, flags *[]FlagData) ([]string, error) {
	ret := make([]string, len(*flags))
	for i, f := range *flags {
		val, err := getOrPrompt(cmd, &f)
		if err != nil {
			return nil, err
		}
		ret[i] = val
	}
	return ret, nil
}

func getOrPrompt(cmd *cobra.Command, flag *FlagData) (string, error) {
	val, _ := cmd.Flags().GetString((*flag).Name)
	if val != "" {
		return val, nil
	}

	switch flag.Type {
	case String:
		return promptString(cmd, (*flag).Name)
	case Password:
		return promptPassword(cmd, (*flag).Name)
	}
	return "", fmt.Errorf("prompt: unknown prompt type for %s", (*flag).Name)
}

func promptString(cmd *cobra.Command, name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateEmptyInput,
	}

	return prompt.Run()
}

func promptPassword(cmd *cobra.Command, name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateEmptyInput,
		Mask:     '*',
	}

	return prompt.Run()
}

// func PromptInteger(cmd *cobra.Command, name string) (int64, error) {
// 	val, err := cmd.Flags().GetInt64(name)
// 	if err == nil {
// 		return val, nil
// 	}

// 	prompt := promptui.Prompt{
// 		Label:    name,
// 		Validate: validateIntegerNumberInput,
// 	}

// 	promptResult, err := prompt.Run()
// 	if err != nil {
// 		return 0, err
// 	}

// 	parseInt, _ := strconv.ParseInt(promptResult, 0, 64)
// 	return parseInt, nil
// }
