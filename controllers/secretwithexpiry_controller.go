/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	expiryv1 "github.com/devops-360-online/k8s-secret-expiry-controller/api/v1"
)

// SecretWithExpiryReconciler reconciles a SecretWithExpiry object
type SecretWithExpirySpec struct {
	SecretName string     `json:"secretName"`
	ExpiryDate metav1.Time `json:"expiryDate"`
}

//+kubebuilder:rbac:groups=expiry.devops-360.online,resources=secretwithexpiries,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=expiry.devops-360.online,resources=secretwithexpiries/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=expiry.devops-360.online,resources=secretwithexpiries/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SecretWithExpiry object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *SecretWithExpiryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("secretwithexpiry", req.NamespacedName)

	var secretWithExpiry expiryv1.SecretWithExpiry
	if err := r.Get(ctx, req.NamespacedName, &secretWithExpiry); err != nil {
		// handle error
	}

	// Your logic here

	return ctrl.Result{}, nil
}

func (r *SecretWithExpiryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&expiryv1.SecretWithExpiry{}).
		Complete(r)
}
