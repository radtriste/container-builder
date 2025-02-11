//go:build integration_podman

/*
 * Copyright 2023 Red Hat, Inc. and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cleaner

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PodmanTestSuite struct {
	suite.Suite
	LocalRegistry PodmanLocalRegistry
	RegistryID    string
	Podman        Podman
}

func (suite *PodmanTestSuite) SetupSuite() {
	localRegistry, registryID, podman := SetupPodmanSocket()
	if len(registryID) > 0 {
		suite.LocalRegistry = localRegistry
		suite.RegistryID = registryID
		suite.Podman = podman
	} else {
		assert.FailNow(suite.T(), "Initialization failed")
	}
}

func (suite *PodmanTestSuite) TearDownSuite() {
	registryID := suite.LocalRegistry.GetRegistryRunningID()
	if len(registryID) > 0 {
		PodmanTearDown(suite.LocalRegistry)
	} else {
		suite.LocalRegistry.StopRegistry()
	}
	suite.Podman.PurgeContainer("", REGISTRY_IMG_FULL)
}
