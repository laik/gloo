package install

import (
	"bytes"
	"fmt"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/kubeutils"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"os/exec"
	"time"
)

func CheckKnativeInstallation() (isInstalled bool, isOurs bool, err error) {
	restCfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		return false, false, err
	}
	kube, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return false, false, err
	}
	namespaces, err := kube.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return false, false, err
	}
	for _, ns := range namespaces.Items {
		if ns.Name == constants.KnativeServingNamespace {
			ours := ns.Labels != nil && ns.Labels["app"] == "gloo"
			return true, ours, nil
		}
	}
	return false, false, nil
}

// Blocks until the given CRDs have been registered.
func WaitForCrdsToBeRegistered(crds []string, timeout, interval time.Duration) error {
	if len(crds) == 0 {
		return nil
	}

	// TODO: think about improving
	// Just pick the last crd in the list an wait for it to be created. It is reasonable to assume that by the time we
	// get to applying the manifest the other ones will be ready as well.
	crdName := crds[len(crds)-1]

	elapsed := time.Duration(0)
	for {
		select {
		case <-time.After(interval):
			if err := Kubectl(nil, "get", crdName); err == nil {
				return nil
			}
			elapsed += interval
			if elapsed > timeout {
				return errors.Errorf("failed to confirm knative crd registration after %v", timeout)
			}
		}
	}
}

//noinspection GoNameStartsWithPackageName
func InstallManifest(manifest []byte, isDryRun bool) error {
	if isDryRun {
		fmt.Printf("%s", manifest)
		// For safety, print a YAML separator so multiple invocations of this function will produce valid output
		fmt.Println("\n---")
		return nil
	}
	if err := kubectlApply(manifest); err != nil {
		return errors.Wrapf(err, "running kubectl apply on manifest")
	}
	return nil
}

func kubectlApply(manifest []byte) error {
	return Kubectl(bytes.NewBuffer(manifest), "apply", "-f", "-")
}

func Kubectl(stdin io.Reader, args ...string) error {
	kubectl := exec.Command("kubectl", args...)
	if stdin != nil {
		kubectl.Stdin = stdin
	}
	kubectl.Stdout = os.Stdout
	kubectl.Stderr = os.Stderr
	return kubectl.Run()
}
