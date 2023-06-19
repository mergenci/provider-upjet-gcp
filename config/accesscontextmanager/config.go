package accesscontextmanager

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_access_context_manager_access_policy", func(r *config.Resource) {})
	// p.AddResourceConfigurator("google_access_context_manager_access_level", func(r *config.Resource) {
	// 	r.References["parent"] = config.Reference{
	// 		Type: "AccessPolicy",
	// 	}
	// })
	// p.AddResourceConfigurator("google_access_context_manager_service_perimeter_resource", func(r *config.Resource) {})
}
