package controllers

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"

	expiryv1 "github.com/devops-360-online/k8s-secret-expiry-controller/api/v1"
)

// SecretWithExpiryReconciler reconciles a SecretWithExpiry object
type SecretWithExpiryReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=expiry.devops-360.online,resources=secretwithexpiries,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=expiry.devops-360.online,resources=secretwithexpiries/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=expiry.devops-360.online,resources=secretwithexpiries/finalizers,verbs=update

func (r *SecretWithExpiryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("secretwithexpiry", req.NamespacedName)

	var secretWithExpiry expiryv1.SecretWithExpiry
	if err := r.Get(ctx, req.NamespacedName, &secretWithExpiry); err != nil {
		// handle error: if the error is related to the SecretWithExpiry object not being found, return without error
		if apierrs.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		// other errors should be returned as they would indicate a problem that needs to be resolved
		logger.Error(err, "Unable to fetch SecretWithExpiry")
		return ctrl.Result{}, err
	}

	// Check if secret is about to expire or has expired
	if time.Now().After(secretWithExpiry.Spec.ExpiryDate.Time) {
		// Secret has expired, generate error event
		r.Recorder.Event(&secretWithExpiry, corev1.EventTypeWarning, "SecretExpired", "The secret " + secretWithExpiry.Name + " in the namespace "+ secretWithExpiry.Namespace +" has expired.")
	} else if time.Now().Add(7 * 24 * time.Hour).After(secretWithExpiry.Spec.ExpiryDate.Time) {
		// Secret will expire in less than 24 hours, generate warning event
		r.Recorder.Event(&secretWithExpiry, corev1.EventTypeWarning, "SecretExpiring", "The secret" + secretWithExpiry.Name + "will expire in less than 7 days.")
	}

	return ctrl.Result{}, nil
}

func (r *SecretWithExpiryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&expiryv1.SecretWithExpiry{}).
		Complete(r)
}
