/*
Copyright (c) 2020 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package component

import (
	"encoding/json"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	k8sJSON "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/tools/clientcmd/api/latest"
	"strings"
)

// Code is a component for code
type Editor struct {
	base
	Config EditorConfig `json:"config"`
}

// CodeConfig is the contents of Code
type EditorConfig struct {
	Code     string `json:"value"`
	ReadOnly bool   `json:"readOnly"`
}

// NewCodeBlock creates a code component
func NewEditor(title []TitleComponent, code string, readOnly bool) *Editor {
	return &Editor{
		base: newBase(typeEditor, title),
		Config: EditorConfig{
			Code:     code,
			ReadOnly: readOnly,
		},
	}
}

func (e *Editor) SetCodeFromObject(object runtime.Object) error {
	yamlSerializer := k8sJSON.NewYAMLSerializer(k8sJSON.DefaultMetaFactory, latest.Scheme, latest.Scheme)

	var sb strings.Builder
	if _, err := sb.WriteString("---\n"); err != nil {
		return err
	}
	if err := yamlSerializer.Encode(object, &sb); err != nil {
		return errors.Wrap(err, "encoding object as YAML")
	}

	e.Config.Code = sb.String()

	return nil
}

// GetMetadata returns the component's metadata.
func (e *Editor) GetMetadata() Metadata {
	return e.Metadata
}

type editorMarshal Editor

// MarshalJSON implements json.Marshaler
func (c *Editor) MarshalJSON() ([]byte, error) {
	m := editorMarshal(*c)
	m.Metadata.Type = typeEditor
	return json.Marshal(&m)
}
