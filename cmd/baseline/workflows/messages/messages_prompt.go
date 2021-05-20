package messages

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/provideservices/provide-cli/cmd/common"
	"github.com/spf13/cobra"
)

var promptArgs []string

const promptStepSend = "Send"

// General Endpoints
func generalPrompt(cmd *cobra.Command, args []string, currentStep string) {
	switch step := currentStep; step {
	case promptStepSend:
		mandatoryFlagsPrompt()
		if flagPrompt(cmd, args) {
			optionalFlagsPrompt(cmd, args)
		}
		sendMessageRun(cmd, args)
	case "":
		emptyPrompt(cmd, args)
	}
}

func emptyPrompt(cmd *cobra.Command, args []string) {
	prompt := promptui.Select{
		Label: "What would you like to do",
		Items: []string{promptStepSend},
	}

	_, result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
		return
	}

	promptArgs = append(promptArgs, result)

	generalPrompt(cmd, args, result)
}

func flagPrompt(cmd *cobra.Command, args []string) bool {
	flagPrompt := promptui.Select{
		Label: "Would you like to set Optional Flags?",
		Items: []string{"No", "Yes"},
	}

	_, flagResult, err := flagPrompt.Run()

	if err != nil {
		os.Exit(1)
		return false
	}
	return flagResult == "Yes"
}

func optionalFlagsPrompt(cmd *cobra.Command, args []string) {
	if baselineID == "" {
		baselineIDFlagPrompt()
	}
}

func mandatoryFlagsPrompt() {
	if common.ApplicationID == "" {
		common.RequireWorkgroup()
	}
	if common.OrganizationID == "" {
		common.RequireOrganization()
	}
	if messageType == defaultBaselineMessageType {
		messageTypeFlagPrompt()
	}
	if id == "" {
		idFlagPrompt()
	}
	if data == "" {
		dataFlagPrompt()
	}
}

// Optional Flag
func baselineIDFlagPrompt() {
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Baseline Id",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
		return
	}

	baselineID = result
}

// Mandatory Flags
func dataFlagPrompt() {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("name cant be nil")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Data",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
		return
	}

	data = result
}

func idFlagPrompt() {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("name cant be nil")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "ID",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
		return
	}

	id = result
}

func messageTypeFlagPrompt() {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("name cant be nil")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Message Type",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		os.Exit(1)
		return
	}

	messageType = result
}