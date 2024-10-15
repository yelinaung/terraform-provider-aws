// Code generated by tools/tfsdk2fw/main.go. Manual editing is required.

package {{ .PackageName }}

import (
	"context"

	{{if .ImportFrameworkAttr }}"github.com/hashicorp/terraform-plugin-framework/attr"{{- end}}
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	{{if .ImportProviderFrameworkTypes }}fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"{{- end}}
)

// @FrameworkDataSource("{{ .TFTypeName }}")
func newDataSource{{ .Name }}(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &dataSource{{ .Name }}{}, nil
}

type dataSource{{ .Name }} struct {
	framework.DataSourceWithConfigure
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (d *dataSource{{ .Name }}) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = "{{ .TFTypeName }}"
}

// Schema returns the schema for this data source.
func (d *dataSource{{ .Name }}) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
    response.Schema = {{ .Schema }}
}

// Read is called when the provider must read data source values in order to update state.
// Config values should be read from the ReadRequest and new state values set on the ReadResponse.
func (d *dataSource{{ .Name }}) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data dataSource{{ .Name }}Data

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue("TODO")

    response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type dataSource{{ .Name }}Data struct {
    {{ .Struct }}
}