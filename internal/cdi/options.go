/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package cdi

import (
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvml"
)

// Option defines a function for passing options to the New() call
type Option func(*cdiHandler)

// WithDriverRoot provides an Option to set the driver root used by the 'cdi' interface
func WithDriverRoot(root string) Option {
	return func(c *cdiHandler) {
		c.driverRoot = root
	}
}

// WithNvidiaCTKPath provides an Option to set the nvidia-ctk path used by the 'cdi' interface
func WithNvidiaCTKPath(path string) Option {
	return func(c *cdiHandler) {
		c.nvidiaCTKPath = path
	}
}

// WithNvml provides an Option to set the NVML library used by the 'cdi' interface
func WithNvml(nvml nvml.Interface) Option {
	return func(c *cdiHandler) {
		c.nvml = nvml
	}
}

// WithDeviceIDStrategy provides an Option to set the device ID strategy used by the 'cdi' interface
func WithDeviceIDStrategy(strategy string) Option {
	return func(c *cdiHandler) {
		c.deviceIDStrategy = strategy
	}
}

// WithClass provides an Option to set the vendor used by the 'cdi' interface
func WithClass(class string) Option {
	return func(c *cdiHandler) {
		c.class = class
	}
}

// WithVendor provides an Option to set the vendor used by the 'cdi' interface
func WithVendor(vendor string) Option {
	return func(c *cdiHandler) {
		c.vendor = vendor
	}
}
