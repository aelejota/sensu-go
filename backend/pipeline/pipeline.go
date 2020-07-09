package pipeline

import (
	"context"
	"time"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/asset"
	"github.com/sensu/sensu-go/backend/licensing"
	"github.com/sensu/sensu-go/backend/secrets"
	"github.com/sensu/sensu-go/backend/store"
	"github.com/sensu/sensu-go/command"
	"github.com/sensu/sensu-go/rpc"
	"github.com/sensu/sensu-go/types"
)

// EventClient specifies the pipeline's requirements for event storage
type EventClient interface {
	FetchEvent(ctx context.Context, entity, check string) (*corev2.Event, error)
	ListEvents(ctx context.Context, pred *store.SelectionPredicate) ([]*corev2.Event, error)
}

// Pipeline takes events as inputs, and treats them in various ways according
// to the event's check configuration.
type Pipeline struct {
	store                  store.Store
	assetGetter            asset.Getter
	backendEntity          *corev2.Entity
	extensionExecutor      ExtensionExecutorGetterFunc
	executor               command.Executor
	storeTimeout           time.Duration
	secretsProviderManager *secrets.ProviderManager
	licenseGetter          licensing.Getter
	eventClient            EventClient
}

// Config holds the configuration for a Pipeline.
type Config struct {
	Store                   store.Store
	AssetGetter             asset.Getter
	BackendEntity           *corev2.Entity
	ExtensionExecutorGetter ExtensionExecutorGetterFunc
	StoreTimeout            time.Duration
	SecretsProviderManager  *secrets.ProviderManager
	LicenseGetter           licensing.Getter
	EventClient             EventClient
}

// Option is a functional option used to configure Pipelines.
type Option func(*Pipeline)

// New creates a new Pipeline from the provided configuration.
func New(c Config, options ...Option) *Pipeline {
	pipeline := &Pipeline{
		store:                  c.Store,
		assetGetter:            c.AssetGetter,
		backendEntity:          c.BackendEntity,
		extensionExecutor:      c.ExtensionExecutorGetter,
		executor:               command.NewExecutor(),
		storeTimeout:           c.StoreTimeout,
		secretsProviderManager: c.SecretsProviderManager,
		licenseGetter:          c.LicenseGetter,
		eventClient:            c.EventClient,
	}
	for _, o := range options {
		o(pipeline)
	}
	return pipeline
}

const (
	// DefaultSocketTimeout specifies the default socket dial
	// timeout in seconds for TCP and UDP handlers.
	DefaultSocketTimeout uint32 = 60
)

// ExtensionExecutorGetterFunc gets an ExtensionExecutor. Used to decouple
// pipelines from gRPC.
type ExtensionExecutorGetterFunc func(*types.Extension) (rpc.ExtensionExecutor, error)
