package v1

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	kapi "k8s.io/kubernetes/pkg/api/v1"
)

// ClusterResourceQuota mirrors ResourceQuota at a cluster scope.  This object is easily convertible to
// synthetic ResourceQuota object to allow quota evaluation re-use.
type ClusterResourceQuota struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	kapi.ObjectMeta `json:"metadata"`

	// Spec defines the desired quota
	Spec ClusterResourceQuotaSpec `json:"spec"`

	// Status defines the actual enforced quota and its current usage
	Status ClusterResourceQuotaStatus `json:"status,omitempty"`
}

// ClusterResourceQuotaSpec defines the desired quota restrictions
type ClusterResourceQuotaSpec struct {
	// Selector is the selector used to match projects.
	// It should only select active projects on the scale of dozens (though it can select
	// many more less active projects).  These projects will contend on object creation through
	// this resource.
	Selector ClusterResourceQuotaSelector `json:"selector"`

	// Quota defines the desired quota
	Quota kapi.ResourceQuotaSpec `json:"quota"`
}

// ClusterResourceQuotaSelector is used to select projects.  At least one of LabelSelector or AnnotationSelector
// must present.  If only one is present, it is the only selection criteria.  If both are specified,
// the project must match both restrictions.
type ClusterResourceQuotaSelector struct {
	// LabelSelector is used to select projects by label.
	LabelSelector *unversioned.LabelSelector `json:"labels"`

	// AnnotationSelector is used to select projects by annotation.
	AnnotationSelector map[string]string `json:"annotations"`
}

// ClusterResourceQuotaStatus defines the actual enforced quota and its current usage
type ClusterResourceQuotaStatus struct {
	// Total defines the actual enforced quota and its current usage across all projects
	Total kapi.ResourceQuotaStatus `json:"total"`

	// Namespaces slices the usage by project.  This division allows for quick resolution of
	// deletion reconcilation inside of a single project without requiring a recalculation
	// across all projects.  This can be used to pull the deltas for a given project.
	Namespaces ResourceQuotasStatusByNamespace `json:"namespaces"`
}

// ClusterResourceQuotaList is a collection of ClusterResourceQuotas
type ClusterResourceQuotaList struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	unversioned.ListMeta `json:"metadata,omitempty"`

	// Items is a list of ClusterResourceQuotas
	Items []ClusterResourceQuota `json:"items"`
}

// ResourceQuotasStatusByNamespace bundles multiple ResourceQuotaStatusByNamespace
type ResourceQuotasStatusByNamespace []ResourceQuotaStatusByNamespace

// ResourceQuotaStatusByNamespace gives status for a particular project
type ResourceQuotaStatusByNamespace struct {
	// Namespace the project this status applies to
	Namespace string `json:"namespace"`

	// Status indicates how many resources have been consumed by this project
	Status kapi.ResourceQuotaStatus `json:"status"`
}

// AppliedClusterResourceQuota mirrors ClusterResourceQuota at a project scope, for projection
// into a project.  It allows a project-admin to know which ClusterResourceQuotas are applied to
// his project and their associated usage.
type AppliedClusterResourceQuota struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	kapi.ObjectMeta `json:"metadata"`

	// Spec defines the desired quota
	Spec ClusterResourceQuotaSpec `json:"spec"`

	// Status defines the actual enforced quota and its current usage
	Status ClusterResourceQuotaStatus `json:"status,omitempty"`
}

// AppliedClusterResourceQuotaList is a collection of AppliedClusterResourceQuotas
type AppliedClusterResourceQuotaList struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object's metadata.
	unversioned.ListMeta `json:"metadata,omitempty"`

	// Items is a list of AppliedClusterResourceQuota
	Items []AppliedClusterResourceQuota `json:"items"`
}
