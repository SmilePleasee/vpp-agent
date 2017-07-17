// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nametoidx

import (
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/idxmap"
	"github.com/ligato/cn-infra/idxmap/mem"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/logroot"
	"github.com/ligato/vpp-agent/idxvpp"
	"strconv"
	"time"
)

const idxKey = "idxKey"

type nameToIdxMeta struct {
	// added index
	idx uint32
	// original user's meta data
	meta interface{}
}

type nameToIdxMem struct {
	logging.Logger
	internal idxmap.NamedMappingRW
}

// NewNameToIdx create new instance implementing NameToIdxRW. Argument indexFunction might
// be nil if you do not want to use secondary indexes.
func NewNameToIdx(logger logging.Logger, owner core.PluginName, title string,
	indexFunction func(interface{}) map[string][]string) idxvpp.NameToIdxRW {
	m := nameToIdxMem{}
	m.Logger = logger
	m.internal = mem.NewNamedMapping(logger, owner, title, func(meta interface{}) map[string][]string {
		var idxs map[string][]string

		internalMeta, ok := meta.(*nameToIdxMeta)
		if !ok {
			return nil
		}
		if indexFunction != nil {
			idxs = indexFunction(internalMeta.meta)
		}
		if idxs == nil {
			idxs = map[string][]string{}
		}
		internal := indexInternalMetadata(meta)
		for k, v := range internal {
			idxs[k] = v
		}
		return idxs
	})
	return &m
}

// RegisterName from namedMapping allows to register a name-to-index mapping in-memory.
func (mem *nameToIdxMem) RegisterName(name string, idx uint32, metadata interface{}) {
	mem.internal.RegisterName(name, &nameToIdxMeta{idx, metadata})
}

// UnregisterName from namedMapping allows to remove mapping from the in-memory registry.
func (mem *nameToIdxMem) UnregisterName(name string) (idx uint32, metadata interface{}, found bool) {

	meta, found := mem.internal.UnregisterName(name)
	if found {
		if internalMeta, ok := meta.(*nameToIdxMeta); ok {
			return internalMeta.idx, internalMeta.meta, found
		}
	}
	return
}

// GetRegistryTitle returns a name assigned to mapping.
func (mem *nameToIdxMem) GetRegistryTitle() string {
	return mem.internal.GetRegistryTitle()
}

// LookupIdx allows to retrieve previously stored index for particular name.
func (mem *nameToIdxMem) LookupIdx(name string) (uint32, interface{}, bool) {

	meta, found := mem.internal.Lookup(name)
	if found {
		if internalMeta, ok := meta.(*nameToIdxMeta); ok {
			return internalMeta.idx, internalMeta.meta, found
		}
	}
	return 0, nil, false
}

func (mem *nameToIdxMem) LookupName(idx uint32) (name string, metadata interface{}, exists bool) {
	res := mem.internal.LookupByMetadata(idxKey, strconv.FormatUint(uint64(idx), 10))
	if len(res) != 1 {
		return
	}
	m, found := mem.internal.Lookup(res[0])
	if found {
		if internalMeta, ok := m.(*nameToIdxMeta); ok {
			return res[0], internalMeta.meta, found
		}
	}
	return
}

func (mem *nameToIdxMem) LookupNameByMetadata(key string, value string) []string {
	return mem.internal.LookupByMetadata(key, value)
}

func (mem *nameToIdxMem) ListNames() (names []string) {
	return mem.internal.ListNames()
}

func (mem *nameToIdxMem) Watch(subscriber core.PluginName, callback func(idxvpp.NameToIdxDto)) {
	watcher := func(dto idxmap.NamedMappingDto) {
		internalMeta, ok := dto.Metadata.(*nameToIdxMeta)
		if !ok {
			return
		}
		msg := idxvpp.NameToIdxDto{
			NameToIdxDtoWithoutMeta: idxvpp.NameToIdxDtoWithoutMeta{
				NamedMappingDtoWithoutMeta: dto.NamedMappingDtoWithoutMeta,
				Idx: internalMeta.idx},
			Metadata: internalMeta.meta,
		}
		callback(msg)
	}
	mem.internal.Watch(subscriber, watcher)
}

// ToChan is an utility that allows to received notification through a channel.
// If a notification can not be delivered until timeout, it is dropped.
func ToChan(ch chan idxvpp.NameToIdxDto) func(dto idxvpp.NameToIdxDto) {
	return func(dto idxvpp.NameToIdxDto) {
		select {
		case ch <- dto:
		case <-time.After(mem.NotifTimeout):
			logroot.Logger().Warn("Unable to deliver notification")
		}
	}
}

func indexInternalMetadata(metaData interface{}) map[string][]string {
	indexes := map[string][]string{}
	internalMeta, ok := metaData.(*nameToIdxMeta)
	if !ok || internalMeta == nil {
		return indexes
	}

	idx := internalMeta.idx
	indexes[idxKey] = []string{strconv.FormatUint(uint64(idx), 10)}

	return indexes
}
