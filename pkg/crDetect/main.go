package crDetect

import (
	"github.com/sirupsen/logrus"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"time"
)

// Background represents a procedure that runs in the background, periodically auto-detecting features
type Background struct {
	dc                  discovery.DiscoveryInterface
	ticker              *time.Ticker
	SubscriptionChannel chan schema.GroupVersionKind
	crds 				[]runtime.Object
}


// New creates a new auto-detect runner
func NewAutoDetect(dc discovery.DiscoveryInterface, CRDs []runtime.Object) (*Background, error) {
	// Create a new channel that GVK type will be sent down
	subChan := make(chan schema.GroupVersionKind, 1)

	return &Background{dc: dc, SubscriptionChannel: subChan, crds:CRDs}, nil
}

// Start initializes the auto-detection process that runs in the background
func (b *Background) Start(interval time.Duration) {
	// periodically attempts to auto detect all the capabilities for this operator
	b.ticker = time.NewTicker(interval * time.Second)

	go func() {
		b.autoDetectCapabilities()

		for range b.ticker.C {
			b.autoDetectCapabilities()
		}
	}()
}

// Stop causes the background process to stop auto detecting capabilities
func (b *Background) Stop() {
	b.ticker.Stop()
	close(b.SubscriptionChannel)
}

func (b *Background) autoDetectCapabilities() {
	logrus.Infof("detecting")
	for _, crd := range b.crds {
		crdgvk := crd.GetObjectKind().GroupVersionKind()
		logrus.Infof("looking for %s", crdgvk.String())
		resourceExists, _ := b.resourceExists(b.dc, crdgvk.GroupVersion().String(), crdgvk.Kind)
		if resourceExists {
			stateManager := GetStateManager()
			stateManager.SetState(crdgvk.Kind, true)

			b.SubscriptionChannel <- crdgvk.GroupVersion().WithKind(crdgvk.Kind)
		}
	}
}

//copied from operator-sdk to avoid pulling in thousands of files in imports
func (b *Background) resourceExists(dc discovery.DiscoveryInterface, apiGroupVersion, kind string) (bool, error) {
	apiLists, err := dc.ServerResources()
	if err != nil {
		return false, err
	}
	logrus.Infof("apiLists: %+v", apiLists)
	for _, apiList := range apiLists {
		logrus.Infof("%s ?? %s", apiList.GroupVersion, apiGroupVersion)
		if apiList.GroupVersion == apiGroupVersion {
			for _, r := range apiList.APIResources {
				logrus.Infof("%s ?? %s", r.Kind, kind)
				if r.Kind == kind {
					return true, nil
				}
			}
		}
	}
	return false, nil
}