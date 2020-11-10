package controllers

import (
	"context"
	"github.com/go-logr/logr"
	cloudlinuxv1 "gitlab.com/cloudmanaged/operator/api/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sync"
)

// CloudManagedAlertReconciler reconciles a CloudManagedAlert object
type CloudManagedAlertReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	mu     sync.Mutex
}

// +kubebuilder:rbac:groups=cloudlinux.com,resources=cloudmanagedalerts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cloudlinux.com,resources=cloudmanagedalerts/status,verbs=get;update;patch
func (r *CloudManagedAlertReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("cloudmanagedalert", req.NamespacedName)

	r.mu.Lock()
	defer r.mu.Unlock()

	// Fetch the CloudManagedAlert instance
	cloudmanagedAlert := &cloudlinuxv1.CloudManagedAlert{}
	err := r.Get(ctx, req.NamespacedName, cloudmanagedAlert)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info(req.Namespace, req.Name, " has been deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get CloudManagedAlert")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *CloudManagedAlertReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudlinuxv1.CloudManagedAlert{}).
		Complete(r)
}
