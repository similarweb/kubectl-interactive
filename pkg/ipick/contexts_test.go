package ipick

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func MockNewContext(fileName string) (*ContextsManager, error) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	paths := []string{fmt.Sprintf("%s/testutils/%s", pwd, fileName)}
	return NewContexts(paths)

}
func TestNewContexts(t *testing.T) {

	t.Run("valid config path", func(t *testing.T) {
		_, err := MockNewContext("kubeconfig")

		if err != nil {
			t.Errorf("context instance not created, got: %v", err)
		}
	})
	t.Run("invalid valid config path", func(t *testing.T) {
		_, err := MockNewContext("invalid-path-kubeconfig")

		if err == nil {
			t.Errorf("context error should be return, got: %s, want: %s", "nil", "error")
		}
	})

}

func TestGetCurrentContext(t *testing.T) {
	context, err := MockNewContext("kubeconfig")
	if err != nil {
		t.Errorf("context instance not created, got: %v", err)
	}

	exemptedCurrentContext := "cluster-b"
	currentContext := context.GetCurrentContext()

	if exemptedCurrentContext != currentContext.Name {
		t.Errorf("incorrect current context, got: %s, want: %s.", currentContext.Name, exemptedCurrentContext)
	}
}

func TestGetContexts(t *testing.T) {
	context, err := MockNewContext("kubeconfig")
	if err != nil {
		t.Errorf("context instance not created, got: %v", err)
	}

	exemptedContexts := 2

	contexts := context.GetContexts()

	if len(contexts) != exemptedContexts {
		t.Errorf("context len if incorrect, got: %d, want: jobs > %d.", len(contexts), exemptedContexts)
	}

}

func TestSetContext(t *testing.T) {
	context, err := MockNewContext("kubeconfig")
	if err != nil {
		t.Errorf("context instance not created, got: %v", err)
	}

	exemptedCurrentContext := "cluster-b"
	currentContext := context.GetCurrentContext()

	if exemptedCurrentContext != currentContext.Name {
		t.Errorf("incorrect current context, got: %s, want: %s.", currentContext.Name, exemptedCurrentContext)
	}

	// New cluster set test
	exemptedNewCurrentContext := "cluster-b"

	context.SetContext(Context{Name: "cluster-b"})
	newCurrentContext := context.GetCurrentContext()

	if exemptedNewCurrentContext != newCurrentContext.Name {
		t.Errorf("incorrect current context, got: %s, want: %s.", newCurrentContext.Name, exemptedNewCurrentContext)
	}

}
