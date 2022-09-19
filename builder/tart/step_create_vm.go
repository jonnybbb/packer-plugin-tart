package tart

import (
	"context"
	"fmt"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepCreateVM struct{}

func (s *stepCreateVM) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packersdk.Ui)

	ui.Say("Creating virtual machine...")

	if _, err := TartExec("create", "--from-ipsw", config.FromIPSW, config.VMName); err != nil {
		err := fmt.Errorf("Failed to create a VM: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())

		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepCreateVM) Cleanup(state multistep.StateBag) {
	// nothing to clean up
}
