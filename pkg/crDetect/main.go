package crDetect

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"time"
)

// Detector represents a procedure that runs in the background, periodically auto-detecting features
type Detector struct {
	dc                  discovery.DiscoveryInterface
	ticker              *time.Ticker
	SubscriptionChannel chan schema.GroupVersionKind
	crds 				map[runtime.Object]trigger
}

type trigger func(runtime.Object)

// New creates a new auto-detect runner
func NewAutoDetect(dc discovery.DiscoveryInterface) (*Detector, error) {
	// Create a new channel that GVK type will be sent down
	subChan := make(chan schema.GroupVersionKind, 1)

	return &Detector{dc: dc, SubscriptionChannel: subChan, crds:map[runtime.Object]trigger{}}, nil
}

func (d *Detector) AddCRDTrigger(crd runtime.Object, trigger trigger) {
	d.crds[crd] = trigger
}

// Start initializes the auto-detection process that runs in the background
func (d *Detector) Start(interval time.Duration) {
	// periodically attempts to auto detect all the capabilities for this operator
	d.ticker = time.NewTicker(interval * time.Second)

	go func() {
		d.autoDetectCapabilities()

		for range d.ticker.C {
			d.autoDetectCapabilities()
		}
	}()
}

// Stop causes the background process to stop auto detecting capabilities
func (d *Detector) Stop() {
	d.ticker.Stop()
	close(d.SubscriptionChannel)
}

func (d *Detector) autoDetectCapabilities() {
	for crd, trigger := range d.crds {
		crdgvk := crd.GetObjectKind().GroupVersionKind()
		resourceExists, _ := d.resourceExists(d.dc, crdgvk.GroupVersion().String(), crdgvk.Kind)
		if resourceExists {
			stateManager := GetStateManager()
			stateManager.SetState(crdgvk.Kind, true)
			trigger(crd)
		}
	}
}

//copied from operator-sdk to avoid pulling in thousands of files in imports
func (d *Detector) resourceExists(dc discovery.DiscoveryInterface, apiGroupVersion, kind string) (bool, error) {
	apiLists, err := dc.ServerResources()
	if err != nil {
		return false, err
	}
	for _, apiList := range apiLists {
		if apiList.GroupVersion == apiGroupVersion {
			for _, r := range apiList.APIResources {
				if r.Kind == kind {
					return true, nil
				}
			}
		}
	}
	return false, nil
}