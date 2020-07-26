package keepaliveservice

import (
	"context"
	k8sv1alpha1 "github.com/myback/k8svc/pkg/apis/k8s/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_keepaliveservice")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new KeepAliveService Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKeepAliveService{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("keepaliveservice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource KeepAliveService
	err = c.Watch(&source.Kind{Type: &k8sv1alpha1.KeepAliveService{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &k8sv1alpha1.KeepAliveService{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Endpoints{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &k8sv1alpha1.KeepAliveService{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileKeepAliveService implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileKeepAliveService{}

// ReconcileKeepAliveService reconciles a KeepAliveService object
type ReconcileKeepAliveService struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a KeepAliveService object and makes changes based on the state read
// and what is in the KeepAliveService.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKeepAliveService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)

	// Fetch the KeepAliveService data
	k8SvcData := &k8sv1alpha1.KeepAliveService{}

	name := request.Name
	if k8SvcData.Spec.Template.Name != "" {
		name = k8SvcData.Spec.Template.Name
	}

	if err := r.client.Get(context.TODO(), request.NamespacedName, k8SvcData); err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			//reqLogger.Info("DeleteObject", "Request.Name", request.Name)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	foundEndp := &corev1.Endpoints{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: request.Namespace}, foundEndp)
	if err != nil && errors.IsNotFound(err) {
		//new Endpoint
		newEndp := newEndpoint(name, k8SvcData)
		if err := controllerutil.SetControllerReference(k8SvcData, newEndp, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		if err := r.client.Create(context.Background(), newEndp); err != nil {
			return reconcile.Result{}, err
		}

	} else if err != nil {
		return reconcile.Result{}, err
	}

	if foundEndp.Name != "" {
		crdEndpSubs := keepAliveServiceEndpointSubsets(k8SvcData.Spec)
		if !endpointsEqual(crdEndpSubs, foundEndp.Subsets) {
			reqLogger.Info("Update endpoint state")
			foundEndp.Subsets = crdEndpSubs

			if err := r.client.Update(context.TODO(), foundEndp); err != nil {
				return reconcile.Result{}, err
			}
		}
	}

	foundSvc := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: request.Namespace}, foundSvc)
	if err != nil && errors.IsNotFound(err) {
		//new Service
		newSvc := newService(name, k8SvcData)
		if err := controllerutil.SetControllerReference(k8SvcData, newSvc, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		if err := r.client.Create(context.TODO(), newSvc); err != nil {
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	if foundSvc.Name != "" {
		crdSvc := keepAliveServiceServiceSpec(k8SvcData.Spec)
		if !serviceEqual(crdSvc, foundSvc.Spec) || crdSvc.Type != foundSvc.Spec.Type ||
			foundSvc.Spec.Selector != nil || foundSvc.Spec.ExternalIPs != nil {
			reqLogger.Info("Update service state")
			foundSvc.Spec.Ports = crdSvc.Ports
			foundSvc.Spec.Type = crdSvc.Type
			foundSvc.Spec.Selector = crdSvc.Selector
			foundSvc.Spec.ExternalIPs = crdSvc.ExternalIPs

			if err := r.client.Update(context.TODO(), foundSvc); err != nil {
				return reconcile.Result{}, err
			}
		}
	}

	return reconcile.Result{}, nil
}
