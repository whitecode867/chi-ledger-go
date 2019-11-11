package integration

import "testing"

func TestIntegrationSuite(t *testing.T) {
	t.Run(`Test API Get List Success`, TestTodosGetListSuccess)
}
