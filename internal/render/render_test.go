// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package render_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Helpers...

func load(t testing.TB, n string) *unstructured.Unstructured {
	raw, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", n))
	require.NoError(t, err)

	var o unstructured.Unstructured
	err = json.Unmarshal(raw, &o)
	require.NoError(t, err)

	return &o
}
