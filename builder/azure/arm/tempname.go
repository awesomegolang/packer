// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in the project root for license information.

package arm

import (
	"fmt"

	"github.com/mitchellh/packer/builder/azure/common"
)

const (
	TempNameAlphabet     = "0123456789bcdfghjklmnpqrstvwxyz"
	TempPasswordAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type TempName struct {
	AdminPassword     string
	ComputeName       string
	DeploymentName    string
	ResourceGroupName string
	OSDiskName        string
}

func NewTempName() *TempName {
	tempName := &TempName{}

	suffix := common.RandomString(TempNameAlphabet, 10)
	tempName.ComputeName = fmt.Sprintf("pkrvm%s", suffix)
	tempName.DeploymentName = fmt.Sprintf("pkrdp%s", suffix)
	tempName.OSDiskName = fmt.Sprintf("pkros%s", suffix)
	tempName.ResourceGroupName = fmt.Sprintf("packer-Resource-Group-%s", suffix)

	tempName.AdminPassword = common.RandomString(TempPasswordAlphabet, 32)

	return tempName
}