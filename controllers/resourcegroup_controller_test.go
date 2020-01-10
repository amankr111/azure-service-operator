// +build all resourcegroup

package controllers

import (
	"context"
	"net/http"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestResourceGroupControllerHappyPath(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()
	ctx := context.Background()

	resourceGroupName := "t-rg-dev-" + helpers.RandomString(10)

	var err error

	// Create the ResourceGroup object and expect the Reconcile to be created
	resourceGroupInstance := &azurev1alpha1.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.ResourceGroupSpec{
			Location: tc.resourceGroupLocation,
		},
	}

	// create rg
	err = tc.k8sClient.Create(ctx, resourceGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}

	// verify sure rg has a finalizer
	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.HasFinalizer(finalizerName)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// verify rg gets submitted
	Eventually(func() string {
		_ = tc.k8sClient.Get(ctx, resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.Status.Message
	}, tc.timeout, tc.retry,
	).Should(ContainSubstring("successfully provisioned"))

	// verify rg exists in azure
	Eventually(func() bool {
		_, err := tc.resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		return err == nil
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// delete rg
	err = tc.k8sClient.Delete(ctx, resourceGroupInstance)
	Expect(err).NotTo(HaveOccurred())

	// verify rg is being deleted
	Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, resourceGroupNamespacedName, resourceGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// verify rg is gone from Azure
	Eventually(func() bool {
		result, _ := tc.resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		if result.Response == nil {
			return false
		}
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}
