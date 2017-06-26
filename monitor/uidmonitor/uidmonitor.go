package uidmonitor

import (
	"fmt"
	"strconv"

	"github.com/aporeto-inc/trireme/constants"
	"github.com/aporeto-inc/trireme/monitor/linuxmonitor/cgnetcls"
	"github.com/aporeto-inc/trireme/monitor/rpcmonitor"
	"github.com/aporeto-inc/trireme/policy"
)

func UidMetadataExtractor(event *rpcmonitor.EventInfo) (*policy.PURuntime, error) {
	if event.Name == "" {
		return nil, fmt.Errorf("EventInfo PUName is empty")
	}
	if event.PUID == "" {
		return nil, fmt.Errorf("EventInfo PUID is empty")
	}

	runtimeTags := policy.NewTagsMap(map[string]string{})

	for k, v := range event.Tags {
		runtimeTags.Tags["@usr:"+k] = v
	}
	//Addd more thing here later
	options := policy.NewTagsMap(map[string]string{
		cgnetcls.PortTag:       "0",
		cgnetcls.CgroupNameTag: event.PUID,
	})

	if _, ok := runtimeTags.Tags[cgnetcls.PortTag]; ok {
		options.Tags[cgnetcls.PortTag] = runtimeTags.Tags[cgnetcls.PortTag]
	}

	options.Tags[cgnetcls.CgroupMarkTag] = strconv.FormatUint(cgnetcls.MarkVal(), 10)
	runtimeIps := policy.NewIPMap(map[string]string{"bridge": "0.0.0.0/0"})
	return policy.NewPURuntime(event.Name, runtimePID, runtimeTags, runtimeIps, constants.LinuxProcessPU, options), nil
}
