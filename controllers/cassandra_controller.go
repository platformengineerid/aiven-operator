// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/aiven/aiven-go-client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aiven/aiven-operator/api/v1alpha1"
)

// CassandraReconciler reconciles a Cassandra object
type CassandraReconciler struct {
	Controller
}

// +kubebuilder:rbac:groups=aiven.io,resources=cassandras,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=aiven.io,resources=cassandras/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=aiven.io,resources=cassandras/finalizers,verbs=update

func (r *CassandraReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.reconcileInstance(ctx, req, newGenericServiceHandler(newCassandraAdapter), &v1alpha1.Cassandra{})
}

// SetupWithManager sets up the controller with the Manager.
func (r *CassandraReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Cassandra{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}

func newCassandraAdapter(object client.Object) (serviceAdapter, error) {
	cassandra, ok := object.(*v1alpha1.Cassandra)
	if !ok {
		return nil, fmt.Errorf("object is not of type v1alpha1.Cassandra")
	}
	return &cassandraAdapter{cassandra}, nil
}

// cassandraAdapter handles an Aiven Cassandra service
type cassandraAdapter struct {
	*v1alpha1.Cassandra
}

func (a *cassandraAdapter) getObjectMeta() *metav1.ObjectMeta {
	return &a.ObjectMeta
}

func (a *cassandraAdapter) getServiceStatus() *v1alpha1.ServiceStatus {
	return &a.Status
}

func (a *cassandraAdapter) getServiceCommonSpec() *v1alpha1.ServiceCommonSpec {
	return &a.Spec.ServiceCommonSpec
}

func (a *cassandraAdapter) getUserConfig() any {
	return &a.Spec.UserConfig
}

func (a *cassandraAdapter) newSecret(s *aiven.Service) *corev1.Secret {
	name := a.Spec.ConnInfoSecretTarget.Name
	if name == "" {
		name = a.Name
	}

	stringData := map[string]string{
		"CASSANDRA_HOST":     s.URIParams["host"],
		"CASSANDRA_PORT":     s.URIParams["port"],
		"CASSANDRA_USER":     s.URIParams["user"],
		"CASSANDRA_PASSWORD": s.URIParams["password"],
		"CASSANDRA_URI":      s.URI,
		"CASSANDRA_HOSTS":    strings.Join(s.ConnectionInfo.CassandraHosts, ","),
	}

	// Removes empties
	for k, v := range stringData {
		if v == "" {
			delete(stringData, k)
		}
	}

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: a.Namespace},
		StringData: stringData,
	}
}

func (a *cassandraAdapter) newCreateRequest() aiven.CreateServiceRequest {
	return aiven.CreateServiceRequest{
		Cloud: a.Spec.CloudName,
		MaintenanceWindow: getMaintenanceWindow(
			a.Spec.MaintenanceWindowDow,
			a.Spec.MaintenanceWindowTime,
		),
		Plan:                  a.Spec.Plan,
		TerminationProtection: a.Spec.TerminationProtection,
		ServiceName:           a.Name,
		ServiceType:           "cassandra",
		ServiceIntegrations:   nil,
		DiskSpaceMB:           v1alpha1.ConvertDiscSpace(a.Spec.DiskSpace),
	}
}

func (a *cassandraAdapter) newUpdateRequest() aiven.UpdateServiceRequest {
	return aiven.UpdateServiceRequest{
		Cloud: a.Spec.CloudName,
		MaintenanceWindow: getMaintenanceWindow(
			a.Spec.MaintenanceWindowDow,
			a.Spec.MaintenanceWindowTime,
		),
		Plan:                  a.Spec.Plan,
		TerminationProtection: a.Spec.TerminationProtection,
		Powered:               true,
		DiskSpaceMB:           v1alpha1.ConvertDiscSpace(a.Spec.DiskSpace),
	}
}