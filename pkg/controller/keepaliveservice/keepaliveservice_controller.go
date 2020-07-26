package keepaliveservice

import (
	"context"
	k8sv1alpha1 "github.com/myback/k8svc/pkg/apis/k8s/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner KeepAliveService
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
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
	//reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)

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

	//new Endpoint
	newEndp := newEndpoint(name, k8SvcData)
	if err := controllerutil.SetControllerReference(k8SvcData, newEndp, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundEndp := &corev1.Endpoints{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: newEndp.Namespace}, foundEndp)
	if err != nil && errors.IsNotFound(err) {
		if err := r.client.Create(context.Background(), newEndp); err != nil {
			return reconcile.Result{}, err
		}

	} else if err != nil {
		return reconcile.Result{}, err
	}

	//new Service
	newSvc := newService(name, k8SvcData)
	if err := controllerutil.SetControllerReference(k8SvcData, newSvc, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	foundSvc := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: newSvc.Namespace}, foundSvc)
	if err != nil && errors.IsNotFound(err) {
		if err := r.client.Create(context.TODO(), newSvc); err != nil {
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	//// Pod already exists - don't requeue
	//reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

func newEndpoint(name string, cr *k8sv1alpha1.KeepAliveService) *corev1.Endpoints {
	endpointAddrs := []corev1.EndpointAddress{}
	endpointPorts := []corev1.EndpointPort{}

	for _, host := range cr.Spec.Hosts {
		endpointAddrs = append(endpointAddrs, corev1.EndpointAddress{IP: host})
	}

	for _, port := range cr.Spec.Ports {
		endpointPorts = append(endpointPorts, corev1.EndpointPort{
			Name: port.Name,
			Port: port.Port,
		})
	}

	return &corev1.Endpoints{
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Namespace:   cr.Namespace,
			Labels:      cr.Spec.Template.Labels,
			Annotations: cr.Spec.Template.Annotations,
		},
		Subsets: []corev1.EndpointSubset{
			{
				Addresses: endpointAddrs,
				//NotReadyAddresses: endpointAddrs,
				Ports: endpointPorts,
			},
		},
	}
}

func newService(name string, cr *k8sv1alpha1.KeepAliveService) *corev1.Service {
	portsSpec := []corev1.ServicePort{}

	for _, portSpec := range cr.Spec.Ports {
		portsSpec = append(portsSpec, *portSpec.DeepCopyServicePort())
	}

	svcType := corev1.ServiceTypeClusterIP
	if cr.Spec.Type != "" {
		svcType = cr.Spec.Type
	}

	return &corev1.Service{
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Namespace:   cr.Namespace,
			Labels:      cr.Spec.Template.Labels,
			Annotations: cr.Spec.Template.Annotations,
		},
		Spec: corev1.ServiceSpec{
			Ports: portsSpec,
			Type:  svcType,
			//HealthCheckNodePort:      0,
		},
	}
}
