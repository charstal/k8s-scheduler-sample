package sampleplugins

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	// Name plugins name
	Name = "sampleplugins"
)

// Args config args
type Args struct {
	KubeConfig string `json:"kubeconfig,omitempty"`
	Master     string `json:"master,omitempty"`
}

// SamplePlugins object
type SamplePlugins struct {
	args   *Args
	handle framework.FrameworkHandle
}

var (
	_ framework.FilterPlugin  = &SamplePlugins{}
	_ framework.PreBindPlugin = &SamplePlugins{}
)

// Name plugins name
func (s *SamplePlugins) Name() string {
	return Name
}

// Filter plugins interface
func (s *SamplePlugins) Filter(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	klog.V(3).Infof("filter pod: %v", pod.Name)
	return framework.NewStatus(framework.Success, "")
}

// PreBind plugins interface
func (s *SamplePlugins) PreBind(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	nodeInfo, ok := s.handle.NodeInfoSnapshot().NodeInfoMap["nodeName"]
	if !ok {
		return framework.NewStatus(framework.Error, "can't find")
	}
	klog.V(3).Infof("prebind node info: %+v", nodeInfo.Node())
	return framework.NewStatus(framework.Success, "")
}

// New plugins constructor
func New(plArgs *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(plArgs, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("####-> args: <-#### %+v", args)
	return &SamplePlugins{
		args:   args,
		handle: handle,
	}, nil
}