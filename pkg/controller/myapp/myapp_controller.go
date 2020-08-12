package myapp

import (
	"context"
	"reflect"

	//added new
	"time"

	appv1alpha1 "github.com/huzefa51/myapp-operator/pkg/apis/myapp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	//Added new
	appsv1 "k8s.io/api/apps/v1"
	"github.com/huzefa51/myapp-operator/cmd/manager/apps"
	"k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var log = logf.Log.WithName("controller_myapp")


/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MyApp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
    return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
    return &ReconcileMyApp{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
    // Create a new controller
    c, err := controller.New("MyApp-controller", mgr, controller.Options{Reconciler: r})
    if err != nil {
        return err
    }

    // Watch for changes to primary resource MyApp
    err = c.Watch(&source.Kind{Type: &appv1alpha1.MyApp{}}, &handler.EnqueueRequestForObject{})
    if err != nil {
        return err
    }

    // TODO(user): Modify this to be the types you create that are owned by the primary resource
    // Watch for changes to secondary resource Pods and requeue the owner MyApp
    err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
        IsController: true,
        OwnerType:    &appv1alpha1.MyApp{},
    })
    if err != nil {
        return err
    }

    //New added
    err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appv1alpha1.MyApp{},
	})
    //Till here
    return nil
}

// blank assignment to verify that ReconcileMyApp implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMyApp{}

// ReconcileMyApp reconciles a MyApp object
type ReconcileMyApp struct {
    // This client, initialized using mgr.Client() above, is a split client
    // that reads objects from the cache and writes to the apiserver
    client client.Client
    scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MyApp object and makes changes based on the state read
// and what is in the MyApp.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMyApp) Reconcile(request reconcile.Request) (reconcile.Result, error) {
    reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
    reqLogger.Info("Reconciling MyApp")

    // Fetch the MyApp instance
    instance := &appv1alpha1.MyApp{}
    err := r.client.Get(context.TODO(), request.NamespacedName, instance)
    if err != nil {
        if errors.IsNotFound(err) {
            // Request object not found, could have been deleted after reconcile request.
            // Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
            // Return and don't requeue
            return reconcile.Result{}, nil
        }
        // Error reading the object - requeue the request.
        return reconcile.Result{}, err
    }
        // List all pods owned by this MyApp instance
	myApp := instance
        podList := &corev1.PodList{}
        lbs := map[string]string{
        "app":     myApp.Name,
        "version": "v0.1",
	}
        labelSelector := labels.SelectorFromSet(lbs)
        listOps := &client.ListOptions{Namespace: myApp.Namespace, LabelSelector: labelSelector}
        if err = r.client.List(context.TODO(), podList, listOps); err != nil {
                return reconcile.Result{}, err
	}



    // Count the pods that are pending or running as available
    var available []corev1.Pod
    for _, pod := range podList.Items {
        if pod.ObjectMeta.DeletionTimestamp != nil {
            continue
        }
        if pod.Status.Phase == corev1.PodRunning || pod.Status.Phase == corev1.PodPending {
            available = append(available, pod)
        }
    }
    numAvailable := int32(len(available))
    availableNames := []string{}
    for _, pod := range available {
        availableNames = append(availableNames, pod.ObjectMeta.Name)
    }



    // Update the status if necessary
    status := appv1alpha1.MyAppStatus{
        PodNames: availableNames,
    }
    if !reflect.DeepEqual(myApp.Status, status) {
        myApp.Status = status
        err = r.client.Status().Update(context.TODO(), myApp)
        if err != nil {
            reqLogger.Error(err, "Failed to update MyApp status")
            return reconcile.Result{}, err
        }
    }




    if numAvailable > myApp.Spec.Replicas {
        reqLogger.Info("Scaling down pods", "Currently available", numAvailable, "Required replicas", myApp.Spec.Replicas)
        diff := numAvailable - myApp.Spec.Replicas
        dpods := available[:diff]
        for _, dpod := range dpods {
            err = r.client.Delete(context.TODO(), &dpod)
            if err != nil {
                reqLogger.Error(err, "Failed to delete pod", "pod.name", dpod.Name)
                return reconcile.Result{}, err
            }
        }
        return reconcile.Result{Requeue: true}, nil
    }

    if numAvailable < myApp.Spec.Replicas {
        reqLogger.Info("Scaling up pods", "Currently available", numAvailable, "Required replicas", myApp.Spec.Replicas)
        // Define a new Pod object
        pod := newPodForCR(myApp)
        // Set myApp instance as the owner and controller
        if err := controllerutil.SetControllerReference(myApp, pod, r.scheme); err != nil {
            return reconcile.Result{}, err
        }
        err = r.client.Create(context.TODO(), pod)
        if err != nil {
            reqLogger.Error(err, "Failed to create pod", "pod.name", pod.Name)
            return reconcile.Result{}, err
        }
        return reconcile.Result{Requeue: true}, nil
    }

    apps := apps.GetApps(instance)
    existingAbc, abc := apps.Abc.GetDeployment()
    err = r.client.Get(context.TODO(), types.NamespacedName{Name: abc.Name, Namespace: abc.Namespace}, existingAbc)
    if err != nil && errors.IsNotFound(err) {
	reqLogger.Info("Creating Abc")
    if err = createK8sObject(instance, abc, r); err != nil {
	return reconcile.Result{}, err
    }
	return requeAfter(5, nil)
   }

    return reconcile.Result{}, nil
}




// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *appv1alpha1.MyApp) *corev1.Pod {
        labels := map[string]string{
            "app":     cr.Name,
            "version": "v0.1",
        }
        return &corev1.Pod{
            ObjectMeta: metav1.ObjectMeta{
                GenerateName: cr.Name + "-pod",
                Namespace:    cr.Namespace,
                Labels:       labels,
            },
            Spec: corev1.PodSpec{
                Containers: []corev1.Container{
                    {
                        Name:    "busybox",
                        Image:   "busybox",
                        Command: []string{"sleep", "3600"},
                    },
                },
            },
        }
    }



func createK8sObject(instance *appv1alpha1.MyApp, obj v1.Object, r *ReconcileMyApp) error {
	var err error
	err = controllerutil.SetControllerReference(instance, obj, r.scheme)

	if err != nil {
		return err
	}

	switch t := obj.(type) {
	case *appsv1.Deployment:
		err = r.client.Create(context.TODO(), t)
	}
	return err
}

func requeAfter(sec int, err error) (reconcile.Result, error) {
	t := time.Duration(sec)
	return reconcile.Result{RequeueAfter: time.Second * t}, err
}
